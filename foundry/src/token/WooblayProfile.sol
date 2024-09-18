// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {ERC721} from "@openzeppelin/token/ERC721/ERC721.sol";
import {ERC721Votes} from "@openzeppelin/token/ERC721/extensions/ERC721Votes.sol";
import {EIP712} from "@openzeppelin/utils/cryptography/EIP712.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";

import {IWooblayProfile} from "./IWooblayProfile.sol";

contract WooblayProfile is
    ERC721,
    EIP712,
    ERC721Votes,
    Ownable,
    IWooblayProfile
{
    uint256 private _totalSupply;

    string private _baseUri;

    constructor(address initialOwner)
        ERC721("Wooblay Profile", "WOOBLAYPROFILE")
        EIP712("Wooblay Profile", "1")
        Ownable(initialOwner)
    {}

    function tokenExists(
        uint256 tokenId
    ) external view override returns (bool) {
        return tokenId < _totalSupply;
    }

    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    function safeMint(address to) external override returns (uint256 tokenId) {
        tokenId = _totalSupply;
        _safeMint(to, tokenId);

        _totalSupply++;
    }

    function setBaseURI(string memory baseUri) external override onlyOwner {
        _baseUri = baseUri;
    }

    function _baseURI() internal view override returns (string memory) {
        return _baseUri;
    }

    function _update(
        address to,
        uint256 tokenId,
        address auth
    ) internal override(ERC721, ERC721Votes) returns (address) {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(
        address account,
        uint128 value
    ) internal override(ERC721, ERC721Votes) {
        super._increaseBalance(account, value);
    }
}
