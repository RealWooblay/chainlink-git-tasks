// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IERC20} from "@openzeppelin/token/ERC20/IERC20.sol";
import {Address} from "@openzeppelin/utils/Address.sol";

import {IWooblayUsers} from "../token/IWooblayUsers.sol";
import {IGithubConnector} from "../connector/IGithubConnector.sol";
import {IGithubPay} from "../connector/IGithubPay.sol";
import {IGithubTasks} from "./IGithubTasks.sol";

contract GithubTasks is IGithubTasks {
    using Address for address payable;

    address public immutable wooblayUsers;
    address public immutable githubConnector;
    address public immutable githubPay;

    // Task ID => Task
    mapping(uint256 => Task) private _tasks;
    uint256 private _totalTasks;

    constructor(address users, address connector, address pay) {
        wooblayUsers = users;
        githubConnector = connector;
        githubPay = pay;
    }

    function newTask(
        address owner,
        address[] memory tokens,
        uint256[] memory amounts,
        uint256 expiration
    ) external payable override returns (uint256 taskId) {
        if (tokens.length != amounts.length) {
            revert();
        }

        for (uint256 i = 0; i < tokens.length; i++) {
            IERC20(tokens[i]).transferFrom(
                msg.sender,
                address(this),
                amounts[i]
            );
        }

        taskId = _totalTasks;
        _totalTasks++;

        _tasks[taskId] = Task(owner, tokens, amounts, msg.value, expiration);
    }

    function cancelTask(uint256 taskId) external override {
        if (msg.sender != _tasks[taskId].owner) {
            revert();
        }

        _payoutTask(taskId, _tasks[taskId].owner);
    }

    function endTask(uint256 taskId, uint256 githubId) external override {
        if (msg.sender != _tasks[taskId].owner) {
            revert();
        }

        if (taskId >= _totalTasks) {
            revert();
        }

        if (block.timestamp < _tasks[taskId].expiration) {
            revert();
        }

        if (IGithubConnector(githubConnector).githubIdExists(githubId)) {
            address receiver = IWooblayUsers(wooblayUsers).ownerOf(
                IGithubConnector(githubConnector).getTokenId(githubId)
            );

            _payoutTask(taskId, receiver);
        } else {
            _approveTaskPayout(taskId, githubPay);

            IGithubPay(githubPay).pay(
                githubId,
                _tasks[taskId].tokens,
                _tasks[taskId].amounts
            );
        }

        _tasks[taskId].expiration = type(uint256).max;
    }

    function _payoutTask(uint256 taskId, address receiver) internal {
        for (uint256 i = 0; i < _tasks[taskId].tokens.length; i++) {
            IERC20(_tasks[taskId].tokens[i]).transfer(
                receiver,
                _tasks[taskId].amounts[i]
            );
        }

        if (_tasks[taskId].value > 0) {
            payable(receiver).sendValue(_tasks[taskId].value);
        }
    }

    function _approveTaskPayout(uint256 taskId, address spender) internal {
        for (uint256 i = 0; i < _tasks[taskId].tokens.length; i++) {
            IERC20(_tasks[taskId].tokens[i]).approve(
                spender,
                _tasks[taskId].amounts[i]
            );
        }
    }
}
