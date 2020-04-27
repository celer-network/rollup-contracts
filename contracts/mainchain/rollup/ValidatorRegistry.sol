pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import "openzeppelin-solidity/contracts/access/Ownable.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import {RollupChain} from "./RollupChain.sol";


contract ValidatorRegistry is Ownable {
    address[] public validators;
    address public currentCommitter;
    uint256 private currentCommitterIndex;
    RollupChain rollupChain;

    modifier onlyRollupChain() {
        require(
            msg.sender == address(rollupChain),
            "Only RollupChain may perform action"
        );
        _;
    }

    constructor(address[] memory _validators) public {
        validators = _validators;
    }

    function setRollupChainAddress(address _rollupChainAddress)
        external
        onlyOwner
    {
        rollupChain = RollupChain(_rollupChainAddress);
    }

    function setValidators(address[] calldata _validators) external onlyOwner {
        validators = _validators;
        currentCommitterIndex = 0;
    }

    function checkSignatures(
        uint256 _blockNumber,
        bytes[] calldata _transitions,
        bytes[] calldata _signatures
    ) external view onlyRollupChain returns (bool) {
        uint256 numValidators = validators.length;
        uint256 numSignatures;
        for (uint256 i = 0; i < numValidators; i++) {
            bytes32 blockHash = keccak256(
                abi.encode(_blockNumber, _transitions)
            );
            bytes32 prefixedHash = ECDSA.toEthSignedMessageHash(blockHash);
            require(
                ECDSA.recover(prefixedHash, _signatures[i]) == validators[i],
                "Signature is invalid!"
            );
            numSignatures++;
        }

        // Require signatures from all the validators if less than 4, or 2/3 of
        // the validators if at least 4.
        bool hasEnoughSignatures = numValidators < 4
            ? numSignatures == numValidators
            : numSignatures * 3 > numValidators * 2;
        require(hasEnoughSignatures, "Not enough signatures");
        return true;
    }

    function pickNextCommitter() external onlyRollupChain {
        currentCommitterIndex = (currentCommitterIndex + 1) % validators.length;
        currentCommitter = validators[currentCommitterIndex];
        rollupChain.setCommitterAddress(currentCommitter);
    }
}
