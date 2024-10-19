// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IFunctionsConsumers {
    function isSubscriptionPublic(
        uint64 subscriptionId
    ) external view returns (bool);
    function isSubscriptionConsumer(
        uint64 subscriptionId,
        address account
    ) external view returns (bool);

    function setSubscriptionConsumers(
        uint64 subscriptionId,
        address[] memory addresses,
        bool[] memory values
    ) external;
}
