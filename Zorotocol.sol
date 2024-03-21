pragma solidity ^0.8.0;
// SPDX-License-Identifier: MIT

contract Zorotocol {
    using SafeMath for uint256;
    receive() external payable {
        require(false, "no receive");
    }
    fallback() external payable {
        require(false, "no fallback");
    }
    bool private _reentrancy = false;
    modifier noReentrancy() {
        require(!_reentrancy, "no reentrancy allowed");
        _reentrancy = true;
        _;
        _reentrancy = false;
    }
    address public owner;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    function transferOwnership(address _newOwner) public onlyOwner {
        if (_newOwner != address(0)) {
            owner = _newOwner;
        }
    }

    constructor(string memory _name,string memory _symbol){
        owner = msg.sender;
        name = _name;
        symbol = _symbol;
        totalSupply = 0;
        decimals = 0;
    }

    mapping(address => uint256) public balances;

    function balanceOf(address _account) public view returns (uint256) {
        return balances[_account];
    }

    uint8 public decimals;
    string public name;
    string public symbol;
    uint256 public totalSupply;

    event Purchase(address _payer, uint256 _hours, uint256 deadline, string email);

    function purchase(uint256 _hours,string memory email) public noReentrancy returns(uint256){
        balances[msg.sender] = balances[msg.sender].sub(_hours);
        uint256 deadline = (_hours.mul(3600)).add(block.timestamp);
        totalSupply = totalSupply.sub(_hours);
        emit Purchase(msg.sender,_hours,deadline,email);
        return deadline;
    }

    event Issue(address _account,uint256 _hours);
    function issue(address _account,uint256 _hours) public onlyOwner {
        balances[_account] = balances[_account].add(_hours);
        totalSupply = totalSupply.add(_hours);
        emit Issue(_account,_hours);
    }

    event Transfer(address, address, uint256);
    function _transfer(address _from, address _to, uint _hours) internal noReentrancy returns(bool){
        require(_to != address(0), "invalid receiver address");
        if (balances[_from] < _hours){
            return false;
        }
        balances[_from] = balances[_from].sub(_hours);
        balances[_to] = balances[_to].add(_hours);
        emit Transfer(_from,_to,_hours);
        return true;
    }

    function transfer(address _to, uint _hours) public returns(bool)  {
        return _transfer(msg.sender,_to,_hours);
    }
}



library SafeMath {
    function mul(uint256 a, uint256 b) internal pure returns (uint256 c) {
        if (a == 0) {
            return 0;
        }
        c = a * b;
        assert(c / a == b);
        return c;
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return a / b;
    }

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        assert(b <= a);
        return a - b;
    }


    function add(uint256 a, uint256 b) internal pure returns (uint256 c) {
        c = a + b;
        assert(c >= a);
        return c;
    }
}