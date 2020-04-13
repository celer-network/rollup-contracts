pragma solidity ^0.5.2;

import {ISidechainERC20} from "./ISidechainERC20.sol";
import {ERC20} from "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import {
    ERC20Detailed
} from "openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";


contract SidechainERC20 is ISidechainERC20, ERC20, ERC20Detailed, Ownable {
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

    // Default hash for EOA accounts returned by extcodehash
    bytes32 internal constant ACCOUNT_HASH = 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470;

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

        // TODO: Check validator signature

        if (!registeredAccounts[account]) {
            registeredAccounts[account] = true;
        }

        _mint(account, amount);

        uint256 oldNonce = nonces[account];
        nonces[account] = oldNonce.add(1);

        emit Deposit(mainchainToken, account, amount, oldNonce, signature);
    }

    function withdraw(address account, uint256 amount, bytes memory signature)
        public
    {
        require(amount > 0 && balanceOf(account) >= amount);

        _burn(account, amount);

        uint256 oldNonce = nonces[account];
        nonces[account] = oldNonce.add(1);

        emit Withdraw(mainchainToken, account, amount, oldNonce, signature);
    }

    function transfer(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory signature
    ) public returns (bool) {
        uint256 oldNonce = nonces[sender];
        if (!isContract(sender)) {
            bytes32 hash = keccak256(
                abi.encodePacked(sender, recipient, amount, oldNonce)
            );
            bytes32 prefixedHash = toEthSignedMessageHash(hash);
            require(
                recoverAddress(prefixedHash, signature) == sender,
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

    /**
     * Returns whether the target address is a contract
     * @dev This function will return false if invoked during the constructor of a contract.
     * @param _address address of the account to check
     * @return Whether the target address is a contract
     */
    function isContract(address _address) internal view returns (bool) {
        bytes32 codehash;

        // Currently there is no better way to check if there is a contract in an address
        // than to check the size of the code at that address or if it has a non-zero code hash or account hash
        assembly {
            codehash := extcodehash(_address)
        }
        return (codehash != 0x0 && codehash != ACCOUNT_HASH);
    }

    function recoverAddress(bytes32 dataHash, bytes memory sig)
        public
        pure
        returns (address result)
    {
        bytes32 r;
        bytes32 s;
        uint8 v;
        if (sig.length != 65) {
            return address(0x0);
        }
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := and(mload(add(sig, 65)), 255)
        }
        // https://github.com/ethereum/go-ethereum/issues/2053
        if (v < 27) {
            v += 27;
        }
        if (v != 27 && v != 28) {
            return address(0x0);
        }
        result = ecrecover(dataHash, v, r, s);
        require(result != address(0x0), "Error in ecrecover");
    }

    /**
     * toEthSignedMessageHash
     * @dev prefix a bytes32 value with "\x19Ethereum Signed Message:"
     * and hash the result
     */
    function toEthSignedMessageHash(bytes32 hash)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
            );
    }
}
