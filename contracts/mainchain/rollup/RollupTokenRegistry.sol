pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";


contract RollupTokenRegistry is Ownable {
    event TokenRegistered(
        address indexed tokenAddress,
        uint256 indexed tokenIndex
    );

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
            emit TokenRegistered(_tokenAddress, numTokens);
            numTokens = numTokens + 1;
        }
    }
}
