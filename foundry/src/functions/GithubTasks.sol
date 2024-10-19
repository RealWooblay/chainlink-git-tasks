// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/token/ERC20/IERC20.sol";
import {Address} from "@openzeppelin/utils/Address.sol";

import {FunctionsClient} from "@chainlink/functions/v1_0_0/FunctionsClient.sol";
import {FunctionsRequest} from "@chainlink/functions/v1_0_0/libraries/FunctionsRequest.sol";

import {FunctionsConsumers} from "./FunctionsConsumers.sol";
import {IWooblayUsers} from "../token/IWooblayUsers.sol";
import {IGithubConnector} from "../connector/IGithubConnector.sol";
import {IGithubPay} from "../connector/IGithubPay.sol";
import {IGithubTasks} from "./IGithubTasks.sol";

contract GithubTasks is FunctionsClient, FunctionsConsumers, IGithubTasks {
    using Address for address payable;

    struct Task {
        address owner;
        uint256 issueId;
        address[] tokens;
        uint256[] amounts;
        uint256 value;
        uint256 expiration;
        bytes32 requestId;
    }

    using FunctionsRequest for FunctionsRequest.Request;

    address public immutable wooblayUsers;
    address public immutable githubConnector;
    address public immutable githubPay;

    bytes32 public immutable donId;

    // Task ID => Task
    mapping(uint256 => Task) private _tasks;
    uint256 private _totalTasks;

    // Request ID => Task ID
    mapping(bytes32 => uint256) private _requestTaskIds;

    constructor(
        address users,
        address connector,
        address pay,
        address subscriptions,
        bytes32 don,
        address router
    ) FunctionsClient(router) FunctionsConsumers(subscriptions) {
        wooblayUsers = users;
        githubConnector = connector;
        githubPay = pay;

        donId = don;
    }

    function newTask(
        address owner,
        uint256 issueId,
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

        _tasks[taskId] = Task(
            owner,
            issueId,
            tokens,
            amounts,
            msg.value,
            expiration,
            ""
        );
    }

    function cancelTask(uint256 taskId) external override {
        if (msg.sender != _tasks[taskId].owner) {
            revert();
        }

        _payoutTask(taskId, _tasks[taskId].owner);
    }

    function endTask(
        uint256 taskId,
        bytes memory encryptedSecretsUrls,
        uint64 subscriptionId,
        uint32 gasLimit
    )
        external
        override
        onlySubscriptionConsumer(subscriptionId, msg.sender)
        returns (bytes32 requestId)
    {
        if (taskId >= _totalTasks) {
            revert();
        }

        if (_tasks[taskId].requestId != "") {
            revert();
        }

        if (block.timestamp < _tasks[taskId].expiration) {
            revert();
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

        _tasks[taskId].requestId = requestId;
        _requestTaskIds[requestId] = taskId;
    }

    function fulfillRequest(
        bytes32 requestId,
        bytes memory response,
        bytes memory err
    ) internal override {
        if (err.length > 0) {
            // Invalid
            return;
        }

        uint256 githubId = abi.decode(response, (uint256));

        if (IGithubConnector(githubConnector).githubIdExists(githubId)) {
            address receiver = IWooblayUsers(wooblayUsers).ownerOf(
                IGithubConnector(githubConnector).getTokenId(githubId)
            );

            _payoutTask(_requestTaskIds[requestId], receiver);
        } else {
            _approveTaskPayout(_requestTaskIds[requestId], githubPay);

            IGithubPay(githubPay).pay(
                githubId,
                _tasks[_requestTaskIds[requestId]].tokens,
                _tasks[_requestTaskIds[requestId]].amounts
            );
        }

        _tasks[_requestTaskIds[requestId]].expiration = type(uint256).max;
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
