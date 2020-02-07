pragma solidity ^0.5.2;
pragma experimental ABIEncoderV2;

import {IERC20} from "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

import {DataTypes as dt} from "./DataTypes.sol";
import {RollupChain} from "./RollupChain.sol";
import {TransitionEvaluator} from "./TransitionEvaluator.sol";
import {TokenRegistry} from "./TokenRegistry.sol";

contract DepositWithdrawManager {
    event Deposit(address accountAddress, address tokenAddress, uint256 amount);

    RollupChain rollupChain;
    TransitionEvaluator transitionEvaluator;
    TokenRegistry tokenRegistry;

    constructor(
        address _rollupChainAddress,
        address _transitionEvaluatorAddress,
        address _tokenRegistryAddress
    ) public {
        rollupChain = RollupChain(_rollupChainAddress);
        transitionEvaluator = TransitionEvaluator(_transitionEvaluatorAddress);
        tokenRegistry = TokenRegistry(_tokenRegistryAddress);
    }

    function deposit(address _tokenAddress, uint256 _amount) external {
        require(
            IERC20(_tokenAddress).transferFrom(
                msg.sender,
                address(this),
                _amount
            ),
            "Token transfer failed"
        );
        emit Deposit(msg.sender, _tokenAddress, _amount);
    }

    function withdraw(dt.IncludedTransition calldata _includedTransition)
        external
    {
        require(
            rollupChain.checkTransitionInclusion(_includedTransition),
            "Withdraw transition must be included"
        );
        dt.WithdrawTransition memory withdrawTransition = transitionEvaluator
            .decodeWithdrawTransition(_includedTransition.transition);
        address tokenAddress = tokenRegistry.tokenIndexToTokenAddress(
            withdrawTransition.tokenIndex
        );
        IERC20(tokenAddress).transfer(msg.sender, withdrawTransition.amount);

    }
}
