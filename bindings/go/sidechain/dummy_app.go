// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sidechain

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

// DummyAppABI is the input ABI used to generate the binding from.
const DummyAppABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"playerOneDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"playerTwoWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DummyAppBin is the compiled bytecode used for deploying new contracts.
var DummyAppBin = "0x608060405234801561001057600080fd5b506040516104a43803806104a48339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018081905550506104098061009b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635ef061221461003b578063f01a2cd114610116575b600080fd5b6101146004803603604081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561008e57600080fd5b8201836020820111156100a057600080fd5b803590602001918460018302840111640100000000831117156100c257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061015a565b005b6101586004803603602081101561012c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102c4565b005b8173ffffffffffffffffffffffffffffffffffffffff166312a837b43330600154856040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561023557808201518184015260208101905061021a565b50505050905090810190601f1680156102625780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561028457600080fd5b505af1158015610298573d6000803e3d6000fd5b505050506040513d60208110156102ae57600080fd5b8101908080519060200190929190505050505050565b8073ffffffffffffffffffffffffffffffffffffffff166312a837b430336001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200180602001828103825260008152602001602001945050505050602060405180830381600087803b15801561039557600080fd5b505af11580156103a9573d6000803e3d6000fd5b505050506040513d60208110156103bf57600080fd5b8101908080519060200190929190505050505056fea265627a7a72315820142711caa27848d04e32cf972bfb79cc3174cc502fc88f2ae5bc87fdbc7ced4764736f6c634300050f0032"

// DeployDummyApp deploys a new Ethereum contract, binding an instance of DummyApp to it.
func DeployDummyApp(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *DummyApp, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyAppABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DummyAppBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DummyApp{DummyAppCaller: DummyAppCaller{contract: contract}, DummyAppTransactor: DummyAppTransactor{contract: contract}, DummyAppFilterer: DummyAppFilterer{contract: contract}}, nil
}

// DummyApp is an auto generated Go binding around an Ethereum contract.
type DummyApp struct {
	DummyAppCaller     // Read-only binding to the contract
	DummyAppTransactor // Write-only binding to the contract
	DummyAppFilterer   // Log filterer for contract events
}

// DummyAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyAppSession struct {
	Contract     *DummyApp         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyAppCallerSession struct {
	Contract *DummyAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DummyAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyAppTransactorSession struct {
	Contract     *DummyAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DummyAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyAppRaw struct {
	Contract *DummyApp // Generic contract binding to access the raw methods on
}

// DummyAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyAppCallerRaw struct {
	Contract *DummyAppCaller // Generic read-only contract binding to access the raw methods on
}

// DummyAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyAppTransactorRaw struct {
	Contract *DummyAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyApp creates a new instance of DummyApp, bound to a specific deployed contract.
func NewDummyApp(address common.Address, backend bind.ContractBackend) (*DummyApp, error) {
	contract, err := bindDummyApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyApp{DummyAppCaller: DummyAppCaller{contract: contract}, DummyAppTransactor: DummyAppTransactor{contract: contract}, DummyAppFilterer: DummyAppFilterer{contract: contract}}, nil
}

// NewDummyAppCaller creates a new read-only instance of DummyApp, bound to a specific deployed contract.
func NewDummyAppCaller(address common.Address, caller bind.ContractCaller) (*DummyAppCaller, error) {
	contract, err := bindDummyApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyAppCaller{contract: contract}, nil
}

// NewDummyAppTransactor creates a new write-only instance of DummyApp, bound to a specific deployed contract.
func NewDummyAppTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyAppTransactor, error) {
	contract, err := bindDummyApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyAppTransactor{contract: contract}, nil
}

// NewDummyAppFilterer creates a new log filterer instance of DummyApp, bound to a specific deployed contract.
func NewDummyAppFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyAppFilterer, error) {
	contract, err := bindDummyApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyAppFilterer{contract: contract}, nil
}

// bindDummyApp binds a generic wrapper to an already deployed contract.
func bindDummyApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyApp *DummyAppRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyApp.Contract.DummyAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyApp *DummyAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyApp.Contract.DummyAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyApp *DummyAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyApp.Contract.DummyAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyApp *DummyAppCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyApp *DummyAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyApp *DummyAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyApp.Contract.contract.Transact(opts, method, params...)
}

// PlayerOneDeposit is a paid mutator transaction binding the contract method 0x5ef06122.
//
// Solidity: function playerOneDeposit(address token, bytes signature) returns()
func (_DummyApp *DummyAppTransactor) PlayerOneDeposit(opts *bind.TransactOpts, token common.Address, signature []byte) (*types.Transaction, error) {
	return _DummyApp.contract.Transact(opts, "playerOneDeposit", token, signature)
}

// PlayerOneDeposit is a paid mutator transaction binding the contract method 0x5ef06122.
//
// Solidity: function playerOneDeposit(address token, bytes signature) returns()
func (_DummyApp *DummyAppSession) PlayerOneDeposit(token common.Address, signature []byte) (*types.Transaction, error) {
	return _DummyApp.Contract.PlayerOneDeposit(&_DummyApp.TransactOpts, token, signature)
}

// PlayerOneDeposit is a paid mutator transaction binding the contract method 0x5ef06122.
//
// Solidity: function playerOneDeposit(address token, bytes signature) returns()
func (_DummyApp *DummyAppTransactorSession) PlayerOneDeposit(token common.Address, signature []byte) (*types.Transaction, error) {
	return _DummyApp.Contract.PlayerOneDeposit(&_DummyApp.TransactOpts, token, signature)
}

// PlayerTwoWithdraw is a paid mutator transaction binding the contract method 0xf01a2cd1.
//
// Solidity: function playerTwoWithdraw(address token) returns()
func (_DummyApp *DummyAppTransactor) PlayerTwoWithdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _DummyApp.contract.Transact(opts, "playerTwoWithdraw", token)
}

// PlayerTwoWithdraw is a paid mutator transaction binding the contract method 0xf01a2cd1.
//
// Solidity: function playerTwoWithdraw(address token) returns()
func (_DummyApp *DummyAppSession) PlayerTwoWithdraw(token common.Address) (*types.Transaction, error) {
	return _DummyApp.Contract.PlayerTwoWithdraw(&_DummyApp.TransactOpts, token)
}

// PlayerTwoWithdraw is a paid mutator transaction binding the contract method 0xf01a2cd1.
//
// Solidity: function playerTwoWithdraw(address token) returns()
func (_DummyApp *DummyAppTransactorSession) PlayerTwoWithdraw(token common.Address) (*types.Transaction, error) {
	return _DummyApp.Contract.PlayerTwoWithdraw(&_DummyApp.TransactOpts, token)
}
