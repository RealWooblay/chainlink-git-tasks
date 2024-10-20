// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IGithubTasks {
    struct Task {
        address owner;
        uint256 org;
        uint256 repo;
        uint256 issueId;
        address[] tokens;
        uint256[] amounts;
        uint256 value;
        uint256 expiration;
        bytes32 requestId;
    }

    function newTask(
        address owner,
        uint256 org,
        uint256 repo,
        uint256 issueId,
        address[] memory tokens,
        uint256[] memory amounts,
        uint256 expiration
    ) external payable returns (uint256 taskId);

    function cancelTask(uint256 taskId) external;

    function endTask(
        uint256 taskId,
        bytes memory encryptedSecretsUrls,
        uint64 subscriptionId,
        uint32 gasLimit
    ) external returns (bytes32 requestId);
}
