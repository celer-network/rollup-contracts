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



// DepositWithdrawManagerABI is the input ABI used to generate the binding from.
const DepositWithdrawManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupChainAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_transitionEvaluatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transition\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transitionIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.TransitionInclusionProof\",\"name\":\"inclusionProof\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.IncludedTransition\",\"name\":\"_includedTransition\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositWithdrawManagerBin is the compiled bytecode used for deploying new contracts.
var DepositWithdrawManagerBin = "0x60806040523480156200001157600080fd5b50604051620010be380380620010be833981810160405281019062000037919062000119565b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620001bd565b6000815190506200011381620001a3565b92915050565b6000806000606084860312156200012f57600080fd5b60006200013f8682870162000102565b9350506020620001528682870162000102565b9250506040620001658682870162000102565b9150509250925092565b60006200017c8262000183565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620001ae816200016f565b8114620001ba57600080fd5b50565b610ef180620001cd6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806347e7ef241461003b57806366fb5de314610057575b600080fd5b61005560048036038101906100509190610839565b610073565b005b610071600480360381019061006c919061089e565b610180565b005b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016100b093929190610b8d565b602060405180830381600087803b1580156100ca57600080fd5b505af11580156100de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101029190610875565b610141576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013890610c0f565b60405180910390fd5b7f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6233838360405161017493929190610b8d565b60405180910390a15050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a14f166826040518263ffffffff1660e01b81526004016101da9190610c4f565b60206040518083038186803b1580156101f257600080fd5b505afa158015610206573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022a9190610875565b610269576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026090610c2f565b60405180910390fd5b610271610471565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc354f6783600001516040518263ffffffff1660e01b81526004016102d09190610bed565b60006040518083038186803b1580156102e857600080fd5b505afa1580156102fc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061032591906108df565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639bf53f1883606001516040518263ffffffff1660e01b81526004016103889190610c71565b60206040518083038186803b1580156103a057600080fd5b505afa1580156103b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d89190610810565b90508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb3384608001516040518363ffffffff1660e01b8152600401610419929190610bc4565b602060405180830381600087803b15801561043357600080fd5b505af1158015610447573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046b9190610875565b50505050565b6040518060c001604052806000815260200160008019168152602001600081526020016000815260200160008152602001606081525090565b6000813590506104b981610e5f565b92915050565b6000815190506104ce81610e5f565b92915050565b600082601f8301126104e557600080fd5b81356104f86104f382610cb9565b610c8c565b9150818183526020840193506020810190508385602084028201111561051d57600080fd5b60005b8381101561054d5781610533888261056c565b845260208401935060208301925050600181019050610520565b5050505092915050565b60008151905061056681610e76565b92915050565b60008135905061057b81610e8d565b92915050565b60008151905061059081610e8d565b92915050565b600082601f8301126105a757600080fd5b81356105ba6105b582610ce1565b610c8c565b915080825260208301602083018583830111156105d657600080fd5b6105e1838284610e0c565b50505092915050565b600082601f8301126105fb57600080fd5b815161060e61060982610ce1565b610c8c565b9150808252602083016020830185838301111561062a57600080fd5b610635838284610e1b565b50505092915050565b60006040828403121561065057600080fd5b61065a6040610c8c565b9050600082013567ffffffffffffffff81111561067657600080fd5b61068284828501610596565b600083015250602082013567ffffffffffffffff8111156106a257600080fd5b6106ae848285016106ba565b60208301525092915050565b6000606082840312156106cc57600080fd5b6106d66060610c8c565b905060006106e6848285016107e6565b60008301525060206106fa848285016107e6565b602083015250604082013567ffffffffffffffff81111561071a57600080fd5b610726848285016104d4565b60408301525092915050565b600060c0828403121561074457600080fd5b61074e60c0610c8c565b9050600061075e848285016107fb565b600083015250602061077284828501610581565b6020830152506040610786848285016107fb565b604083015250606061079a848285016107fb565b60608301525060806107ae848285016107fb565b60808301525060a082015167ffffffffffffffff8111156107ce57600080fd5b6107da848285016105ea565b60a08301525092915050565b6000813590506107f581610ea4565b92915050565b60008151905061080a81610ea4565b92915050565b60006020828403121561082257600080fd5b6000610830848285016104bf565b91505092915050565b6000806040838503121561084c57600080fd5b600061085a858286016104aa565b925050602061086b858286016107e6565b9150509250929050565b60006020828403121561088757600080fd5b600061089584828501610557565b91505092915050565b6000602082840312156108b057600080fd5b600082013567ffffffffffffffff8111156108ca57600080fd5b6108d68482850161063e565b91505092915050565b6000602082840312156108f157600080fd5b600082015167ffffffffffffffff81111561090b57600080fd5b61091784828501610732565b91505092915050565b600061092c83836109b4565b60208301905092915050565b61094181610dd6565b82525050565b61095081610d84565b82525050565b600061096182610d1d565b61096b8185610d40565b935061097683610d0d565b8060005b838110156109a757815161098e8882610920565b975061099983610d33565b92505060018101905061097a565b5085935050505092915050565b6109bd81610da2565b82525050565b60006109ce82610d28565b6109d88185610d51565b93506109e8818560208601610e1b565b6109f181610e4e565b840191505092915050565b6000610a0782610d28565b610a118185610d62565b9350610a21818560208601610e1b565b610a2a81610e4e565b840191505092915050565b6000610a42601583610d73565b91507f546f6b656e207472616e73666572206661696c656400000000000000000000006000830152602082019050919050565b6000610a82602483610d73565b91507f5769746864726177207472616e736974696f6e206d75737420626520696e636c60008301527f75646564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006040830160008301518482036000860152610af882826109c3565b91505060208301518482036020860152610b128282610b1f565b9150508091505092915050565b6000606083016000830151610b376000860182610b6f565b506020830151610b4a6020860182610b6f565b5060408301518482036040860152610b628282610956565b9150508091505092915050565b610b7881610dcc565b82525050565b610b8781610dcc565b82525050565b6000606082019050610ba26000830186610938565b610baf6020830185610947565b610bbc6040830184610b7e565b949350505050565b6000604082019050610bd96000830185610938565b610be66020830184610b7e565b9392505050565b60006020820190508181036000830152610c0781846109fc565b905092915050565b60006020820190508181036000830152610c2881610a35565b9050919050565b60006020820190508181036000830152610c4881610a75565b9050919050565b60006020820190508181036000830152610c698184610adb565b905092915050565b6000602082019050610c866000830184610b7e565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610caf57600080fd5b8060405250919050565b600067ffffffffffffffff821115610cd057600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610cf857600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610d8f82610dac565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610de182610de8565b9050919050565b6000610df382610dfa565b9050919050565b6000610e0582610dac565b9050919050565b82818337600083830152505050565b60005b83811015610e39578082015181840152602081019050610e1e565b83811115610e48576000848401525b50505050565b6000601f19601f8301169050919050565b610e6881610d84565b8114610e7357600080fd5b50565b610e7f81610d96565b8114610e8a57600080fd5b50565b610e9681610da2565b8114610ea157600080fd5b50565b610ead81610dcc565b8114610eb857600080fd5b5056fea2646970667358221220469e4ccbb4972ea98d4cec0860d23b2702270c3c9ee236bd8170ad198fa40b2564736f6c63430006060033"

// DeployDepositWithdrawManager deploys a new Ethereum contract, binding an instance of DepositWithdrawManager to it.
func DeployDepositWithdrawManager(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupChainAddress common.Address, _transitionEvaluatorAddress common.Address, _tokenRegistryAddress common.Address) (common.Address, *types.Transaction, *DepositWithdrawManager, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositWithdrawManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositWithdrawManagerBin), backend, _rollupChainAddress, _transitionEvaluatorAddress, _tokenRegistryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositWithdrawManager{DepositWithdrawManagerCaller: DepositWithdrawManagerCaller{contract: contract}, DepositWithdrawManagerTransactor: DepositWithdrawManagerTransactor{contract: contract}, DepositWithdrawManagerFilterer: DepositWithdrawManagerFilterer{contract: contract}}, nil
}

// DepositWithdrawManager is an auto generated Go binding around an Ethereum contract.
type DepositWithdrawManager struct {
	DepositWithdrawManagerCaller     // Read-only binding to the contract
	DepositWithdrawManagerTransactor // Write-only binding to the contract
	DepositWithdrawManagerFilterer   // Log filterer for contract events
}

// DepositWithdrawManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositWithdrawManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositWithdrawManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositWithdrawManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositWithdrawManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositWithdrawManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositWithdrawManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositWithdrawManagerSession struct {
	Contract     *DepositWithdrawManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DepositWithdrawManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositWithdrawManagerCallerSession struct {
	Contract *DepositWithdrawManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DepositWithdrawManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositWithdrawManagerTransactorSession struct {
	Contract     *DepositWithdrawManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DepositWithdrawManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositWithdrawManagerRaw struct {
	Contract *DepositWithdrawManager // Generic contract binding to access the raw methods on
}

// DepositWithdrawManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositWithdrawManagerCallerRaw struct {
	Contract *DepositWithdrawManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DepositWithdrawManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositWithdrawManagerTransactorRaw struct {
	Contract *DepositWithdrawManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositWithdrawManager creates a new instance of DepositWithdrawManager, bound to a specific deployed contract.
func NewDepositWithdrawManager(address common.Address, backend bind.ContractBackend) (*DepositWithdrawManager, error) {
	contract, err := bindDepositWithdrawManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositWithdrawManager{DepositWithdrawManagerCaller: DepositWithdrawManagerCaller{contract: contract}, DepositWithdrawManagerTransactor: DepositWithdrawManagerTransactor{contract: contract}, DepositWithdrawManagerFilterer: DepositWithdrawManagerFilterer{contract: contract}}, nil
}

// NewDepositWithdrawManagerCaller creates a new read-only instance of DepositWithdrawManager, bound to a specific deployed contract.
func NewDepositWithdrawManagerCaller(address common.Address, caller bind.ContractCaller) (*DepositWithdrawManagerCaller, error) {
	contract, err := bindDepositWithdrawManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositWithdrawManagerCaller{contract: contract}, nil
}

// NewDepositWithdrawManagerTransactor creates a new write-only instance of DepositWithdrawManager, bound to a specific deployed contract.
func NewDepositWithdrawManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositWithdrawManagerTransactor, error) {
	contract, err := bindDepositWithdrawManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositWithdrawManagerTransactor{contract: contract}, nil
}

// NewDepositWithdrawManagerFilterer creates a new log filterer instance of DepositWithdrawManager, bound to a specific deployed contract.
func NewDepositWithdrawManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositWithdrawManagerFilterer, error) {
	contract, err := bindDepositWithdrawManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositWithdrawManagerFilterer{contract: contract}, nil
}

// bindDepositWithdrawManager binds a generic wrapper to an already deployed contract.
func bindDepositWithdrawManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositWithdrawManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositWithdrawManager *DepositWithdrawManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositWithdrawManager.Contract.DepositWithdrawManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositWithdrawManager *DepositWithdrawManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.DepositWithdrawManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositWithdrawManager *DepositWithdrawManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.DepositWithdrawManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositWithdrawManager *DepositWithdrawManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositWithdrawManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DepositWithdrawManager.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Deposit(&_DepositWithdrawManager.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Deposit(&_DepositWithdrawManager.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x66fb5de3.
//
// Solidity: function withdraw(DataTypesIncludedTransition _includedTransition) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactor) Withdraw(opts *bind.TransactOpts, _includedTransition DataTypesIncludedTransition) (*types.Transaction, error) {
	return _DepositWithdrawManager.contract.Transact(opts, "withdraw", _includedTransition)
}

// Withdraw is a paid mutator transaction binding the contract method 0x66fb5de3.
//
// Solidity: function withdraw(DataTypesIncludedTransition _includedTransition) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerSession) Withdraw(_includedTransition DataTypesIncludedTransition) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Withdraw(&_DepositWithdrawManager.TransactOpts, _includedTransition)
}

// Withdraw is a paid mutator transaction binding the contract method 0x66fb5de3.
//
// Solidity: function withdraw(DataTypesIncludedTransition _includedTransition) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorSession) Withdraw(_includedTransition DataTypesIncludedTransition) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Withdraw(&_DepositWithdrawManager.TransactOpts, _includedTransition)
}

// DepositWithdrawManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the DepositWithdrawManager contract.
type DepositWithdrawManagerDepositIterator struct {
	Event *DepositWithdrawManagerDeposit // Event containing the contract specifics and raw log

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
func (it *DepositWithdrawManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositWithdrawManagerDeposit)
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
		it.Event = new(DepositWithdrawManagerDeposit)
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
func (it *DepositWithdrawManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositWithdrawManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositWithdrawManagerDeposit represents a Deposit event raised by the DepositWithdrawManager contract.
type DepositWithdrawManagerDeposit struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*DepositWithdrawManagerDepositIterator, error) {

	logs, sub, err := _DepositWithdrawManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &DepositWithdrawManagerDepositIterator{contract: _DepositWithdrawManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *DepositWithdrawManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _DepositWithdrawManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositWithdrawManagerDeposit)
				if err := _DepositWithdrawManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) ParseDeposit(log types.Log) (*DepositWithdrawManagerDeposit, error) {
	event := new(DepositWithdrawManagerDeposit)
	if err := _DepositWithdrawManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}
