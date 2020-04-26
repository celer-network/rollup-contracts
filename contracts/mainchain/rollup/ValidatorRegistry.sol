pragma solidity ^0.6.6;

import "openzeppelin-solidity/contracts/access/Ownable.sol";
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

    function pickNextCommitter() external onlyRollupChain {
        currentCommitterIndex = (currentCommitterIndex + 1) % validators.length;
        currentCommitter = validators[currentCommitterIndex];
        rollupChain.setCommitterAddress(currentCommitter);
    }
}
