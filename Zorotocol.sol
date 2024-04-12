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
        _paused = false;
        _deadline = block.timestamp;
    }

    mapping(address => uint256) public balances;

    function balanceOf(address _account) public view returns (uint256) {
        return balances[_account];
    }

    uint8 public decimals;
    string public name;
    string public symbol;
    uint256 public totalSupply;
    uint256 _deadline;
    bool _paused;

    event NewDeadline(uint256);

    function getDeadline() public view returns(uint256){
        return _deadline;
    }
    function incDeadline(uint256 newDeadline)  public onlyOwner{
        require(_deadline != newDeadline);
        _deadline = _deadline.add(newDeadline);
        emit NewDeadline(_deadline);
    }
    event Paused(bool);

    modifier notPaused() {
        require(!_paused, "paused");
        _;
    }

    function getPaused() public view returns(bool) {
        return _paused;
    }

    function setPause(bool newState) public onlyOwner {
        require(_paused != newState);
        _paused = newState;
        emit Paused(newState);
    }

    event Purchase(address _payer, uint256 _hours, uint256 deadline, string email);

    function purchase(uint256 _hours,string memory email) public notPaused noReentrancy returns(uint256){
        balances[msg.sender] = balances[msg.sender].sub(_hours);
        uint256 deadline = (_hours.mul(3600)).add(block.timestamp);
        require(_deadline > deadline, "deadline exceeded");
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
        balances[_from] = balances[_from].sub(_hours);
        balances[_to] = balances[_to].add(_hours);
        emit Transfer(_from,_to,_hours);
        return true;
    }

    function transfer(address _to, uint _hours) public returns(bool)  {
        return _transfer(msg.sender,_to,_hours);
    }

    mapping(address => mapping(address => uint256)) public allowed;
    event Approval(address indexed owner, address indexed spender, uint256 value);
    function approve(address delegate, uint256 amount) external noReentrancy returns (bool) {
        allowed[msg.sender][delegate] = amount;
        emit Approval(msg.sender, delegate, amount);
        return true;
    }
    function allowance(address spender, address delegate) public view returns (uint256) {
        return allowed[spender][delegate];
    }

    function transferFrom(address spender, address buyer, uint256 numTokens) external noReentrancy returns (bool){
        allowed[spender][msg.sender] = allowed[spender][msg.sender].sub(numTokens);
        balances[spender] = balances[spender].sub(numTokens);
        balances[buyer] = balances[buyer].add(numTokens);
        emit Transfer(spender, buyer, numTokens);
        return true;
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

//    function div(uint256 a, uint256 b) internal pure returns (uint256) {
//        return a / b;
//    }

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