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

// RollupTokenRegistryABI is the input ABI used to generate the binding from.
const RollupTokenRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenAddressToTokenIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndexToTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupTokenRegistryBin is the compiled bytecode used for deploying new contracts.
var RollupTokenRegistryBin = "0x6080604052600060035560006100196100bc60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100c4565b600033905090565b610869806100d36000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638f32d59b1161005b5780638f32d59b1461011a5780639bf53f181461013c578063e6250c77146101aa578063f2fde38b146102025761007d565b806309824a8014610082578063715018a6146100c65780638da5cb5b146100d0575b600080fd5b6100c46004803603602081101561009857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610246565b005b6100ce610431565b005b6100d861056a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610122610593565b604051808215151515815260200191505060405180910390f35b6101686004803603602081101561015257600080fd5b81019080803590602001909291905050506105f1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101ec600480360360208110156101c057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610624565b6040518082815260200191505060405180910390f35b6102446004803603602081101561021857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061063c565b005b61024e610593565b6102c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561033c57506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b1561042e57600354600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508060026000600354815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003548173ffffffffffffffffffffffffffffffffffffffff167fd7ca5dc2f8c6bb37c3a4de2a81499b25f8ca8bbb3082010244fe747077d0f6cc60405160405180910390a36001600354016003819055505b50565b610439610593565b6104ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166105d56106c2565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b610644610593565b6106b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6106bf816106ca565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610750576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061080f6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a265627a7a7231582053f6ef03864b4d059096b28949c44ac3484415def1ccab9cb16ad806a92484b764736f6c634300050f0032"

// DeployRollupTokenRegistry deploys a new Ethereum contract, binding an instance of RollupTokenRegistry to it.
func DeployRollupTokenRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTokenRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTokenRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTokenRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTokenRegistry{RollupTokenRegistryCaller: RollupTokenRegistryCaller{contract: contract}, RollupTokenRegistryTransactor: RollupTokenRegistryTransactor{contract: contract}, RollupTokenRegistryFilterer: RollupTokenRegistryFilterer{contract: contract}}, nil
}

// RollupTokenRegistry is an auto generated Go binding around an Ethereum contract.
type RollupTokenRegistry struct {
	RollupTokenRegistryCaller     // Read-only binding to the contract
	RollupTokenRegistryTransactor // Write-only binding to the contract
	RollupTokenRegistryFilterer   // Log filterer for contract events
}

// RollupTokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTokenRegistrySession struct {
	Contract     *RollupTokenRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RollupTokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTokenRegistryCallerSession struct {
	Contract *RollupTokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RollupTokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTokenRegistryTransactorSession struct {
	Contract     *RollupTokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RollupTokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTokenRegistryRaw struct {
	Contract *RollupTokenRegistry // Generic contract binding to access the raw methods on
}

// RollupTokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTokenRegistryCallerRaw struct {
	Contract *RollupTokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTokenRegistryTransactorRaw struct {
	Contract *RollupTokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTokenRegistry creates a new instance of RollupTokenRegistry, bound to a specific deployed contract.
func NewRollupTokenRegistry(address common.Address, backend bind.ContractBackend) (*RollupTokenRegistry, error) {
	contract, err := bindRollupTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTokenRegistry{RollupTokenRegistryCaller: RollupTokenRegistryCaller{contract: contract}, RollupTokenRegistryTransactor: RollupTokenRegistryTransactor{contract: contract}, RollupTokenRegistryFilterer: RollupTokenRegistryFilterer{contract: contract}}, nil
}

// NewRollupTokenRegistryCaller creates a new read-only instance of RollupTokenRegistry, bound to a specific deployed contract.
func NewRollupTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*RollupTokenRegistryCaller, error) {
	contract, err := bindRollupTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTokenRegistryCaller{contract: contract}, nil
}

// NewRollupTokenRegistryTransactor creates a new write-only instance of RollupTokenRegistry, bound to a specific deployed contract.
func NewRollupTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTokenRegistryTransactor, error) {
	contract, err := bindRollupTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTokenRegistryTransactor{contract: contract}, nil
}

// NewRollupTokenRegistryFilterer creates a new log filterer instance of RollupTokenRegistry, bound to a specific deployed contract.
func NewRollupTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTokenRegistryFilterer, error) {
	contract, err := bindRollupTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTokenRegistryFilterer{contract: contract}, nil
}

// bindRollupTokenRegistry binds a generic wrapper to an already deployed contract.
func bindRollupTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTokenRegistry *RollupTokenRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTokenRegistry.Contract.RollupTokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTokenRegistry *RollupTokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.RollupTokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTokenRegistry *RollupTokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.RollupTokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTokenRegistry *RollupTokenRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTokenRegistry *RollupTokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTokenRegistry *RollupTokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RollupTokenRegistry *RollupTokenRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RollupTokenRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RollupTokenRegistry *RollupTokenRegistrySession) IsOwner() (bool, error) {
	return _RollupTokenRegistry.Contract.IsOwner(&_RollupTokenRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RollupTokenRegistry *RollupTokenRegistryCallerSession) IsOwner() (bool, error) {
	return _RollupTokenRegistry.Contract.IsOwner(&_RollupTokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RollupTokenRegistry *RollupTokenRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RollupTokenRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RollupTokenRegistry *RollupTokenRegistrySession) Owner() (common.Address, error) {
	return _RollupTokenRegistry.Contract.Owner(&_RollupTokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RollupTokenRegistry *RollupTokenRegistryCallerSession) Owner() (common.Address, error) {
	return _RollupTokenRegistry.Contract.Owner(&_RollupTokenRegistry.CallOpts)
}

// TokenAddressToTokenIndex is a free data retrieval call binding the contract method 0xe6250c77.
//
// Solidity: function tokenAddressToTokenIndex(address ) constant returns(uint256)
func (_RollupTokenRegistry *RollupTokenRegistryCaller) TokenAddressToTokenIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RollupTokenRegistry.contract.Call(opts, out, "tokenAddressToTokenIndex", arg0)
	return *ret0, err
}

// TokenAddressToTokenIndex is a free data retrieval call binding the contract method 0xe6250c77.
//
// Solidity: function tokenAddressToTokenIndex(address ) constant returns(uint256)
func (_RollupTokenRegistry *RollupTokenRegistrySession) TokenAddressToTokenIndex(arg0 common.Address) (*big.Int, error) {
	return _RollupTokenRegistry.Contract.TokenAddressToTokenIndex(&_RollupTokenRegistry.CallOpts, arg0)
}

// TokenAddressToTokenIndex is a free data retrieval call binding the contract method 0xe6250c77.
//
// Solidity: function tokenAddressToTokenIndex(address ) constant returns(uint256)
func (_RollupTokenRegistry *RollupTokenRegistryCallerSession) TokenAddressToTokenIndex(arg0 common.Address) (*big.Int, error) {
	return _RollupTokenRegistry.Contract.TokenAddressToTokenIndex(&_RollupTokenRegistry.CallOpts, arg0)
}

// TokenIndexToTokenAddress is a free data retrieval call binding the contract method 0x9bf53f18.
//
// Solidity: function tokenIndexToTokenAddress(uint256 ) constant returns(address)
func (_RollupTokenRegistry *RollupTokenRegistryCaller) TokenIndexToTokenAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RollupTokenRegistry.contract.Call(opts, out, "tokenIndexToTokenAddress", arg0)
	return *ret0, err
}

// TokenIndexToTokenAddress is a free data retrieval call binding the contract method 0x9bf53f18.
//
// Solidity: function tokenIndexToTokenAddress(uint256 ) constant returns(address)
func (_RollupTokenRegistry *RollupTokenRegistrySession) TokenIndexToTokenAddress(arg0 *big.Int) (common.Address, error) {
	return _RollupTokenRegistry.Contract.TokenIndexToTokenAddress(&_RollupTokenRegistry.CallOpts, arg0)
}

// TokenIndexToTokenAddress is a free data retrieval call binding the contract method 0x9bf53f18.
//
// Solidity: function tokenIndexToTokenAddress(uint256 ) constant returns(address)
func (_RollupTokenRegistry *RollupTokenRegistryCallerSession) TokenIndexToTokenAddress(arg0 *big.Int) (common.Address, error) {
	return _RollupTokenRegistry.Contract.TokenIndexToTokenAddress(&_RollupTokenRegistry.CallOpts, arg0)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _tokenAddress) returns()
func (_RollupTokenRegistry *RollupTokenRegistryTransactor) RegisterToken(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _RollupTokenRegistry.contract.Transact(opts, "registerToken", _tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _tokenAddress) returns()
func (_RollupTokenRegistry *RollupTokenRegistrySession) RegisterToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.RegisterToken(&_RollupTokenRegistry.TransactOpts, _tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _tokenAddress) returns()
func (_RollupTokenRegistry *RollupTokenRegistryTransactorSession) RegisterToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.RegisterToken(&_RollupTokenRegistry.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupTokenRegistry *RollupTokenRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTokenRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupTokenRegistry *RollupTokenRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.RenounceOwnership(&_RollupTokenRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupTokenRegistry *RollupTokenRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.RenounceOwnership(&_RollupTokenRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupTokenRegistry *RollupTokenRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupTokenRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupTokenRegistry *RollupTokenRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.TransferOwnership(&_RollupTokenRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupTokenRegistry *RollupTokenRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupTokenRegistry.Contract.TransferOwnership(&_RollupTokenRegistry.TransactOpts, newOwner)
}

// RollupTokenRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RollupTokenRegistry contract.
type RollupTokenRegistryOwnershipTransferredIterator struct {
	Event *RollupTokenRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupTokenRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupTokenRegistryOwnershipTransferred)
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
		it.Event = new(RollupTokenRegistryOwnershipTransferred)
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
func (it *RollupTokenRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupTokenRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupTokenRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the RollupTokenRegistry contract.
type RollupTokenRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupTokenRegistry *RollupTokenRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupTokenRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupTokenRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupTokenRegistryOwnershipTransferredIterator{contract: _RollupTokenRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupTokenRegistry *RollupTokenRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupTokenRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupTokenRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupTokenRegistryOwnershipTransferred)
				if err := _RollupTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RollupTokenRegistry *RollupTokenRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RollupTokenRegistryOwnershipTransferred, error) {
	event := new(RollupTokenRegistryOwnershipTransferred)
	if err := _RollupTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupTokenRegistryTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the RollupTokenRegistry contract.
type RollupTokenRegistryTokenRegisteredIterator struct {
	Event *RollupTokenRegistryTokenRegistered // Event containing the contract specifics and raw log

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
func (it *RollupTokenRegistryTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupTokenRegistryTokenRegistered)
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
		it.Event = new(RollupTokenRegistryTokenRegistered)
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
func (it *RollupTokenRegistryTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupTokenRegistryTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupTokenRegistryTokenRegistered represents a TokenRegistered event raised by the RollupTokenRegistry contract.
type RollupTokenRegistryTokenRegistered struct {
	TokenAddress common.Address
	TokenIndex   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xd7ca5dc2f8c6bb37c3a4de2a81499b25f8ca8bbb3082010244fe747077d0f6cc.
//
// Solidity: event TokenRegistered(address indexed tokenAddress, uint256 indexed tokenIndex)
func (_RollupTokenRegistry *RollupTokenRegistryFilterer) FilterTokenRegistered(opts *bind.FilterOpts, tokenAddress []common.Address, tokenIndex []*big.Int) (*RollupTokenRegistryTokenRegisteredIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var tokenIndexRule []interface{}
	for _, tokenIndexItem := range tokenIndex {
		tokenIndexRule = append(tokenIndexRule, tokenIndexItem)
	}

	logs, sub, err := _RollupTokenRegistry.contract.FilterLogs(opts, "TokenRegistered", tokenAddressRule, tokenIndexRule)
	if err != nil {
		return nil, err
	}
	return &RollupTokenRegistryTokenRegisteredIterator{contract: _RollupTokenRegistry.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xd7ca5dc2f8c6bb37c3a4de2a81499b25f8ca8bbb3082010244fe747077d0f6cc.
//
// Solidity: event TokenRegistered(address indexed tokenAddress, uint256 indexed tokenIndex)
func (_RollupTokenRegistry *RollupTokenRegistryFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *RollupTokenRegistryTokenRegistered, tokenAddress []common.Address, tokenIndex []*big.Int) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var tokenIndexRule []interface{}
	for _, tokenIndexItem := range tokenIndex {
		tokenIndexRule = append(tokenIndexRule, tokenIndexItem)
	}

	logs, sub, err := _RollupTokenRegistry.contract.WatchLogs(opts, "TokenRegistered", tokenAddressRule, tokenIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupTokenRegistryTokenRegistered)
				if err := _RollupTokenRegistry.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0xd7ca5dc2f8c6bb37c3a4de2a81499b25f8ca8bbb3082010244fe747077d0f6cc.
//
// Solidity: event TokenRegistered(address indexed tokenAddress, uint256 indexed tokenIndex)
func (_RollupTokenRegistry *RollupTokenRegistryFilterer) ParseTokenRegistered(log types.Log) (*RollupTokenRegistryTokenRegistered, error) {
	event := new(RollupTokenRegistryTokenRegistered)
	if err := _RollupTokenRegistry.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
