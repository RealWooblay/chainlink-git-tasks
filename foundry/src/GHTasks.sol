// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {IGHTasks} from "./IGHTasks.sol";

import {Ownable} from "@openzeppelin/access/Ownable.sol";

import {FunctionsClient} from "@chainlink/functions/v1_0_0/FunctionsClient.sol";
import {FunctionsRequest} from "@chainlink/functions/v1_0_0/libraries/FunctionsRequest.sol";

contract GHTasks is IGHTasks, FunctionsClient, Ownable {
    using FunctionsRequest for FunctionsRequest.Request;

    uint64 public subscriptionId;

    uint8 public donSecretsSlotId;
    uint64 public donSecretsVersion;

    bytes32 public donId;

    mapping(uint256 => string) public sources;

    mapping(uint256 => Org) private _orgs;
    uint256 public totalOrgs;

    constructor(
        address router,
        uint64 subId,
        bytes32 don,
        address initialOwner
    ) FunctionsClient(router) Ownable(initialOwner) {
        subscriptionId = subId;
        donId = don;
    }

    function newOrg(
        address[] memory initalOwners
    ) external override returns (uint256 orgId) {
        orgId = totalOrgs;
        totalOrgs++;

        for (uint256 i = 0; i < initalOwners.length; i++) {
            _orgs[orgId].owners[initalOwners[i]] = true;
        }
    
        emit NewOrg(orgId, initalOwners);
    }

    function setOrgOwners(
        uint256 orgId,
        address[] memory owners,
        bool[] memory values
    ) external override {
        if (owners.length != values.length) {
            revert("Address and values arrays unequal length!");
        }

        if (!_orgs[orgId].owners[msg.sender]) {
            revert("Sender is not an org owner!");
        }

        for (uint256 i = 0; i < owners.length; i++) {
            _orgs[orgId].owners[owners[i]] = values[i];
        }

        emit SetOrgOwners(orgId, owners, values);
    }

    function setOrgRepo(
        uint256 orgId,
        string memory repoOwner,
        string memory repoName,
        bool value
    ) external override onlyOwner {
        _orgs[orgId].repos[repoOwner][repoName] = value;

        emit SetOrgRepo(orgId, repoOwner, repoName, value);
    }

    function setSource(
        uint256 sourceId,
        string memory source
    ) external onlyOwner {
        sources[sourceId] = source;

        emit SetSource(sourceId, source);
    }

    function request(
        uint256 orgId,
        uint256 sourceId,
        string[] memory args,
        uint32 gasLimit
    ) external override returns (bytes32 requestId) {
        if (bytes(sources[sourceId]).length == 0) {
            revert("Nil source!");
        }

        if (!_orgs[orgId].owners[address(0)] || !_orgs[orgId].owners[msg.sender]) {
            revert("Sender is not an org owner!");
        }

        if (!_orgs[orgId].repos[args[0]][args[1]]) {
            revert("Org does not manage repo!");
        }

        FunctionsRequest.Request memory req;

        req.initializeRequestForInlineJavaScript(sources[sourceId]);
        req.setArgs(args);

        if (donSecretsVersion > 0) {
            req.addDONHostedSecrets(donSecretsSlotId, donSecretsVersion);
        }

        requestId = _sendRequest(req.encodeCBOR(), subscriptionId, gasLimit, donId);

        emit Request(orgId, sourceId, args, requestId);
    }

    function fulfillRequest(
        bytes32 requestId,
        bytes memory response,
        bytes memory err
    ) internal override {
        emit RequestFulfilled(requestId, response, err);        
    }
}
