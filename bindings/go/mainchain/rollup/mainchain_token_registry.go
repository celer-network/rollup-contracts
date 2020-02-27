// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollup

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MainchainTokenRegistryABI is the input ABI used to generate the binding from.
const MainchainTokenRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRegisteredOnMainchain\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenAddressToTokenIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndexToTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MainchainTokenRegistryBin is the compiled bytecode used for deploying new contracts.
var MainchainTokenRegistryBin = "0x6080604052600060035560006100196100bc60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100c4565b600033905090565b610886806100d36000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638f32d59b1161005b5780638f32d59b1461011a5780639bf53f181461013c578063e6250c77146101aa578063f2fde38b146102025761007d565b806309824a8014610082578063715018a6146100c65780638da5cb5b146100d0575b600080fd5b6100c46004803603602081101561009857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610246565b005b6100ce61044e565b005b6100d8610587565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101226105b0565b604051808215151515815260200191505060405180910390f35b6101686004803603602081101561015257600080fd5b810190808035906020019092919050505061060e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101ec600480360360208110156101c057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610641565b6040518082815260200191505060405180910390f35b6102446004803603602081101561021857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610659565b005b61024e6105b0565b6102c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561033c57506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b1561044b57600354600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508060026000600354815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600354016003819055507ff575fd7740ca9354404444fc33465bfc6da92b01bdfe4dd67693a49b6a539eb981604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b50565b6104566105b0565b6104c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166105f26106df565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b6106616105b0565b6106d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6106dc816106e7565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561076d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061082c6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a265627a7a72315820d1fb745e9a82cee238fb8001751fa5beed897cacace61ab7f50d0ddea76075b364736f6c634300050f0032"

// DeployMainchainTokenRegistry deploys a new Ethereum contract, binding an instance of MainchainTokenRegistry to it.
func DeployMainchainTokenRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MainchainTokenRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(MainchainTokenRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainchainTokenRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainchainTokenRegistry{MainchainTokenRegistryCaller: MainchainTokenRegistryCaller{contract: contract}, MainchainTokenRegistryTransactor: MainchainTokenRegistryTransactor{contract: contract}, MainchainTokenRegistryFilterer: MainchainTokenRegistryFilterer{contract: contract}}, nil
}

// MainchainTokenRegistry is an auto generated Go binding around an Ethereum contract.
type MainchainTokenRegistry struct {
	MainchainTokenRegistryCaller     // Read-only binding to the contract
	MainchainTokenRegistryTransactor // Write-only binding to the contract
	MainchainTokenRegistryFilterer   // Log filterer for contract events
}

// MainchainTokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainchainTokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainchainTokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainchainTokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainchainTokenRegistrySession struct {
	Contract     *MainchainTokenRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MainchainTokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainchainTokenRegistryCallerSession struct {
	Contract *MainchainTokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MainchainTokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainchainTokenRegistryTransactorSession struct {
	Contract     *MainchainTokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MainchainTokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainchainTokenRegistryRaw struct {
	Contract *MainchainTokenRegistry // Generic contract binding to access the raw methods on
}

// MainchainTokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainchainTokenRegistryCallerRaw struct {
	Contract *MainchainTokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MainchainTokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainchainTokenRegistryTransactorRaw struct {
	Contract *MainchainTokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainchainTokenRegistry creates a new instance of MainchainTokenRegistry, bound to a specific deployed contract.
func NewMainchainTokenRegistry(address common.Address, backend bind.ContractBackend) (*MainchainTokenRegistry, error) {
	contract, err := bindMainchainTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainchainTokenRegistry{MainchainTokenRegistryCaller: MainchainTokenRegistryCaller{contract: contract}, MainchainTokenRegistryTransactor: MainchainTokenRegistryTransactor{contract: contract}, MainchainTokenRegistryFilterer: MainchainTokenRegistryFilterer{contract: contract}}, nil
}

// NewMainchainTokenRegistryCaller creates a new read-only instance of MainchainTokenRegistry, bound to a specific deployed contract.
func NewMainchainTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*MainchainTokenRegistryCaller, error) {
	contract, err := bindMainchainTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTokenRegistryCaller{contract: contract}, nil
}

// NewMainchainTokenRegistryTransactor creates a new write-only instance of MainchainTokenRegistry, bound to a specific deployed contract.
func NewMainchainTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MainchainTokenRegistryTransactor, error) {
	contract, err := bindMainchainTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTokenRegistryTransactor{contract: contract}, nil
}

// NewMainchainTokenRegistryFilterer creates a new log filterer instance of MainchainTokenRegistry, bound to a specific deployed contract.
func NewMainchainTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MainchainTokenRegistryFilterer, error) {
	contract, err := bindMainchainTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainchainTokenRegistryFilterer{contract: contract}, nil
}

// bindMainchainTokenRegistry binds a generic wrapper to an already deployed contract.
func bindMainchainTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainchainTokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTokenRegistry *MainchainTokenRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainchainTokenRegistry.Contract.MainchainTokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTokenRegistry *MainchainTokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.MainchainTokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTokenRegistry *MainchainTokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.MainchainTokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTokenRegistry *MainchainTokenRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainchainTokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_MainchainTokenRegistry *MainchainTokenRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainchainTokenRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_MainchainTokenRegistry *MainchainTokenRegistrySession) IsOwner() (bool, error) {
	return _MainchainTokenRegistry.Contract.IsOwner(&_MainchainTokenRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_MainchainTokenRegistry *MainchainTokenRegistryCallerSession) IsOwner() (bool, error) {
	return _MainchainTokenRegistry.Contract.IsOwner(&_MainchainTokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MainchainTokenRegistry *MainchainTokenRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainchainTokenRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MainchainTokenRegistry *MainchainTokenRegistrySession) Owner() (common.Address, error) {
	return _MainchainTokenRegistry.Contract.Owner(&_MainchainTokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MainchainTokenRegistry *MainchainTokenRegistryCallerSession) Owner() (common.Address, error) {
	return _MainchainTokenRegistry.Contract.Owner(&_MainchainTokenRegistry.CallOpts)
}

// TokenAddressToTokenIndex is a free data retrieval call binding the contract method 0xe6250c77.
//
// Solidity: function tokenAddressToTokenIndex(address ) constant returns(uint256)
func (_MainchainTokenRegistry *MainchainTokenRegistryCaller) TokenAddressToTokenIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainchainTokenRegistry.contract.Call(opts, out, "tokenAddressToTokenIndex", arg0)
	return *ret0, err
}

// TokenAddressToTokenIndex is a free data retrieval call binding the contract method 0xe6250c77.
//
// Solidity: function tokenAddressToTokenIndex(address ) constant returns(uint256)
func (_MainchainTokenRegistry *MainchainTokenRegistrySession) TokenAddressToTokenIndex(arg0 common.Address) (*big.Int, error) {
	return _MainchainTokenRegistry.Contract.TokenAddressToTokenIndex(&_MainchainTokenRegistry.CallOpts, arg0)
}

// TokenAddressToTokenIndex is a free data retrieval call binding the contract method 0xe6250c77.
//
// Solidity: function tokenAddressToTokenIndex(address ) constant returns(uint256)
func (_MainchainTokenRegistry *MainchainTokenRegistryCallerSession) TokenAddressToTokenIndex(arg0 common.Address) (*big.Int, error) {
	return _MainchainTokenRegistry.Contract.TokenAddressToTokenIndex(&_MainchainTokenRegistry.CallOpts, arg0)
}

// TokenIndexToTokenAddress is a free data retrieval call binding the contract method 0x9bf53f18.
//
// Solidity: function tokenIndexToTokenAddress(uint256 ) constant returns(address)
func (_MainchainTokenRegistry *MainchainTokenRegistryCaller) TokenIndexToTokenAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainchainTokenRegistry.contract.Call(opts, out, "tokenIndexToTokenAddress", arg0)
	return *ret0, err
}

// TokenIndexToTokenAddress is a free data retrieval call binding the contract method 0x9bf53f18.
//
// Solidity: function tokenIndexToTokenAddress(uint256 ) constant returns(address)
func (_MainchainTokenRegistry *MainchainTokenRegistrySession) TokenIndexToTokenAddress(arg0 *big.Int) (common.Address, error) {
	return _MainchainTokenRegistry.Contract.TokenIndexToTokenAddress(&_MainchainTokenRegistry.CallOpts, arg0)
}

// TokenIndexToTokenAddress is a free data retrieval call binding the contract method 0x9bf53f18.
//
// Solidity: function tokenIndexToTokenAddress(uint256 ) constant returns(address)
func (_MainchainTokenRegistry *MainchainTokenRegistryCallerSession) TokenIndexToTokenAddress(arg0 *big.Int) (common.Address, error) {
	return _MainchainTokenRegistry.Contract.TokenIndexToTokenAddress(&_MainchainTokenRegistry.CallOpts, arg0)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _tokenAddress) returns()
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactor) RegisterToken(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _MainchainTokenRegistry.contract.Transact(opts, "registerToken", _tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _tokenAddress) returns()
func (_MainchainTokenRegistry *MainchainTokenRegistrySession) RegisterToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.RegisterToken(&_MainchainTokenRegistry.TransactOpts, _tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _tokenAddress) returns()
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactorSession) RegisterToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.RegisterToken(&_MainchainTokenRegistry.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTokenRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainchainTokenRegistry *MainchainTokenRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.RenounceOwnership(&_MainchainTokenRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.RenounceOwnership(&_MainchainTokenRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MainchainTokenRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MainchainTokenRegistry *MainchainTokenRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.TransferOwnership(&_MainchainTokenRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MainchainTokenRegistry *MainchainTokenRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MainchainTokenRegistry.Contract.TransferOwnership(&_MainchainTokenRegistry.TransactOpts, newOwner)
}

// MainchainTokenRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MainchainTokenRegistry contract.
type MainchainTokenRegistryOwnershipTransferredIterator struct {
	Event *MainchainTokenRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainchainTokenRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainTokenRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainchainTokenRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainchainTokenRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainTokenRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainTokenRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the MainchainTokenRegistry contract.
type MainchainTokenRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MainchainTokenRegistry *MainchainTokenRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainchainTokenRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MainchainTokenRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainchainTokenRegistryOwnershipTransferredIterator{contract: _MainchainTokenRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MainchainTokenRegistry *MainchainTokenRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainchainTokenRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MainchainTokenRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainTokenRegistryOwnershipTransferred)
				if err := _MainchainTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MainchainTokenRegistry *MainchainTokenRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*MainchainTokenRegistryOwnershipTransferred, error) {
	event := new(MainchainTokenRegistryOwnershipTransferred)
	if err := _MainchainTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MainchainTokenRegistryTokenRegisteredOnMainchainIterator is returned from FilterTokenRegisteredOnMainchain and is used to iterate over the raw logs and unpacked data for TokenRegisteredOnMainchain events raised by the MainchainTokenRegistry contract.
type MainchainTokenRegistryTokenRegisteredOnMainchainIterator struct {
	Event *MainchainTokenRegistryTokenRegisteredOnMainchain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainchainTokenRegistryTokenRegisteredOnMainchainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainTokenRegistryTokenRegisteredOnMainchain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainchainTokenRegistryTokenRegisteredOnMainchain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainchainTokenRegistryTokenRegisteredOnMainchainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainTokenRegistryTokenRegisteredOnMainchainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainTokenRegistryTokenRegisteredOnMainchain represents a TokenRegisteredOnMainchain event raised by the MainchainTokenRegistry contract.
type MainchainTokenRegistryTokenRegisteredOnMainchain struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRegisteredOnMainchain is a free log retrieval operation binding the contract event 0xf575fd7740ca9354404444fc33465bfc6da92b01bdfe4dd67693a49b6a539eb9.
//
// Solidity: event TokenRegisteredOnMainchain(address tokenAddress)
func (_MainchainTokenRegistry *MainchainTokenRegistryFilterer) FilterTokenRegisteredOnMainchain(opts *bind.FilterOpts) (*MainchainTokenRegistryTokenRegisteredOnMainchainIterator, error) {

	logs, sub, err := _MainchainTokenRegistry.contract.FilterLogs(opts, "TokenRegisteredOnMainchain")
	if err != nil {
		return nil, err
	}
	return &MainchainTokenRegistryTokenRegisteredOnMainchainIterator{contract: _MainchainTokenRegistry.contract, event: "TokenRegisteredOnMainchain", logs: logs, sub: sub}, nil
}

// WatchTokenRegisteredOnMainchain is a free log subscription operation binding the contract event 0xf575fd7740ca9354404444fc33465bfc6da92b01bdfe4dd67693a49b6a539eb9.
//
// Solidity: event TokenRegisteredOnMainchain(address tokenAddress)
func (_MainchainTokenRegistry *MainchainTokenRegistryFilterer) WatchTokenRegisteredOnMainchain(opts *bind.WatchOpts, sink chan<- *MainchainTokenRegistryTokenRegisteredOnMainchain) (event.Subscription, error) {

	logs, sub, err := _MainchainTokenRegistry.contract.WatchLogs(opts, "TokenRegisteredOnMainchain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainTokenRegistryTokenRegisteredOnMainchain)
				if err := _MainchainTokenRegistry.contract.UnpackLog(event, "TokenRegisteredOnMainchain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenRegisteredOnMainchain is a log parse operation binding the contract event 0xf575fd7740ca9354404444fc33465bfc6da92b01bdfe4dd67693a49b6a539eb9.
//
// Solidity: event TokenRegisteredOnMainchain(address tokenAddress)
func (_MainchainTokenRegistry *MainchainTokenRegistryFilterer) ParseTokenRegisteredOnMainchain(log types.Log) (*MainchainTokenRegistryTokenRegisteredOnMainchain, error) {
	event := new(MainchainTokenRegistryTokenRegisteredOnMainchain)
	if err := _MainchainTokenRegistry.contract.UnpackLog(event, "TokenRegisteredOnMainchain", log); err != nil {
		return nil, err
	}
	return event, nil
}
