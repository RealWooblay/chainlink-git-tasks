// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IGithubPay {
    function pay(
        uint256 githubId,
        address[] memory tokens,
        uint256[] memory amounts
    ) external payable;
    function claim(uint256 githubId, address[] memory tokens) external;
}
