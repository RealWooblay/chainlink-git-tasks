// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

interface IGHTasks {
    event NewOrg(uint256 indexed orgId, address[] initalOwners);
    event SetOrgOwners(uint256 indexed orgId, address[] owners, bool[] values);
    event SetOrgRepo(
        uint256 indexed orgId,
        string repoOwner,
        string repoName,
        bool value
    );

    event SetSource(uint256 indexed sourceId, string source);
    event Request(
        uint256 indexed orgId,
        uint256 indexed sourceId,
        string[] args,
        bytes32 requestId
    );
    event RequestFulfilled(bytes32 requestId, bytes response, bytes error);

    struct Org {
        mapping(address => bool) owners;
        mapping(string => mapping(string => bool)) repos;
    }

    function newOrg(
        address[] memory initalOwners
    ) external returns (uint256 orgId);

    function setOrgOwners(
        uint256 orgId,
        address[] memory owners,
        bool[] memory values
    ) external;

    function setOrgRepo(
        uint256 orgId,
        string memory repoOwner,
        string memory repoName,
        bool value
    ) external;

    function setSource(uint256 sourceId, string memory source) external;

    function request(
        uint256 orgId,
        uint256 sourceId,
        string[] memory args,
        uint32 gasLimit
    ) external returns (bytes32 requestId);
}
