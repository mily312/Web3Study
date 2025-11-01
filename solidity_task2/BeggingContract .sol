// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/*
一个 mapping 来记录每个捐赠者的捐赠金额。
一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
一个 withdraw 函数，允许合约所有者提取所有资金。
一个 getDonation 函数，允许查询某个地址的捐赠金额。
使用 payable 修饰符和 address.transfer 实现支付和提款。
*/

contract BeggingContract is Ownable{

    mapping(address => uint256) public donations; // 记录每个捐赠者的捐赠金额

    event donateEvent(address from, uint256 amount);

     constructor()
        Ownable(msg.sender)
    {}

    function donate() external payable {
        uint256 donationAmount = donations[msg.sender];
        donations[msg.sender] = donationAmount + msg.value;
        emit donateEvent(msg.sender, msg.value);
    }

    function withdraw() external payable onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    function getDonation(address addr) external view returns (uint256) {
        return donations[addr];
    }

    receive() external payable { 
        
    }

}