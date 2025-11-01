// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/*
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
*/
contract  Voting is Ownable{
    
    mapping(address => uint256) votesRestore; // 存储候选人的得票数

    //为了遍历mapping votesRestore
    mapping(uint8 => address) indexAddr;

    uint8 index = 0;

    address[] public voteAddrs;

    constructor()
        Ownable(msg.sender)
    {}

    function vote(address to) external checkAddr(to){

        uint256 voteNum = votesRestore[to];
        voteNum++;
        votesRestore[to] = voteNum;

        if(!contains(voteAddrs,to)){
            voteAddrs.push(to);
        }

    }

    function getVotes(address voteAddr) external checkAddr(voteAddr) view returns (uint256){
        
        return votesRestore[voteAddr];
    }

    function resetVotes() external onlyOwner{
        for(uint i=0; i < voteAddrs.length;i++){
            votesRestore[voteAddrs[i]] = 0;
        }
    }

    modifier checkAddr(address addr){
        require(addr != address(0), "vote address incorecct");
        _;
    }

    function contains(address[] memory addrs, address voteaddr) public pure returns(bool) {
        uint length = addrs.length;
        for (uint i = 0; i < length; i++){
            if(addrs[i] == voteaddr){
                return true;
            }
        }

        return false;
    }
}