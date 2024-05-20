// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {ERC721} from "@openzeppelin/token/ERC721/ERC721.sol";
import {AccessControl} from "@openzeppelin/access/AccessControl.sol";

contract Contributors is ERC721, AccessControl {
    bytes32 public constant GITHUB_ID_ROLE = keccak256("GITHUB_ID_ROLE");
    bytes32 public constant BASE_URI_ROLE = keccak256("BASE_URI_ROLE");

    // Token ID => GitHub ID
    mapping(uint256 => uint256) private _githubIds;

    uint256 private _nextTokenId;

    string private _baseUri;

    constructor(
        address defaultAdmin
    ) ERC721("Contributors", "WBLY") {
        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
    }

    function safeMint(address to) external {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
    }

    function setGithubId(uint256 tokenId, uint256 githubId) external onlyRole(GITHUB_ID_ROLE) {
        _githubIds[tokenId] = githubId;
    }

    function setBaseURI(string memory baseUri) external onlyRole(BASE_URI_ROLE) {
        _baseUri = baseUri;
    }   

    function _baseURI() internal view override returns (string memory) {
        return _baseUri;
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
