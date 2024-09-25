// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Strings} from "@openzeppelin/utils/Strings.sol";

import {FunctionsClient} from "@chainlink/functions/v1_0_0/FunctionsClient.sol";
import {FunctionsRequest} from "@chainlink/functions/v1_0_0/libraries/FunctionsRequest.sol";

import {IWooblayUser} from "../token/IWooblayUser.sol";
import {IGithubConnector} from "./IGithubConnector.sol";

contract GithubConnector is IGithubConnector, FunctionsClient {
    using Strings for uint256;
    using FunctionsRequest for FunctionsRequest.Request;

    uint256 public constant NULL_GITHUB_ID = 0;
    uint256 public constant NULL_TOKEN_ID = 0;

    address public immutable wooblayUsers;

    bytes32 public immutable donId;

    // Token ID => GitHub ID + 1
    mapping(uint256 => uint256) private _githubIds;
    // Github ID => Token ID + 1
    mapping(uint256 => uint256) private _tokenIds;
    
    // Request ID => Token ID
    mapping(bytes32 => uint256) private _requests;

    constructor(address users, bytes32 don, address router) FunctionsClient(router) {
        wooblayUsers = users;
        donId = don;
    }

    function githubIdExists(uint256 githubId) public view override returns (bool) {
        return _tokenIds[githubId] != NULL_TOKEN_ID;
    }

    function hasGithubId(uint256 tokenId) public view override returns (bool) {
        return _githubIds[tokenId] != NULL_GITHUB_ID;
    }

    function getGithubId(uint256 tokenId) external view override returns (uint256) {
        if(!hasGithubId(tokenId)) {
            revert("");
        }
    
        return _getGithubId(tokenId);
    }

    function getTokenId(uint256 githubId) external view override returns (uint256) {
        if(!githubIdExists(githubId)) {
            revert("");
        }
    
        return _getTokenId(githubId);
    }

    function requestGithubId(
        bytes memory encryptedSecretsUrls,
        uint256 tokenId,
        uint64 subscriptionId,
        uint32 gasLimit
    ) external override returns (bytes32 requestId) {
        if(!IWooblayUser(wooblayUsers).tokenExists(tokenId)) {
            revert("Token doesn't exist!");
        }

        if(IWooblayUser(wooblayUsers).ownerOf(tokenId) != msg.sender) {
            revert("Token doesn't belong to you!");
        }

        if(!hasGithubId(tokenId)) {
            revert("Token already assigned a Github ID!");
        }

        FunctionsRequest.Request memory req;
        req.initializeRequestForInlineJavaScript(""); // TODO: Setup source text
        req.addSecretsReference(encryptedSecretsUrls);

        requestId = _sendRequest(
            req.encodeCBOR(),
            subscriptionId,
            gasLimit,
            donId
        );

        _requests[requestId] = tokenId;
    }

    function fulfillRequest(bytes32 requestId, bytes memory response, bytes memory err) internal override {
        uint256 githubId = abi.decode(response, (uint256));
        
        if(err.length > 0 || githubIdExists(githubId)) {
            // Invalid
            return;
        }

        uint256 tokenId = _requests[requestId];

        _setGithubId(tokenId, githubId);
        _setTokenId(tokenId, githubId);
    }
    
    function _getGithubId(uint256 tokenId) internal view returns(uint256) {
        return _githubIds[tokenId] - 1;
    }

    function _getTokenId(uint256 githubId) internal view returns(uint256) {
        return _tokenIds[githubId] - 1;
    }

    function _setGithubId(uint256 tokenId, uint256 githubId) internal {
        _githubIds[tokenId] = githubId + 1;
    }

    function _setTokenId(uint256 tokenId, uint256 githubId) internal {
        _tokenIds[githubId] = tokenId + 1;
    }
}