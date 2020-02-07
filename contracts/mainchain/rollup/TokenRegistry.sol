pragma solidity ^0.5.2;

contract TokenRegistry {
    mapping(address => uint256) public tokenAddressToTokenIndex;
    mapping(uint256 => address) public tokenIndexToTokenAddress;
    uint256 numTokens = 0;

    function registerToken(address _tokenAddress) external {
        // Register token with an index if it isn't already
        if (
            _tokenAddress != address(0) &&
            tokenAddressToTokenIndex[_tokenAddress] == 0
        ) {
            tokenAddressToTokenIndex[_tokenAddress] = numTokens;
            tokenIndexToTokenAddress[numTokens] = _tokenAddress;
            numTokens = numTokens + 1;
        }
    }
}
