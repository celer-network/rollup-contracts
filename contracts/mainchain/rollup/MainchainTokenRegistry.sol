pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";


contract MainchainTokenRegistry is Ownable {
    event TokenRegisteredOnMainchain(address tokenAddress);

    mapping(address => uint256) public tokenAddressToTokenIndex;
    mapping(uint256 => address) public tokenIndexToTokenAddress;
    uint256 numTokens = 0;

    function registerToken(address _tokenAddress) external onlyOwner {
        // Register token with an index if it isn't already
        if (
            _tokenAddress != address(0) &&
            tokenAddressToTokenIndex[_tokenAddress] == 0
        ) {
            tokenAddressToTokenIndex[_tokenAddress] = numTokens;
            tokenIndexToTokenAddress[numTokens] = _tokenAddress;
            numTokens = numTokens + 1;
            emit TokenRegisteredOnMainchain(_tokenAddress);
        }
    }
}
