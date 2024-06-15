// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Storage {
    uint256 value;

    function store(uint256 number) public{
        value = number;
    }

    function retrieve() public view returns (uint256){
        return value;
    }
}