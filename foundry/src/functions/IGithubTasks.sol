// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IGithubTasks {
    struct Task {
        address owner;
        address[] tokens;
        uint256[] amounts;
        uint256 value;
        uint256 expiration;
    }

    function newTask(
        address owner,
        address[] memory tokens,
        uint256[] memory amounts,
        uint256 expiration
    ) external payable returns (uint256 taskId);

    function cancelTask(uint256 taskId) external;

    function endTask(uint256 taskId, uint256 githubId) external;
}
