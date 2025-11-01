// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract MailBox {
    uint public totalLetters;

    struct Letter {
        string message;
        address sender;
    }

    Letter[] private letters;

    function write(string memory _message) public {
        totalLetters++;

        //letters.push(Letter(_message,msg.sender)); 
        letters.push(Letter({message: _message,sender: msg.sender}));
    }

    function getAllLetters() public view returns (Letter[] memory) {
        return letters;
    }
}