pragma solidity ^0.5.2;
pragma experimental ABIEncoderV2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import {IERC20} from "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

/* Internal Imports */
import {DataTypes as dt} from "./DataTypes.sol";
import {RollupMerkleUtils} from "./RollupMerkleUtils.sol";
import {TransitionEvaluator} from "./TransitionEvaluator.sol";
import {RollupTokenRegistry} from "./RollupTokenRegistry.sol";


contract RollupChain {
    using SafeMath for uint256;

    /* Fields */
    // The state transition evaluator
    TransitionEvaluator transitionEvaluator;
    // The Rollup Merkle Tree library (currently a contract for ease of testing)
    RollupMerkleUtils merkleUtils;
    // The token registry
    RollupTokenRegistry tokenRegistry;
    // All the blocks!
    dt.Block[] public blocks;
    bytes32 public constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;
    // State tree height
    uint256 constant STATE_TREE_HEIGHT = 32;

    address aggregatorAddress;

    /* Events */
    event DecodedTransition(bool success, bytes returnData);
    event NewRollupBlock(bytes[] block, uint256 blockNumber);
    event Deposit(address account, address token, uint256 amount);

    /***************
     * Constructor *
     **************/
    constructor(
        address _transitionEvaluatorAddress,
        address _rollupMerkleUtilsAddress,
        address _tokenRegistryAddress,
        address _aggregatorAddress
    ) public {
        transitionEvaluator = TransitionEvaluator(_transitionEvaluatorAddress);
        merkleUtils = RollupMerkleUtils(_rollupMerkleUtilsAddress);
        tokenRegistry = RollupTokenRegistry(_tokenRegistryAddress);
        aggregatorAddress = _aggregatorAddress;
    }

    /* Methods */
    function pruneBlocksAfter(uint256 _blockNumber) internal {
        for (uint256 i = _blockNumber; i < blocks.length; i++) {
            delete blocks[i];
        }
    }

    /**
     * Submits a new block which is then rolled up.
     */
    function submitBlock(bytes[] calldata _block) external returns (bytes32) {
        require(
            msg.sender == aggregatorAddress,
            "Only the aggregator may submit blocks"
        );

        bytes32 root = merkleUtils.getMerkleRoot(_block);
        dt.Block memory rollupBlock = dt.Block({
            rootHash: root,
            blockSize: _block.length
        });
        blocks.push(rollupBlock);
        // NOTE: Toggle the event if you'd like easier historical block queries
        emit NewRollupBlock(_block, blocks.length);
        return root;
    }

    /**********************
     * Proving Invalidity *
     *********************/

    /**
     * Verify inclusion of the claimed includedStorageSlot & store their results.
     * Note the complexity here is we need to store an empty storage slot as being 32 bytes of zeros
     * to be what the sparse merkle tree expects.
     */
    function verifyAndStoreStorageSlotInclusionProof(
        dt.IncludedStorageSlot memory _includedStorageSlot
    ) private {
        bytes memory accountInfoBytes = getAccountInfoBytes(
            _includedStorageSlot.storageSlot.value
        );
        merkleUtils.verifyAndStore(
            accountInfoBytes,
            uint256(_includedStorageSlot.storageSlot.slotIndex),
            _includedStorageSlot.siblings
        );
    }

    function getStateRootsAndStorageSlots(
        bytes memory _preStateTransition,
        bytes memory _invalidTransition
    ) public returns (bool, bytes32, bytes32, uint256[] memory) {
        bool success;
        bytes memory returnData;
        bytes32 preStateRoot;
        bytes32 postStateRoot;
        uint256[] memory preStateStorageSlots;
        uint256[] memory storageSlots;
        // First decode the prestate root
        (success, returnData) = address(transitionEvaluator).call(
            abi.encodeWithSelector(
                transitionEvaluator
                    .getTransitionStateRootAndAccessList
                    .selector,
                _preStateTransition
            )
        );
        // Emit the output as an event
        emit DecodedTransition(success, returnData);
        // Make sure the call was successful
        require(
            success,
            "If the preStateRoot is invalid, then prove that invalid instead!"
        );
        (preStateRoot, preStateStorageSlots) = abi.decode(
            (returnData),
            (bytes32, uint256[])
        );
        // Now that we have the prestateRoot, let's decode the postState
        (success, returnData) = address(transitionEvaluator).call(
            abi.encodeWithSelector(
                transitionEvaluator
                    .getTransitionStateRootAndAccessList
                    .selector,
                _invalidTransition
            )
        );
        // Emit the output as an event
        emit DecodedTransition(success, returnData);
        // If the call was successful let's decode!
        if (success) {
            (postStateRoot, storageSlots) = abi.decode(
                (returnData),
                (bytes32, uint256[])
            );
        }
        return (success, preStateRoot, postStateRoot, storageSlots);
    }

    /**
     * Checks if a transition is invalid and if it is prunes that block and it's children from the chain.
     */
    function proveTransitionInvalid(
        dt.IncludedTransition memory _preStateIncludedTransition,
        dt.IncludedTransition memory _invalidIncludedTransition,
        dt.IncludedStorageSlot[] memory _transitionStorageSlots
    ) public {
        // For convenience store the transitions
        bytes memory preStateTransition = _preStateIncludedTransition
            .transition;
        bytes memory invalidTransition = _invalidIncludedTransition.transition;

        /********* #1: CHECK_SEQUENTIAL_TRANSITIONS *********/
        // First verify that the transitions are sequential
        verifySequentialTransitions(
            _preStateIncludedTransition,
            _invalidIncludedTransition
        );

        /********* #2: DECODE_TRANSITIONS *********/
        // Decode our transitions and determine which storage slots we'll need in order to validate the transition
        (
            bool success,
            bytes32 preStateRoot,
            bytes32 postStateRoot,
            uint256[] memory storageSlotIndexes
        ) = getStateRootsAndStorageSlots(preStateTransition, invalidTransition);
        // If not success something went wrong with the decoding...
        if (!success) {
            // Prune the block if it has an incorrectly encoded transition!
            pruneBlocksAfter(
                _invalidIncludedTransition.inclusionProof.blockNumber
            );
            return;
        }

        /********* #3: VERIFY_TRANSITION_STORAGE_SLOTS *********/
        // Make sure the storage slots we were given are correct
        for (uint256 i = 0; i < _transitionStorageSlots.length; i++) {
            require(
                _transitionStorageSlots[i].storageSlot.slotIndex ==
                    storageSlotIndexes[i],
                "Supplied storage slot index is incorrect!"
            );
        }

        /********* #4: STORE_STORAGE_INCLUSION_PROOFS *********/
        // Now verify and store the storage inclusion proofs
        merkleUtils.setMerkleRootAndHeight(preStateRoot, STATE_TREE_HEIGHT);
        for (uint256 i = 0; i < _transitionStorageSlots.length; i++) {
            verifyAndStoreStorageSlotInclusionProof(_transitionStorageSlots[i]);
        }

        /********* #5: EVALUATE_TRANSITION *********/
        // Now that we've verified and stored our storage in the state tree, lets apply the transaction
        // To do this first let's pull out the storage slots we care about
        dt.StorageSlot[] memory storageSlots;
        for (uint256 i = 0; i < _transitionStorageSlots.length; i++) {
            storageSlots[i] = _transitionStorageSlots[i].storageSlot;
        }
        bytes memory returnData;
        // Make the external call
        (success, returnData) = address(transitionEvaluator).call(
            abi.encodeWithSelector(
                transitionEvaluator.evaluateTransition.selector,
                invalidTransition,
                storageSlots
            )
        );
        // Check if it was successful. If not, we've got to prune.
        if (!success) {
            pruneBlocksAfter(
                _invalidIncludedTransition.inclusionProof.blockNumber
            );
            return;
        }
        // It was successful so let's decode the outputs to get the new leaf nodes we'll have to insert
        bytes32[] memory outputs = abi.decode((returnData), (bytes32[]));

        /********* #6: UPDATE_STATE_ROOT *********/
        // Now we need to check if the state root is incorrect, to do this we first insert the new leaf values
        for (uint256 i = 0; i < _transitionStorageSlots.length; i++) {
            merkleUtils.updateLeaf(
                outputs[i],
                _transitionStorageSlots[i].storageSlot.slotIndex
            );
        }

        /********* #7: COMPARE_STATE_ROOTS *********/
        // Check if the calculated state root equals what we expect
        if (postStateRoot != merkleUtils.getRoot()) {
            // Prune the block because we found an invalid post state root! Cryptoeconomic validity ftw!
            pruneBlocksAfter(
                _invalidIncludedTransition.inclusionProof.blockNumber
            );
            return;
        }

        // Woah! Looks like there's no fraud!
        revert("No fraud detected!");
    }

    /**
     * Verifies that two transitions were included one after another.
     * This is used to make sure we are comparing the correct
     * prestate & poststate.
     */
    function verifySequentialTransitions(
        dt.IncludedTransition memory _transition0,
        dt.IncludedTransition memory _transition1
    ) public view returns (bool) {
        // Verify inclusion
        require(
            checkTransitionInclusion(_transition0),
            "The first transition must be included!"
        );
        require(
            checkTransitionInclusion(_transition1),
            "The second transition must be included!"
        );

        // Verify that the two transitions are one after another

        // Start by checking if they are in the same block
        if (
            _transition0.inclusionProof.blockNumber ==
            _transition1.inclusionProof.blockNumber
        ) {
            // If the blocknumber is the same, simply check that transition0 preceeds transition1
            require(
                _transition0.inclusionProof.transitionIndex ==
                    _transition1.inclusionProof.transitionIndex - 1,
                "Transitions must be sequential!"
            );
            // Hurray! The transition is valid!
            return true;
        }

        // If not in the same block, we check that:
        // 0) the blocks are one after another
        require(
            _transition0.inclusionProof.blockNumber ==
                _transition1.inclusionProof.blockNumber - 1,
            "Blocks must be one after another or equal."
        );
        // 1) the transitionIndex of transition0 is the last in the block; and
        require(
            _transition0.inclusionProof.transitionIndex ==
                blocks[_transition0.inclusionProof.blockNumber].blockSize - 1,
            "_transition0 must be last in its block."
        );
        // 2) the transitionIndex of transition1 is the first in the block
        require(
            _transition1.inclusionProof.transitionIndex == 0,
            "_transition0 must be first in its block."
        );

        // Success!
        return true;
    }

    /**
     * Check to see if a transition was indeed included.
     */
    function checkTransitionInclusion(
        dt.IncludedTransition memory _includedTransition
    ) public view returns (bool) {
        bytes32 rootHash = blocks[_includedTransition
            .inclusionProof
            .blockNumber]
            .rootHash;
        bool isIncluded = merkleUtils.verify(
            rootHash,
            _includedTransition.transition,
            _includedTransition.inclusionProof.transitionIndex,
            _includedTransition.inclusionProof.siblings
        );
        return isIncluded;
    }

    /**
     * Get the hash of the transition.
     */
    function getTransitionHash(bytes memory _transition)
        public
        pure
        returns (bytes32)
    {
        return keccak256(_transition);
    }

    /**
     * Get the bytes value for this storage.
     */
    function getAccountInfoBytes(dt.AccountInfo memory _accountInfo)
        public
        pure
        returns (bytes memory)
    {
        // If it's an empty storage slot, return 32 bytes of zeros (empty value)
        if (
            _accountInfo.account ==
            0x0000000000000000000000000000000000000000 &&
            _accountInfo.balances.length == 0 &&
            _accountInfo.nonces.length == 0
        ) {
            return abi.encodePacked(uint256(0));
        }
        // Here we don't use `abi.encode([struct])` because it's not clear
        // how to generate that encoding client-side.
        return
            abi.encode(
                _accountInfo.account,
                _accountInfo.balances,
                _accountInfo.nonces
            );
    }

    function deposit(address _token, uint256 _amount) external {
        require(
            IERC20(_token).transferFrom(msg.sender, address(this), _amount),
            "Token transfer failed"
        );
        emit Deposit(msg.sender, _token, _amount);
    }

    function withdraw(dt.IncludedTransition memory _includedTransition) public {
        require(
            checkTransitionInclusion(_includedTransition),
            "Withdraw transition must be included"
        );
        dt.WithdrawTransition memory withdrawTransition = transitionEvaluator
            .decodeWithdrawTransition(_includedTransition.transition);
        address token = tokenRegistry.tokenIndexToTokenAddress(
            withdrawTransition.tokenIndex
        );
        IERC20(token).transfer(msg.sender, withdrawTransition.amount);
    }
}
