// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Ownable} from "@openzeppelin/access/Ownable.sol";

import {IWooblayProfile} from "../token/IWooblayProfile.sol";

import {IGithubConnector} from "./IGithubConnector.sol";

contract GithubConnector is Ownable, IGithubConnector {
    address public immutable wooblayProfiles;

    uint256 public constant NULL_GITHUB_ID = 0;
    uint256 public constant NULL_TOKEN_ID = 0;

    // Token ID => GitHub ID + 1
    mapping(uint256 => uint256) private _githubIds;
    // Github ID => Token ID + 1
    mapping(uint256 => uint256) private _tokenIds;

    constructor(address profiles, address initialOwner) Ownable(initialOwner) {
        wooblayProfiles = profiles;
    }

    function githubIdExists(uint256 githubId) public view override returns (bool) {
        return _tokenIds[githubId] != NULL_TOKEN_ID;
    }

    function getTokenId(uint256 githubId) external view override returns (uint256) {
        if(_tokenIds[githubId] == NULL_TOKEN_ID) {
            revert("");
        }
    
        return _tokenIds[githubId] - 1;
    }

    function getGithubId(uint256 tokenId) external view override returns (uint256) {
        if(_githubIds[tokenId] == NULL_GITHUB_ID) {
            revert("");
        }
    
        return _githubIds[tokenId] - 1;
    }

    function setGithubId(uint256 tokenId, uint256 githubId) external override onlyOwner {
        if(githubIdExists(githubId)) {
            revert("Github ID already assigned!");
        }

        if(IWooblayProfile(wooblayProfiles).tokenExists(tokenId)) {
            revert("");
        }

        _githubIds[tokenId] = githubId + 1;
        _tokenIds[githubId] = tokenId + 1;
    }
}