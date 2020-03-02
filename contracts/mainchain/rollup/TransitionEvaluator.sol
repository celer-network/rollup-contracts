pragma solidity ^0.5.2;
pragma experimental ABIEncoderV2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";

/* Internal Imports */
import {DataTypes as dt} from "./DataTypes.sol";
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
        dt.StorageSlot[] calldata _storageSlots
    ) external view returns (bytes32[] memory) {
        // Convert our inputs to memory
        bytes memory transition = _transition;
        dt.StorageSlot[] memory storageSlots = new dt.StorageSlot[](
            _storageSlots.length
        );
        // Direct copy not supported by Solidity yet
        for (uint256 i = 0; i < _storageSlots.length; i++) {
            uint256 slotIndex = _storageSlots[i].slotIndex;
            address account = _storageSlots[i].value.account;
            uint256[] memory balances = _storageSlots[i].value.balances;
            uint256[] memory nonces = _storageSlots[i].value.nonces;
            storageSlots[i] = dt.StorageSlot(
                slotIndex,
                dt.AccountInfo(account, balances, nonces)
            );
        }
        // Extract the transition type
        uint256 transitionType = extractTransitionType(transition);
        bytes32[] memory outputs;
        // Apply the transition and record the resulting storage slots
        if (transitionType == TRANSITION_TYPE_INITIAL_DEPOSIT) {

                dt.InitialDepositTransition memory initialDeposit
             = decodeInitialDepositTransition(transition);


                dt.AccountInfo memory updatedAccountInfo
             = applyInitialDepositTransition(initialDeposit, storageSlots[0]);
            outputs = new bytes32[](1);
            outputs[0] = getAccountInfoHash(updatedAccountInfo);
        } else if (transitionType == TRANSITION_TYPE_DEPOSIT) {
            dt.DepositTransition memory deposit = decodeDepositTransition(
                transition
            );
            dt.AccountInfo memory updatedAccountInfo = applyDepositTransition(
                deposit,
                storageSlots[0]
            );
            outputs = new bytes32[](1);
            outputs[0] = getAccountInfoHash(updatedAccountInfo);
        } else if (transitionType == TRANSITION_TYPE_WITHDRAW) {
            dt.WithdrawTransition memory withdraw = decodeWithdrawTransition(
                transition
            );
            dt.AccountInfo memory updatedAccountInfo = applyWithdrawTransition(
                withdraw,
                storageSlots[0]
            );
            outputs = new bytes32[](1);
            outputs[0] = getAccountInfoHash(updatedAccountInfo);
        } else if (transitionType == TRANSITION_TYPE_TRANSFER) {
            dt.TransferTransition memory transfer = decodeTransferTransition(
                transition
            );


                dt.AccountInfo[2] memory updatedAccountInfos
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

                dt.InitialDepositTransition memory transition
             = decodeInitialDepositTransition(rawTransition);
            stateRoot = transition.stateRoot;
            storageSlots[0] = transition.accountSlotIndex;
        } else if (transitionType == TRANSITION_TYPE_DEPOSIT) {
            dt.DepositTransition memory transition = decodeDepositTransition(
                rawTransition
            );
            stateRoot = transition.stateRoot;
            storageSlots[0] = transition.accountSlotIndex;
        } else if (transitionType == TRANSITION_TYPE_WITHDRAW) {
            dt.WithdrawTransition memory transition = decodeWithdrawTransition(
                rawTransition
            );
            stateRoot = transition.stateRoot;
            storageSlots[0] = transition.accountSlotIndex;
        } else if (transitionType == TRANSITION_TYPE_TRANSFER) {
            dt.TransferTransition memory transition = decodeTransferTransition(
                rawTransition
            );
            stateRoot = transition.stateRoot;
            storageSlots[0] = transition.senderSlotIndex;
            storageSlots[1] = transition.recipientSlotIndex;
        }
        return (stateRoot, storageSlots);
    }

    function getWithdrawTxHash(dt.WithdrawTx memory _withdrawTx)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encode(
                    _withdrawTx.account,
                    _withdrawTx.token,
                    _withdrawTx.amount
                )
            );
    }

    function getTransferTxHash(dt.TransferTx memory _transferTx)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encode(
                    _transferTx.sender,
                    _transferTx.recipient,
                    _transferTx.token,
                    _transferTx.amount,
                    _transferTx.nonce
                )
            );
    }

    function verifyEmptyAccountInfo(dt.AccountInfo memory _accountInfo)
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
        dt.InitialDepositTransition memory _transition,
        dt.StorageSlot memory _storageSlot
    ) public view returns (dt.AccountInfo memory) {
        // Verify that the AccountInfo is empty
        verifyEmptyAccountInfo(_storageSlot.value);
        // Now set storage slot to have the address of the registered account
        _storageSlot.value.account = _transition.account;
        // Next create a DepositTransition based on this InitialDepositTransition
        dt.DepositTransition memory depositTransition = dt.DepositTransition(
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
        dt.DepositTransition memory _transition,
        dt.StorageSlot memory _storageSlot
    ) public view returns (dt.AccountInfo memory) {
        address account = _storageSlot.value.account;

        dt.DepositTx memory depositTx = dt.DepositTx(
            account,
            tokenRegistry.tokenIndexToTokenAddress(_transition.tokenIndex),
            _transition.amount
        );

        // TODO (dominator008): Verify signature of depositer

        dt.AccountInfo memory outputStorage;
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
        dt.WithdrawTransition memory _transition,
        dt.StorageSlot memory _storageSlot
    ) public view returns (dt.AccountInfo memory) {
        address account = _storageSlot.value.account;

        dt.WithdrawTx memory withdrawTx = dt.WithdrawTx(
            account,
            tokenRegistry.tokenIndexToTokenAddress(_transition.tokenIndex),
            _transition.amount
        );

        require(
            verifyEcdsaSignatureOnHash(
                _transition.signature,
                getWithdrawTxHash(withdrawTx),
                account
            ),
            "Transfer signature is invalid!"
        );
        dt.AccountInfo memory outputStorage;
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
        dt.TransferTransition memory _transition,
        dt.StorageSlot[2] memory _storageSlots
    ) public view returns (dt.AccountInfo[2] memory) {
        // First construct the transaction from the storage slots
        address sender = _storageSlots[0].value.account;
        address recipient = _storageSlots[1].value.account;
        dt.TransferTx memory transferTx = dt.TransferTx(
            sender,
            recipient,
            tokenRegistry.tokenIndexToTokenAddress(_transition.tokenIndex),
            _transition.amount,
            _transition.nonce
        );

        // Next check to see if the signature is valid
        require(
            verifyEcdsaSignatureOnHash(
                _transition.signature,
                getTransferTxHash(transferTx),
                sender
            ),
            "Transfer signature is invalid!"
        );

        // Create an array to store our output storage slots
        dt.AccountInfo[2] memory outputStorage;
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
    function getAccountInfoHash(dt.AccountInfo memory _accountInfo)
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
        returns (dt.InitialDepositTransition memory)
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
        dt.InitialDepositTransition memory transition = dt
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
        returns (dt.DepositTransition memory)
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
        dt.DepositTransition memory transition = dt.DepositTransition(
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
        returns (dt.WithdrawTransition memory)
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
        dt.WithdrawTransition memory transition = dt.WithdrawTransition(
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
        returns (dt.TransferTransition memory)
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
        dt.TransferTransition memory transition = dt.TransferTransition(
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

    // splits a signature string into v, r, s
    function splitSignature(bytes memory sig)
        internal
        pure
        returns (uint8 v, bytes32 r, bytes32 s)
    {
        require(sig.length == 65, "invalid signature length.");

        assembly {
            // first 32 bytes, after the length prefix.
            r := mload(add(sig, 32))
            // second 32 bytes.
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes).
            v := byte(0, mload(add(sig, 96)))
        }

        return (v, r, s);
    }

    // verifies a signature on a 32 byte value -- note this means we must be
    // signing / verifying the hash of our transactions, not the encodings
    // themselves -- luckily, DefaultSignatureProvider now does this.
    function verifyEcdsaSignatureOnHash(
        bytes memory _signature,
        bytes32 _hash,
        address _pubkey
    ) private pure returns (bool) {
        bytes memory prefixedMessage = abi.encodePacked(
            "\x19Ethereum Signed Message:\n32",
            _hash
        );
        bytes32 digest = keccak256(prefixedMessage);
        (uint8 v, bytes32 r, bytes32 s) = splitSignature(_signature);
        return ecrecover(digest, v, r, s) == _pubkey;
    }
}
