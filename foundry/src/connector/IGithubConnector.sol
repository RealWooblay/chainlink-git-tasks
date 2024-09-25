// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IGithubConnector {
    function hasGithubId(uint256 tokenId) external view returns (bool);
    function githubIdExists(uint256 githubId) external view returns (bool);

    function getTokenId(uint256 githubId) external view returns (uint256);
    function getGithubId(uint256 tokenId) external view returns (uint256);

    function requestGithubId(
        bytes memory encryptedSecretsUrls,
        uint256 tokenId,
        uint64 subscriptionId,
        uint32 gasLimit
    ) external returns (bytes32 requestId);
}
