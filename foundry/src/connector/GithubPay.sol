// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IERC20} from "@openzeppelin/token/ERC20/IERC20.sol";
import {Address} from "@openzeppelin/utils/Address.sol";

import {IWooblayUsers} from "../token/IWooblayUsers.sol";
import {IGithubConnector} from "./IGithubConnector.sol";
import {IGithubPay} from "./IGithubPay.sol";

contract GithubPay is IGithubPay {
    using Address for address payable;

    address public immutable wooblayUsers;
    address public immutable githubConnector;

    address public constant ETH_ADDRESS = address(0);

    // Github ID => Token => Balance
    mapping(uint256 => mapping(address => uint256)) private _balances;

    constructor(address users, address githubCon) {
        wooblayUsers = users;
        githubConnector = githubCon;
    }

    function pay(
        uint256 githubId,
        address[] memory tokens,
        uint256[] memory amounts
    ) external payable override {
        if (tokens.length != amounts.length) {
            revert();
        }

        if (IGithubConnector(githubConnector).githubIdExists(githubId)) {
            address receiver = IWooblayUsers(wooblayUsers).ownerOf(
                IGithubConnector(githubConnector).getTokenId(githubId)
            );

            for (uint256 i = 0; i < tokens.length; i++) {
                IERC20(tokens[i]).transferFrom(
                    msg.sender,
                    receiver,
                    amounts[i]
                );
            }

            if (msg.value > 0) {
                payable(receiver).sendValue(msg.value);
            }
        } else {
            for (uint256 i = 0; i < tokens.length; i++) {
                IERC20(tokens[i]).transferFrom(
                    msg.sender,
                    address(this),
                    amounts[i]
                );
                _balances[githubId][tokens[i]] += amounts[i];
            }

            if (msg.value > 0) {
                _balances[githubId][ETH_ADDRESS] += msg.value;
            }
        }
    }

    function claim(
        uint256 githubId,
        address[] memory tokens
    ) external override {
        if (!IGithubConnector(githubConnector).githubIdExists(githubId)) {
            revert();
        }

        address receiver = IWooblayUsers(wooblayUsers).ownerOf(
            IGithubConnector(githubConnector).getTokenId(githubId)
        );

        for (uint256 i = 0; i < tokens.length; i++) {
            if (tokens[i] == ETH_ADDRESS) {
                payable(receiver).sendValue(_balances[githubId][tokens[i]]);
            } else {
                IERC20(tokens[i]).transfer(
                    receiver,
                    _balances[githubId][tokens[i]]
                );
            }

            _balances[githubId][tokens[i]] = 0;
        }
    }
}
