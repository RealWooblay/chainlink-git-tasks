// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {FunctionsClient} from "@chainlink/functions/v1_0_0/FunctionsClient.sol";

import {FunctionsConsumers} from "../functions/FunctionsConsumers.sol";

abstract contract Connector is FunctionsClient, FunctionsConsumers {
    bytes32 public immutable donId;

    constructor(
        address subscriptions,
        bytes32 don,
        address router
    ) FunctionsClient(router) FunctionsConsumers(subscriptions) {
        donId = don;
    }
}
