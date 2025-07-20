// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.2 <0.9.0;

contract MessageStorage {
    mapping(uint256 => string) private messages;
    uint256 private messageIdCounter;

    function setMessage(string memory _message) public returns (uint256) {
        messageIdCounter++;
        messages[messageIdCounter] = _message;
        return messageIdCounter;
    }

    function getMessage(uint256 _id) public view returns (string memory) {
        require(_id > 0 && _id <= messageIdCounter, "Invalid message ID");
        return messages[_id];
    }

    function getTotalMessages() public view returns (uint256) {
        return messageIdCounter;
    }
}