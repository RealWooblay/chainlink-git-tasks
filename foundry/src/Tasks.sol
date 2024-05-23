// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {FunctionsClient} from "@chainlink/functions/v1_0_0/FunctionsClient.sol";
import {FunctionsRequest} from "@chainlink/functions/v1_0_0/libraries/FunctionsRequest.sol";

contract Tasks is FunctionsClient {
    using FunctionsRequest for FunctionsRequest.Request;

    uint64 private _subscriptionId = 50;
    bytes32 private _donId =
        0x66756e2d626173652d7365706f6c69612d310000000000000000000000000000;
    uint8 private _donHostedSecretsSlotID;
    uint64 private _donHostedSecretsVersion;

    constructor(address router) FunctionsClient(router) {}

    function setdonHostedSecrets(uint8 slotID, uint64 version) external {
        _donHostedSecretsSlotID = slotID;
        _donHostedSecretsVersion = version;
    }

    function sendRequest(
        string memory source,
        string[] memory args,
        uint32 gasLimit
    ) external returns (bytes32 requestId) {
        FunctionsRequest.Request memory req;
        req.initializeRequestForInlineJavaScript(source);
        if (_donHostedSecretsVersion > 0) {
            req.addDONHostedSecrets(
                _donHostedSecretsSlotID,
                _donHostedSecretsVersion
            );
        }
        if (args.length > 0) req.setArgs(args);
        return
            _sendRequest(req.encodeCBOR(), _subscriptionId, gasLimit, _donId);
    }

    function fulfillRequest(
        bytes32 requestId,
        bytes memory response,
        bytes memory err
    ) internal override {}
}
