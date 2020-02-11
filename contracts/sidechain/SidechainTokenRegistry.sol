pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./SidechainERC20.sol";

contract SidechainTokenRegistry is Ownable {
    event TokenRegisteredOnSidechain(
        address mainchainToken,
        address sidechainToken
    );

    mapping(address => address) public mainchainTokenToSidechainToken;
    mapping(address => address) public sidechainTokenToMainchainToken;

    function registerToken(
        address _mainchainToken,
        string memory _name,
        string memory _symbol,
        uint8 _decimals
    ) public onlyOwner returns (address token) {
        require(
            mainchainTokenToSidechainToken[_mainchainToken] == address(0x0),
            "Token already mapped"
        );

        address sidechainToken = address(
            new SidechainERC20(_mainchainToken, _name, _symbol, _decimals)
        );

        mainchainTokenToSidechainToken[_mainchainToken] = sidechainToken;
        sidechainTokenToMainchainToken[sidechainToken] = _mainchainToken;

        emit TokenRegisteredOnSidechain(_mainchainToken, sidechainToken);

        return sidechainToken;
    }
}
