pragma solidity ^0.5.2;

import {ERC20} from "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import {
    ERC20Detailed
} from "openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";


contract SidechainERC20 is ERC20, ERC20Detailed, Ownable {
    using SafeMath for uint256;

    event Deposit(
        address indexed mainchainToken,
        address indexed account,
        uint256 amount,
        uint256 nonce,
        bytes signature
    );

    event Withdraw(
        address indexed mainchainToken,
        address indexed account,
        uint256 amount,
        uint256 nonce,
        bytes signature
    );

    event Transfer(
        address indexed mainchainToken,
        address indexed sender,
        address indexed recipient,
        uint256 amount,
        uint256 nonce,
        bytes signature
    );

    address public mainchainToken;
    mapping(address => uint256) public nonces;

    constructor(
        address _mainchainToken,
        string memory _name,
        string memory _symbol,
        uint8 _decimals
    ) public ERC20Detailed(_name, _symbol, _decimals) {
        require(_mainchainToken != address(0x0));
        mainchainToken = _mainchainToken;
    }

    function deposit(address account, uint256 amount, bytes memory signature)
        public
    {
        require(amount > 0 && account != address(0x0));

        _mint(account, amount);

        uint256 oldNonce = nonces[account];
        nonces[account] = oldNonce.add(1);

        emit Deposit(mainchainToken, account, amount, oldNonce, signature);
    }

    function withdraw(uint256 amount, bytes memory signature) public {
        address account = msg.sender;
        require(amount > 0 && balanceOf(account) >= amount);

        _burn(account, amount);

        uint256 oldNonce = nonces[account];
        nonces[account] = oldNonce.add(1);

        emit Withdraw(mainchainToken, account, amount, oldNonce, signature);
    }

    function transfer(address recipient, uint256 amount, bytes memory signature)
        public
        returns (bool)
    {
        address sender = msg.sender;

        _transfer(sender, recipient, amount);

        uint256 oldNonce = nonces[sender];
        nonces[sender] = oldNonce.add(1);

        emit Transfer(
            mainchainToken,
            sender,
            recipient,
            amount,
            oldNonce,
            signature
        );
        return true;
    }

    function transfer(address, uint256) public returns (bool) {
        revert("Disabled feature");
    }

    function allowance(address, address) public view returns (uint256) {
        revert("Disabled feature");
    }

    function approve(address, uint256) public returns (bool) {
        revert("Disabled feature");
    }

    function transferFrom(address, address, uint256) public returns (bool) {
        revert("Disabled feature");
    }
}
