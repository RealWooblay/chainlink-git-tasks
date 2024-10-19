// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IFunctionsSubscriptions} from "@chainlink/functions/v1_0_0/interfaces/IFunctionsSubscriptions.sol";

import {IFunctionsConsumers} from "./IFunctionsConsumers.sol";

contract FunctionsConsumers is IFunctionsConsumers {
    address public immutable functionsSubscriptions;

    mapping(uint64 => mapping(address => bool)) private _consumers;

    constructor(address subscriptions) {
        functionsSubscriptions = subscriptions;
    }

    modifier onlySubscriptionConsumer(uint64 subscriptionId, address account) {
        if (!isSubscriptionConsumer(subscriptionId, account)) {
            revert();
        }
        _;
    }

    function isSubscriptionPublic(
        uint64 subscriptionId
    ) public view override returns (bool) {
        return _consumers[subscriptionId][address(0)] == true;
    }

    function isSubscriptionConsumer(
        uint64 subscriptionId,
        address account
    ) public view override returns (bool) {
        return
            isSubscriptionPublic(subscriptionId) ||
            _consumers[subscriptionId][account] == true;
    }

    function setSubscriptionConsumers(
        uint64 subscriptionId,
        address[] memory addresses,
        bool[] memory values
    ) external override {
        IFunctionsSubscriptions.Subscription
            memory subscription = IFunctionsSubscriptions(
                functionsSubscriptions
            ).getSubscription(subscriptionId);

        if (msg.sender != subscription.owner) {
            revert();
        }

        if (addresses.length != values.length) {
            revert();
        }

        for (uint256 i = 0; i < addresses.length; i++) {
            _consumers[subscriptionId][addresses[i]] = values[i];
        }
    }
}
