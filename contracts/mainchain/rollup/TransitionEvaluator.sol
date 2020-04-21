pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";

/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {RollupTokenRegistry} from "./RollupTokenRegistry.sol";


contract TransitionEvaluator {
    using SafeMath for uint256;

    bytes32 constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;
    // Transition Types
    uint256 constant TRANSITION_TYPE_INITIAL_DEPOSIT = 0;
    uint256 constant TRANSITION_TYPE_DEPOSIT = 1;
    uint256 constant TRANSITION_TYPE_WITHDRAW = 2;
    uint256 constant TRANSITION_TYPE_TRANSFER = 3;

    RollupTokenRegistry tokenRegistry;

    constructor(address _tokenRegistryAddress) public {
        tokenRegistry = RollupTokenRegistry(_tokenRegistryAddress);
    }

    function evaluateTransition(
        bytes calldata _transition,
        DataTypes.StorageSlot[] calldata _storageSlots
    ) external view returns (bytes32[] memory) {
        // Convert our inputs to memory
        bytes memory transition = _transition;


            DataTypes.StorageSlot[] memory storageSlots
         = new DataTypes.StorageSlot[](_storageSlots.length);
        // Direct copy not supported by Solidity yet
        for (uint256 i = 0; i < _storageSlots.length; i++) {
            uint256 slotIndex = _storageSlots[i].slotIndex;
            address account = _storageSlots[i].value.account;
            uint256[] memory balances = _storageSlots[i].value.balances;
            uint256[] memory nonces = _storageSlots[i].value.nonces;
            storageSlots[i] = DataTypes.StorageSlot(
                slotIndex,
                DataTypes.AccountInfo(account, balances, nonces)
            );
        }

        // Extract the transition type
        uint256 transitionType = extractTransitionType(transition);
        bytes32[] memory outputs;
        // Apply the transition and record the resulting storage slots
        if (transitionType == TRANSITION_TYPE_INITIAL_DEPOSIT) {

                DataTypes.InitialDepositTransition memory initialDeposit
             = decodeInitialDepositTransition(transition);


                DataTypes.AccountInfo memory updatedAccountInfo
             = applyInitialDepositTransition(initialDeposit, storageSlots[0]);
            outputs = new bytes32[](1);
            outputs[0] = getAccountInfoHash(updatedAccountInfo);
        } else if (transitionType == TRANSITION_TYPE_DEPOSIT) {

                DataTypes.DepositTransition memory deposit
             = decodeDepositTransition(transition);


                DataTypes.AccountInfo memory updatedAccountInfo
             = applyDepositTransition(deposit, storageSlots[0]);
            outputs = new bytes32[](1);
            outputs[0] = getAccountInfoHash(updatedAccountInfo);
        } else if (transitionType == TRANSITION_TYPE_WITHDRAW) {

                DataTypes.WithdrawTransition memory withdraw
             = decodeWithdrawTransition(transition);


                DataTypes.AccountInfo memory updatedAccountInfo
             = applyWithdrawTransition(withdraw, storageSlots[0]);
            outputs = new bytes32[](1);
            outputs[0] = getAccountInfoHash(updatedAccountInfo);
        } else if (transitionType == TRANSITION_TYPE_TRANSFER) {

                DataTypes.TransferTransition memory transfer
             = decodeTransferTransition(transition);


                DataTypes.AccountInfo[2] memory updatedAccountInfos
             = applyTransferTransition(
                transfer,
                [storageSlots[0], storageSlots[1]]
            );
            // Return the hash of both storage (leaf nodes to insert into the tree)
            outputs = new bytes32[](2);
            for (uint256 i = 0; i < updatedAccountInfos.length; i++) {
                outputs[i] = getAccountInfoHash(updatedAccountInfos[i]);
            }
        } else {
            revert("Transition type not recognized!");
        }
        return outputs;
    }

    function extractTransitionType(bytes memory _bytes)
        internal
        pure
        returns (uint256)
    {
        uint256 transitionType;

        assembly {
            transitionType := mload(add(_bytes, 0x20))
        }

        return transitionType;
    }

    /**
     * Return the access list for this transition.
     */
    function getTransitionStateRootAndAccessList(bytes calldata _rawTransition)
        external
        pure
        returns (bytes32, uint256[] memory)
    {
        // Initialize memory rawTransition
        bytes memory rawTransition = _rawTransition;
        // Initialize stateRoot and storageSlots
        bytes32 stateRoot;
        uint256[] memory storageSlots;
        uint256 transitionType = extractTransitionType(rawTransition);
        if (transitionType == TRANSITION_TYPE_INITIAL_DEPOSIT) {

                DataTypes.InitialDepositTransition memory transition
             = decodeInitialDepositTransition(rawTransition);
            stateRoot = transition.stateRoot;
            storageSlots = new uint256[](1);
            storageSlots[0] = transition.accountSlotIndex;
        } else if (transitionType == TRANSITION_TYPE_DEPOSIT) {

                DataTypes.DepositTransition memory transition
             = decodeDepositTransition(rawTransition);
            stateRoot = transition.stateRoot;
            storageSlots = new uint256[](1);
            storageSlots[0] = transition.accountSlotIndex;
        } else if (transitionType == TRANSITION_TYPE_WITHDRAW) {

                DataTypes.WithdrawTransition memory transition
             = decodeWithdrawTransition(rawTransition);
            stateRoot = transition.stateRoot;
            storageSlots = new uint256[](1);
            storageSlots[0] = transition.accountSlotIndex;
        } else if (transitionType == TRANSITION_TYPE_TRANSFER) {

                DataTypes.TransferTransition memory transition
             = decodeTransferTransition(rawTransition);
            stateRoot = transition.stateRoot;
            storageSlots = new uint256[](2);
            storageSlots[0] = transition.senderSlotIndex;
            storageSlots[1] = transition.recipientSlotIndex;
        }
        return (stateRoot, storageSlots);
    }

    function getWithdrawTxHash(DataTypes.WithdrawTx memory _withdrawTx)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encodePacked(
                    _withdrawTx.account,
                    _withdrawTx.token,
                    _withdrawTx.amount
                )
            );
    }

    function getTransferTxHash(DataTypes.TransferTx memory _transferTx)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encodePacked(
                    _transferTx.sender,
                    _transferTx.recipient,
                    _transferTx.token,
                    _transferTx.amount,
                    _transferTx.nonce
                )
            );
    }

    function verifyEmptyAccountInfo(DataTypes.AccountInfo memory _accountInfo)
        internal
        pure
    {
        require(
            _accountInfo.account == 0x0000000000000000000000000000000000000000,
            "Address of empty account must be zero"
        );
        require(
            _accountInfo.balances.length == 0,
            "Balance array must be empty"
        );
        require(_accountInfo.nonces.length == 0, "Nonce array must be empty");
    }

    /**
     * Apply an InitialDepositTransition.
     */
    function applyInitialDepositTransition(
        DataTypes.InitialDepositTransition memory _transition,
        DataTypes.StorageSlot memory _storageSlot
    ) public view returns (DataTypes.AccountInfo memory) {
        // Verify that the AccountInfo is empty
        verifyEmptyAccountInfo(_storageSlot.value);
        // Now set storage slot to have the address of the registered account
        _storageSlot.value.account = _transition.account;
        // Next create a DepositTransition based on this InitialDepositTransition
        DataTypes.DepositTransition memory depositTransition = DataTypes
            .DepositTransition(
            TRANSITION_TYPE_DEPOSIT,
            _transition.stateRoot,
            _transition.accountSlotIndex,
            _transition.tokenIndex,
            _transition.amount,
            _transition.signature
        );
        // Now simply apply the deposit transition as usual
        return applyDepositTransition(depositTransition, _storageSlot);
    }

    /**
     * Apply a DepositTransition.
     */
    function applyDepositTransition(
        DataTypes.DepositTransition memory _transition,
        DataTypes.StorageSlot memory _storageSlot
    ) public view returns (DataTypes.AccountInfo memory) {
        address account = _storageSlot.value.account;

        DataTypes.DepositTx memory depositTx = DataTypes.DepositTx(
            account,
            tokenRegistry.tokenIndexToTokenAddress(_transition.tokenIndex),
            _transition.amount
        );

        // TODO (dominator008): Verify signature of depositer

        DataTypes.AccountInfo memory outputStorage;
        uint256 tokenIndex = _transition.tokenIndex;
        uint256 oldBalance = _storageSlot.value.balances[tokenIndex];
        _storageSlot.value.balances[tokenIndex] = oldBalance.add(
            depositTx.amount
        );
        outputStorage = _storageSlot.value;
        return outputStorage;
    }

    /**
     * Apply a WithdrawTransition.
     */
    function applyWithdrawTransition(
        DataTypes.WithdrawTransition memory _transition,
        DataTypes.StorageSlot memory _storageSlot
    ) public view returns (DataTypes.AccountInfo memory) {
        address account = _storageSlot.value.account;

        DataTypes.WithdrawTx memory withdrawTx = DataTypes.WithdrawTx(
            account,
            tokenRegistry.tokenIndexToTokenAddress(_transition.tokenIndex),
            _transition.amount
        );

        bytes32 txHash = getWithdrawTxHash(withdrawTx);
        bytes32 prefixedHash = ECDSA.toEthSignedMessageHash(txHash);
        require(
            ECDSA.recover(prefixedHash, _transition.signature) == account,
            "Withdraw signature is invalid!"
        );
        DataTypes.AccountInfo memory outputStorage;
        uint256 tokenIndex = _transition.tokenIndex;
        uint256 oldBalance = _storageSlot.value.balances[tokenIndex];
        _storageSlot.value.balances[tokenIndex] = oldBalance.sub(
            withdrawTx.amount
        );
        outputStorage = _storageSlot.value;
        return outputStorage;
    }

    /**
     * Apply a TransferTransition.
     */
    function applyTransferTransition(
        DataTypes.TransferTransition memory _transition,
        DataTypes.StorageSlot[2] memory _storageSlots
    ) public view returns (DataTypes.AccountInfo[2] memory) {
        // First construct the transaction from the storage slots
        address sender = _storageSlots[0].value.account;
        address recipient = _storageSlots[1].value.account;
        DataTypes.TransferTx memory transferTx = DataTypes.TransferTx(
            sender,
            recipient,
            tokenRegistry.tokenIndexToTokenAddress(_transition.tokenIndex),
            _transition.amount,
            _transition.nonce
        );

        // Next check to see if the signature is valid
        bytes32 txHash = getTransferTxHash(transferTx);
        bytes32 prefixedHash = ECDSA.toEthSignedMessageHash(txHash);
        require(
            ECDSA.recover(prefixedHash, _transition.signature) == sender,
            "Transfer signature is invalid!"
        );

        // Create an array to store our output storage slots
        DataTypes.AccountInfo[2] memory outputStorage;
        // Now we know the signature is correct, let's compute the output of the transaction
        uint256 tokenIndex = _transition.tokenIndex;
        uint256 senderBalance = _storageSlots[0].value.balances[tokenIndex];

        // First let's make sure the sender has enough money
        require(
            senderBalance > transferTx.amount,
            "Sender does not have enough money!"
        );

        // Update the storage slots with the new balances
        uint256 senderOldBalance = _storageSlots[0].value.balances[tokenIndex];
        uint256 recipientOldBalance = _storageSlots[1]
            .value
            .balances[tokenIndex];
        _storageSlots[0].value.balances[tokenIndex] = senderOldBalance.sub(
            transferTx.amount
        );
        _storageSlots[1].value.balances[tokenIndex] = recipientOldBalance.add(
            transferTx.amount
        );
        // Set the outputs
        outputStorage[0] = _storageSlots[0].value;
        outputStorage[1] = _storageSlots[1].value;
        // Return the outputs!
        return outputStorage;
    }

    /**
     * Get the hash of the AccountInfo.
     */
    function getAccountInfoHash(DataTypes.AccountInfo memory _accountInfo)
        public
        pure
        returns (bytes32)
    {
        // Here we don't use `abi.encode([struct])` because it's not clear
        // how to generate that encoding client-side.
        return
            keccak256(
                abi.encode(
                    _accountInfo.account,
                    _accountInfo.balances,
                    _accountInfo.nonces
                )
            );
    }

    /************
     * Decoding *
     ***********/

    /**
     * Decode an InitialDepositTransition
     * TODO: Decode directly into a struct.
     */
    function decodeInitialDepositTransition(bytes memory _rawBytes)
        internal
        pure
        returns (DataTypes.InitialDepositTransition memory)
    {
        (
            uint256 transitionType,
            bytes32 stateRoot,
            uint256 accountSlotIndex,
            address account,
            uint256 tokenIndex,
            uint256 amount,
            bytes memory signature
        ) = abi.decode(
            (_rawBytes),
            (uint256, bytes32, uint256, address, uint256, uint256, bytes)
        );
        DataTypes.InitialDepositTransition memory transition = DataTypes
            .InitialDepositTransition(
            transitionType,
            stateRoot,
            accountSlotIndex,
            account,
            tokenIndex,
            amount,
            signature
        );
        return transition;
    }

    function decodeDepositTransition(bytes memory _rawBytes)
        internal
        pure
        returns (DataTypes.DepositTransition memory)
    {
        (
            uint256 transitionType,
            bytes32 stateRoot,
            uint256 accountSlotIndex,
            uint256 tokenIndex,
            uint256 amount,
            bytes memory signature
        ) = abi.decode(
            (_rawBytes),
            (uint256, bytes32, uint256, uint256, uint256, bytes)
        );
        DataTypes.DepositTransition memory transition = DataTypes
            .DepositTransition(
            transitionType,
            stateRoot,
            accountSlotIndex,
            tokenIndex,
            amount,
            signature
        );
        return transition;
    }

    function decodeWithdrawTransition(bytes memory _rawBytes)
        public
        pure
        returns (DataTypes.WithdrawTransition memory)
    {
        (
            uint256 transitionType,
            bytes32 stateRoot,
            uint256 accountSlotIndex,
            uint256 tokenIndex,
            uint256 amount,
            bytes memory signature
        ) = abi.decode(
            (_rawBytes),
            (uint256, bytes32, uint256, uint256, uint256, bytes)
        );
        DataTypes.WithdrawTransition memory transition = DataTypes
            .WithdrawTransition(
            transitionType,
            stateRoot,
            accountSlotIndex,
            tokenIndex,
            amount,
            signature
        );
        return transition;
    }

    /**
     * Decode a TransferTransition
     * TODO: Decode directly into a struct.
     */
    function decodeTransferTransition(bytes memory _rawBytes)
        internal
        pure
        returns (DataTypes.TransferTransition memory)
    {
        (
            uint256 transitionType,
            bytes32 stateRoot,
            uint256 senderSlotIndex,
            uint256 recipientSlotIndex,
            uint256 tokenIndex,
            uint256 amount,
            uint256 nonce,
            bytes memory signature
        ) = abi.decode(
            (_rawBytes),
            (
                uint256,
                bytes32,
                uint256,
                uint256,
                uint256,
                uint256,
                uint256,
                bytes
            )
        );
        DataTypes.TransferTransition memory transition = DataTypes
            .TransferTransition(
            transitionType,
            stateRoot,
            senderSlotIndex,
            recipientSlotIndex,
            tokenIndex,
            amount,
            nonce,
            signature
        );
        return transition;
    }
}
