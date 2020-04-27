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

// AccountRegistryABI is the input ABI used to generate the binding from.
const AccountRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccountRegistryBin is the compiled bytecode used for deploying new contracts.
var AccountRegistryBin = "0x608060405260006100146100b760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100bf565b600033905090565b610c60806100ce6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063715018a61461005c5780638da5cb5b14610066578063d8fa638d146100b0578063f2cd61c91461010c578063f2fde38b146101a5575b600080fd5b6100646101e9565b005b61006e610371565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100f2600480360360208110156100c657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061039a565b604051808215151515815260200191505060405180910390f35b6101a36004803603604081101561012257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561015f57600080fd5b82018360208201111561017157600080fd5b8035906020019184600183028401116401000000008311171561019357600080fd5b90919293919293905050506103ba565b005b6101e7600480360360208110156101bb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106b8565b005b6101f16108c5565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915054906101000a900460ff1681565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561047a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4163636f756e7420616c7265616479207265676973746572656400000000000081525060200191505060405180910390fd5b600030604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401807f72656769737465724163636f756e740000000000000000000000000000000000815250600f019150506040516020818303038152906040528051906020012090506000610507826108cd565b90508473ffffffffffffffffffffffffffffffffffffffff1661056e8286868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610925565b73ffffffffffffffffffffffffffffffffffffffff16146105f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5265676973746572207369676e617475726520697320696e76616c696421000081525060200191505060405180910390fd5b60018060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fcd822dc9688e20acea68724a2fbcfe4f3e526d20ecaa37b18fe3047ab377d6a585604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15050505050565b6106c06108c5565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610781576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610807576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610bc16026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b60008160405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c01828152602001915050604051602081830303815290604052805190602001209050919050565b6000604182511461099e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45434453413a20696e76616c6964207369676e6174757265206c656e6774680081525060200191505060405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610a37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610be76022913960400191505060405180910390fd5b601b8160ff1614158015610a4f5750601c8160ff1614155b15610aa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610c096022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610b04573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610bb3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f45434453413a20696e76616c6964207369676e6174757265000000000000000081525060200191505060405180910390fd5b809450505050509291505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c7565a26469706673582212209d42860dbe5f9c252c627d4ec2b8bf1a932df05eccb01412ce6081d4b02358e464736f6c63430006060033"

// DeployAccountRegistry deploys a new Ethereum contract, binding an instance of AccountRegistry to it.
func DeployAccountRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountRegistry{AccountRegistryCaller: AccountRegistryCaller{contract: contract}, AccountRegistryTransactor: AccountRegistryTransactor{contract: contract}, AccountRegistryFilterer: AccountRegistryFilterer{contract: contract}}, nil
}

// AccountRegistry is an auto generated Go binding around an Ethereum contract.
type AccountRegistry struct {
	AccountRegistryCaller     // Read-only binding to the contract
	AccountRegistryTransactor // Write-only binding to the contract
	AccountRegistryFilterer   // Log filterer for contract events
}

// AccountRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountRegistrySession struct {
	Contract     *AccountRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountRegistryCallerSession struct {
	Contract *AccountRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AccountRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountRegistryTransactorSession struct {
	Contract     *AccountRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AccountRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRegistryRaw struct {
	Contract *AccountRegistry // Generic contract binding to access the raw methods on
}

// AccountRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountRegistryCallerRaw struct {
	Contract *AccountRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountRegistryTransactorRaw struct {
	Contract *AccountRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountRegistry creates a new instance of AccountRegistry, bound to a specific deployed contract.
func NewAccountRegistry(address common.Address, backend bind.ContractBackend) (*AccountRegistry, error) {
	contract, err := bindAccountRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountRegistry{AccountRegistryCaller: AccountRegistryCaller{contract: contract}, AccountRegistryTransactor: AccountRegistryTransactor{contract: contract}, AccountRegistryFilterer: AccountRegistryFilterer{contract: contract}}, nil
}

// NewAccountRegistryCaller creates a new read-only instance of AccountRegistry, bound to a specific deployed contract.
func NewAccountRegistryCaller(address common.Address, caller bind.ContractCaller) (*AccountRegistryCaller, error) {
	contract, err := bindAccountRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountRegistryCaller{contract: contract}, nil
}

// NewAccountRegistryTransactor creates a new write-only instance of AccountRegistry, bound to a specific deployed contract.
func NewAccountRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountRegistryTransactor, error) {
	contract, err := bindAccountRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountRegistryTransactor{contract: contract}, nil
}

// NewAccountRegistryFilterer creates a new log filterer instance of AccountRegistry, bound to a specific deployed contract.
func NewAccountRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountRegistryFilterer, error) {
	contract, err := bindAccountRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountRegistryFilterer{contract: contract}, nil
}

// bindAccountRegistry binds a generic wrapper to an already deployed contract.
func bindAccountRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountRegistry *AccountRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountRegistry.Contract.AccountRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountRegistry *AccountRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountRegistry.Contract.AccountRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountRegistry *AccountRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountRegistry.Contract.AccountRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountRegistry *AccountRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountRegistry *AccountRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountRegistry *AccountRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountRegistry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountRegistry *AccountRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountRegistry *AccountRegistrySession) Owner() (common.Address, error) {
	return _AccountRegistry.Contract.Owner(&_AccountRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountRegistry *AccountRegistryCallerSession) Owner() (common.Address, error) {
	return _AccountRegistry.Contract.Owner(&_AccountRegistry.CallOpts)
}

// RegisteredAccounts is a free data retrieval call binding the contract method 0xd8fa638d.
//
// Solidity: function registeredAccounts(address ) view returns(bool)
func (_AccountRegistry *AccountRegistryCaller) RegisteredAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountRegistry.contract.Call(opts, out, "registeredAccounts", arg0)
	return *ret0, err
}

// RegisteredAccounts is a free data retrieval call binding the contract method 0xd8fa638d.
//
// Solidity: function registeredAccounts(address ) view returns(bool)
func (_AccountRegistry *AccountRegistrySession) RegisteredAccounts(arg0 common.Address) (bool, error) {
	return _AccountRegistry.Contract.RegisteredAccounts(&_AccountRegistry.CallOpts, arg0)
}

// RegisteredAccounts is a free data retrieval call binding the contract method 0xd8fa638d.
//
// Solidity: function registeredAccounts(address ) view returns(bool)
func (_AccountRegistry *AccountRegistryCallerSession) RegisteredAccounts(arg0 common.Address) (bool, error) {
	return _AccountRegistry.Contract.RegisteredAccounts(&_AccountRegistry.CallOpts, arg0)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0xf2cd61c9.
//
// Solidity: function registerAccount(address _account, bytes _signature) returns()
func (_AccountRegistry *AccountRegistryTransactor) RegisterAccount(opts *bind.TransactOpts, _account common.Address, _signature []byte) (*types.Transaction, error) {
	return _AccountRegistry.contract.Transact(opts, "registerAccount", _account, _signature)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0xf2cd61c9.
//
// Solidity: function registerAccount(address _account, bytes _signature) returns()
func (_AccountRegistry *AccountRegistrySession) RegisterAccount(_account common.Address, _signature []byte) (*types.Transaction, error) {
	return _AccountRegistry.Contract.RegisterAccount(&_AccountRegistry.TransactOpts, _account, _signature)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0xf2cd61c9.
//
// Solidity: function registerAccount(address _account, bytes _signature) returns()
func (_AccountRegistry *AccountRegistryTransactorSession) RegisterAccount(_account common.Address, _signature []byte) (*types.Transaction, error) {
	return _AccountRegistry.Contract.RegisterAccount(&_AccountRegistry.TransactOpts, _account, _signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountRegistry *AccountRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountRegistry *AccountRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountRegistry.Contract.RenounceOwnership(&_AccountRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountRegistry *AccountRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountRegistry.Contract.RenounceOwnership(&_AccountRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountRegistry *AccountRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountRegistry *AccountRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountRegistry.Contract.TransferOwnership(&_AccountRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountRegistry *AccountRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountRegistry.Contract.TransferOwnership(&_AccountRegistry.TransactOpts, newOwner)
}

// AccountRegistryAccountRegisteredIterator is returned from FilterAccountRegistered and is used to iterate over the raw logs and unpacked data for AccountRegistered events raised by the AccountRegistry contract.
type AccountRegistryAccountRegisteredIterator struct {
	Event *AccountRegistryAccountRegistered // Event containing the contract specifics and raw log

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
func (it *AccountRegistryAccountRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountRegistryAccountRegistered)
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
		it.Event = new(AccountRegistryAccountRegistered)
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
func (it *AccountRegistryAccountRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountRegistryAccountRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountRegistryAccountRegistered represents a AccountRegistered event raised by the AccountRegistry contract.
type AccountRegistryAccountRegistered struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountRegistered is a free log retrieval operation binding the contract event 0xcd822dc9688e20acea68724a2fbcfe4f3e526d20ecaa37b18fe3047ab377d6a5.
//
// Solidity: event AccountRegistered(address account)
func (_AccountRegistry *AccountRegistryFilterer) FilterAccountRegistered(opts *bind.FilterOpts) (*AccountRegistryAccountRegisteredIterator, error) {

	logs, sub, err := _AccountRegistry.contract.FilterLogs(opts, "AccountRegistered")
	if err != nil {
		return nil, err
	}
	return &AccountRegistryAccountRegisteredIterator{contract: _AccountRegistry.contract, event: "AccountRegistered", logs: logs, sub: sub}, nil
}

// WatchAccountRegistered is a free log subscription operation binding the contract event 0xcd822dc9688e20acea68724a2fbcfe4f3e526d20ecaa37b18fe3047ab377d6a5.
//
// Solidity: event AccountRegistered(address account)
func (_AccountRegistry *AccountRegistryFilterer) WatchAccountRegistered(opts *bind.WatchOpts, sink chan<- *AccountRegistryAccountRegistered) (event.Subscription, error) {

	logs, sub, err := _AccountRegistry.contract.WatchLogs(opts, "AccountRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountRegistryAccountRegistered)
				if err := _AccountRegistry.contract.UnpackLog(event, "AccountRegistered", log); err != nil {
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

// ParseAccountRegistered is a log parse operation binding the contract event 0xcd822dc9688e20acea68724a2fbcfe4f3e526d20ecaa37b18fe3047ab377d6a5.
//
// Solidity: event AccountRegistered(address account)
func (_AccountRegistry *AccountRegistryFilterer) ParseAccountRegistered(log types.Log) (*AccountRegistryAccountRegistered, error) {
	event := new(AccountRegistryAccountRegistered)
	if err := _AccountRegistry.contract.UnpackLog(event, "AccountRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountRegistry contract.
type AccountRegistryOwnershipTransferredIterator struct {
	Event *AccountRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountRegistryOwnershipTransferred)
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
		it.Event = new(AccountRegistryOwnershipTransferred)
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
func (it *AccountRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the AccountRegistry contract.
type AccountRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountRegistry *AccountRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountRegistryOwnershipTransferredIterator{contract: _AccountRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountRegistry *AccountRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountRegistryOwnershipTransferred)
				if err := _AccountRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountRegistry *AccountRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*AccountRegistryOwnershipTransferred, error) {
	event := new(AccountRegistryOwnershipTransferred)
	if err := _AccountRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
