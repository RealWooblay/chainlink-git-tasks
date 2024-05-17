// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {FunctionsClient} from "@chainlink/functions/v1_0_0/FunctionsClient.sol";
import {ConfirmedOwner} from "@chainlink/shared/access/ConfirmedOwner.sol";
import {FunctionsRequest} from "@chainlink/functions/v1_0_0/libraries/FunctionsRequest.sol";

contract Tasks is FunctionsClient, ConfirmedOwner {
    using FunctionsRequest for FunctionsRequest.Request;

    constructor(
        address router
    ) FunctionsClient(router) ConfirmedOwner(msg.sender) {}

    function fulfillRequest(
        bytes32 requestId,
        bytes memory response,
        bytes memory err
    ) internal override {}
}
