pragma solidity ^0.6.6;

import {ISidechainERC20} from "./ISidechainERC20.sol";
import {ERC20} from "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/utils/Address.sol";
import "openzeppelin-solidity/contracts/access/Ownable.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";


contract SidechainERC20 is ISidechainERC20, ERC20, Ownable {
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
    mapping(address => bool) public registeredAccounts;

    constructor(
        address _mainchainToken,
        string memory _name,
        string memory _symbol,
        uint8 _decimals
    ) public ERC20(_name, _symbol) {
        require(_mainchainToken != address(0x0));
        mainchainToken = _mainchainToken;
        _setupDecimals(_decimals);
    }

    function deposit(
        address account,
        uint256 amount,
        bytes memory signature
    ) public {
        require(amount > 0 && account != address(0x0));

        // TODO: Check validator signature

        if (!registeredAccounts[account]) {
            registeredAccounts[account] = true;
        }

        _mint(account, amount);

        uint256 oldNonce = nonces[account];
        nonces[account] = oldNonce.add(1);

        emit Deposit(mainchainToken, account, amount, oldNonce, signature);
    }

    function withdraw(
        address account,
        uint256 amount,
        bytes memory signature
    ) public {
        require(amount > 0 && balanceOf(account) >= amount);

        _burn(account, amount);

        uint256 oldNonce = nonces[account];
        nonces[account] = oldNonce.add(1);

        emit Withdraw(mainchainToken, account, amount, oldNonce, signature);
    }

    // prettier-ignore
    function transfer(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory signature
    ) public override returns (bool) {
        uint256 oldNonce = nonces[sender];
        if (!Address.isContract(sender)) {
            bytes32 hash = keccak256(
                abi.encodePacked(
                    sender,
                    recipient,
                    mainchainToken,
                    amount,
                    oldNonce
                )
            );
            bytes32 prefixedHash = ECDSA.toEthSignedMessageHash(hash);
            require(
                ECDSA.recover(prefixedHash, signature) == sender,
                "Wrong signature"
            );
        }

        _transfer(sender, recipient, amount);
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

    // prettier-ignore
    function transfer(address, uint256) public override returns (bool) {
        revert("Disabled feature");
    }

    // prettier-ignore
    function allowance(address, address) public override view returns (uint256) {
        revert("Disabled feature");
    }

    // prettier-ignore
    function approve(address, uint256) public override returns (bool) {
        revert("Disabled feature");
    }

    // prettier-ignore
    function transferFrom(
        address,
        address,
        uint256
    ) public override returns (bool) {
        revert("Disabled feature");
    }
}
