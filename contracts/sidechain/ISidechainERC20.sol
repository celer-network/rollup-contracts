pragma solidity ^0.5.2;


interface ISidechainERC20 {
    function transfer(
        address sender,
        address recipient,
        uint256 amount,
        bytes calldata signature
    ) external returns (bool);
}
