// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IGithubConnector {
    function githubIdExists(uint256 githubId) external view returns (bool);

    function getTokenId(uint256 githubId) external view returns (uint256);
    function getGithubId(uint256 tokenId) external view returns (uint256);

    function setGithubId(uint256 tokenId, uint256 githubId) external;
}