// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IERC721Metadata} from "@openzeppelin/token/ERC721/extensions/IERC721Metadata.sol";
import {IERC5805} from "@openzeppelin/interfaces/IERC5805.sol";

interface IWooblayProfile is IERC721Metadata, IERC5805 {
    function tokenExists(uint256 tokenId) external view returns (bool);
    function totalSupply() external view returns (uint256);

    function safeMint(address to) external returns (uint256 tokenId);

    function setBaseURI(string memory baseUri) external;
}