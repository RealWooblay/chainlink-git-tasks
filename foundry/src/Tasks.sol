// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Ownable} from "@openzeppelin/access/Ownable.sol";

import {FunctionsClient} from "@chainlink/functions/v1_0_0/FunctionsClient.sol";
import {FunctionsRequest} from "@chainlink/functions/v1_0_0/libraries/FunctionsRequest.sol";

contract Tasks is FunctionsClient, Ownable {
    struct Repo {
        uint256 apiId;
        string owner;
        string name;
    }

    struct Org {
        Repo[] repos;
    }

    using FunctionsRequest for FunctionsRequest.Request;

    uint64 private _subscriptionId;

    mapping(uint256 => Org) private _orgs;

    constructor(
        address router,
        uint64 subId,
        address owner
    ) FunctionsClient(router) Ownable(owner) {
        _subscriptionId = subId;
    }

    function setSubscriptionId(uint64 subId) external onlyOwner {
        _subscriptionId = subId;
    }

    function setOrgsRepo(
        uint256 daoId,
        uint256 apiId,
        string memory owner,
        string memory name
    ) external onlyOwner {
        _orgs[daoId].repos.push(Repo(apiId, owner, name));
    }

    function request(uint32 gasLimit) external {}

    function fulfillRequest(
        bytes32 requestId,
        bytes memory response,
        bytes memory err
    ) internal override {}
}
