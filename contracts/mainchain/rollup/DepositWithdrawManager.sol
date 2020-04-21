pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import {IERC20} from "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

import {DataTypes as dt} from "./DataTypes.sol";
import {RollupChain} from "./RollupChain.sol";
import {TransitionEvaluator} from "./TransitionEvaluator.sol";
import {RollupTokenRegistry} from "./RollupTokenRegistry.sol";


contract DepositWithdrawManager {
    event Deposit(address account, address token, uint256 amount);

    RollupChain rollupChain;
    TransitionEvaluator transitionEvaluator;
    RollupTokenRegistry tokenRegistry;

    constructor(
        address _rollupChainAddress,
        address _transitionEvaluatorAddress,
        address _tokenRegistryAddress
    ) public {
        rollupChain = RollupChain(_rollupChainAddress);
        transitionEvaluator = TransitionEvaluator(_transitionEvaluatorAddress);
        tokenRegistry = RollupTokenRegistry(_tokenRegistryAddress);
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
            rollupChain.checkTransitionInclusion(_includedTransition),
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
