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

// RollupMerkleUtilsABI is the input ABI used to generate the binding from.
const RollupMerkleUtilsABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"defaultHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getChildren\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getLeftSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_dataBlocks\",\"type\":\"bytes[]\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_intVal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNthBitFromRight\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getRightSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"getSiblings\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"setMerkleRootAndHeight\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"storeLeaf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_leftChild\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rightChild\",\"type\":\"bytes32\"}],\"name\":\"storeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"updateLeaf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyAndStore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupMerkleUtilsBin is the compiled bytecode used for deploying new contracts.
var RollupMerkleUtilsBin = "0x60806040523480156200001157600080fd5b50620000226200002860201b60201c565b62000194565b60006040516020016200003c91906200014f565b6040516020818303038152906040528051906020012060008060a081106200006057fe5b01819055506000600190505b60a0811015620000e65760006001820360a081106200008757fe5b015460006001830360a081106200009a57fe5b0154604051602001620000af9291906200011f565b6040516020818303038152906040528051906020012060008260a08110620000d357fe5b018190555080806001019150506200006c565b50565b620000fe620000f8826200016c565b62000180565b82525050565b62000119620001138262000176565b6200018a565b82525050565b60006200012d8285620000e9565b6020820191506200013f8284620000e9565b6020820191508190509392505050565b60006200015d828462000104565b60208201915081905092915050565b6000819050919050565b6000819050919050565b6000819050919050565b6000819050919050565b6112f580620001a46000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80635ca1e165116100a2578063d37684ff11610071578063d37684ff146102b2578063db0787cb146102e3578063df7c726314610313578063e913e47f14610343578063fd54b228146103735761010b565b80635ca1e1651461024057806363327f891461025e5780639c0de5201461027a578063c3b45234146102965761010b565b806340ff34ef116100de57806340ff34ef146101a85780634359356d146101d857806348419ad8146101f45780635c22b6d9146102245761010b565b8063101b166c14610110578063158933ad14610140578063272684b51461015c57806330d90a7614610178575b600080fd5b61012a60048036036101259190810190610e50565b610392565b6040516101379190611083565b60405180910390f35b61015a60048036036101559190810190610dd1565b610471565b005b61017660048036036101719190810190610cda565b61048e565b005b610192600480360361018d9190810190610c47565b6104ab565b60405161019f91906110a5565b60405180910390f35b6101c260048036036101bd9190810190610b8a565b610535565b6040516101cf91906110c0565b60405180910390f35b6101f260048036036101ed9190810190610dd1565b61079b565b005b61020e60048036036102099190810190610e50565b6107fd565b60405161021b91906110c0565b60405180910390f35b61023e60048036036102399190810190610cda565b610815565b005b61024861082d565b60405161025591906110c0565b60405180910390f35b61027860048036036102739190810190610bf8565b61083a565b005b610294600480360361028f9190810190610d16565b610885565b005b6102b060048036036102ab9190810190610d7d565b610924565b005b6102cc60048036036102c79190810190610bcf565b610941565b6040516102da9291906110db565b60405180910390f35b6102fd60048036036102f89190810190610e79565b61098b565b60405161030a919061114d565b60405180910390f35b61032d60048036036103289190810190610bcf565b61099c565b60405161033a91906110c0565b60405180910390f35b61035d60048036036103589190810190610bcf565b6109cb565b60405161036a91906110c0565b60405180910390f35b61037b6109fa565b604051610389929190611104565b60405180910390f35b60608060a0600101546040519080825280602002602001820160405280156103c95781602001602082028038833980820191505090505b509050600060a0600001549050600060a06001015490505b60008111156104665760006001820390506000806103fe85610941565b91509150600061040e898561098b565b60ff161415610438578194508086848151811061042757fe5b602002602001018181525050610455565b8094508186848151811061044857fe5b6020026020010181815250505b5050508080600190039150506103e1565b508192505050919050565b600083805190602001209050610488818484610885565b50505050565b606061049982610392565b90506104a6838383610885565b505050565b6000808480519060200120905060008090505b83518110156105265760008482815181106104d557fe5b6020026020010151905060006104eb878461098b565b905060008160ff16141561050a576105038483610a0c565b9350610517565b6105148285610a0c565b93505b505080806001019150506104be565b50858114915050949350505050565b60008083839050905060008090506060600183016040519080825280602002602001820160405280156105775781602001602082028038833980820191505090505b50905060008090505b868690508110156106265786868281811061059757fe5b90506020028101803560016020038336030381126105b457600080fd5b8083019250508135905060208201915067ffffffffffffffff8111156105d957600080fd5b6001810236038213156105eb57600080fd5b6040516105f992919061106a565b604051809103902082828151811061060d57fe5b6020026020010181815250508080600101915050610580565b506001868690501415610652578060008151811061064057fe5b60200260200101519350505050610795565b60016002848161065e57fe5b0614156106935760008260a0811061067257fe5b015481848151811061068057fe5b6020026020010181815250506001830192505b5b600183111561077a5760018201915060008090505b600284816106b357fe5b04811015610719576106f48260028302815181106106cd57fe5b60200260200101518360016002850201815181106106e757fe5b6020026020010151610a0c565b82828151811061070057fe5b60200260200101818152505080806001019150506106a9565b506002838161072457fe5b04925060016002848161073357fe5b06148015610742575060018314155b156107755760008260a0811061075457fe5b015481848151811061076257fe5b6020026020010181815250506001830192505b610694565b8060008151811061078757fe5b602002602001015193505050505b92915050565b600060a06000015490506107b0848484610471565b8060a060000154146107f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ee9061112d565b60405180910390fd5b50505050565b60008160a0811061080a57fe5b016000915090505481565b8160a0600001819055508060a0600101819055505050565b600060a060000154905090565b8160a0600201600061084b8661099c565b8152602001908152602001600020819055508060a0600201600061086e866109cb565b815260200190815260200160002081905550505050565b600083905060008090505b8251811015610913576000808483815181106108a857fe5b6020026020010151905060006108be878561098b565b905060008160ff1614156108e8576108d68583610a0c565b92506108e383868461083a565b610900565b6108f28286610a0c565b92506108ff83838761083a565b5b8294505050508080600101915050610890565b508060a06000018190555050505050565b606061092f82610392565b905061093c838383610471565b505050565b60008060a060020160006109548561099c565b81526020019081526020016000205460a06002016000610973866109cb565b81526020019081526020016000205491509150915091565b600060018284901c16905092915050565b60007f011111111111111111111111111111111111111111111111111111111111111160001b82169050919050565b60007f100000000000000000000000000000000000000000000000000000000000000060001b82179050919050565b60a08060000154908060010154905082565b60008282604051602001610a2192919061103e565b60405160208183030381529060405280519060200120905092915050565b600082601f830112610a5057600080fd5b8135610a63610a5e82611195565b611168565b91508181835260208401935060208101905083856020840282011115610a8857600080fd5b60005b83811015610ab85781610a9e8882610b0c565b845260208401935060208301925050600181019050610a8b565b5050505092915050565b60008083601f840112610ad457600080fd5b8235905067ffffffffffffffff811115610aed57600080fd5b602083019150836020820283011115610b0557600080fd5b9250929050565b600081359050610b1b81611284565b92915050565b600082601f830112610b3257600080fd5b8135610b45610b40826111bd565b611168565b91508082526020830160208301858383011115610b6157600080fd5b610b6c83828461126b565b50505092915050565b600081359050610b848161129b565b92915050565b60008060208385031215610b9d57600080fd5b600083013567ffffffffffffffff811115610bb757600080fd5b610bc385828601610ac2565b92509250509250929050565b600060208284031215610be157600080fd5b6000610bef84828501610b0c565b91505092915050565b600080600060608486031215610c0d57600080fd5b6000610c1b86828701610b0c565b9350506020610c2c86828701610b0c565b9250506040610c3d86828701610b0c565b9150509250925092565b60008060008060808587031215610c5d57600080fd5b6000610c6b87828801610b0c565b945050602085013567ffffffffffffffff811115610c8857600080fd5b610c9487828801610b21565b9350506040610ca587828801610b75565b925050606085013567ffffffffffffffff811115610cc257600080fd5b610cce87828801610a3f565b91505092959194509250565b60008060408385031215610ced57600080fd5b6000610cfb85828601610b0c565b9250506020610d0c85828601610b75565b9150509250929050565b600080600060608486031215610d2b57600080fd5b6000610d3986828701610b0c565b9350506020610d4a86828701610b75565b925050604084013567ffffffffffffffff811115610d6757600080fd5b610d7386828701610a3f565b9150509250925092565b60008060408385031215610d9057600080fd5b600083013567ffffffffffffffff811115610daa57600080fd5b610db685828601610b21565b9250506020610dc785828601610b75565b9150509250929050565b600080600060608486031215610de657600080fd5b600084013567ffffffffffffffff811115610e0057600080fd5b610e0c86828701610b21565b9350506020610e1d86828701610b75565b925050604084013567ffffffffffffffff811115610e3a57600080fd5b610e4686828701610a3f565b9150509250925092565b600060208284031215610e6257600080fd5b6000610e7084828501610b75565b91505092915050565b60008060408385031215610e8c57600080fd5b6000610e9a85828601610b75565b9250506020610eab85828601610b75565b9150509250929050565b6000610ec18383610f3a565b60208301905092915050565b6000610ed8826111f9565b610ee28185611211565b9350610eed836111e9565b8060005b83811015610f1e578151610f058882610eb5565b9750610f1083611204565b925050600181019050610ef1565b5085935050505092915050565b610f348161123e565b82525050565b610f438161124a565b82525050565b610f528161124a565b82525050565b610f69610f648261124a565b61127a565b82525050565b6000610f7b8385611222565b9350610f8883858461126b565b82840190509392505050565b6000610fa160568361122d565b91507f4661696c65642073616d6520726f6f7420766572696669636174696f6e20636860008301527f65636b2120546869732077617320616e20696e636c7573696f6e2070726f6f6660208301527f20666f72206120646966666572656e74207472656521000000000000000000006040830152606082019050919050565b61102981611254565b82525050565b6110388161125e565b82525050565b600061104a8285610f58565b60208201915061105a8284610f58565b6020820191508190509392505050565b6000611077828486610f6f565b91508190509392505050565b6000602082019050818103600083015261109d8184610ecd565b905092915050565b60006020820190506110ba6000830184610f2b565b92915050565b60006020820190506110d56000830184610f49565b92915050565b60006040820190506110f06000830185610f49565b6110fd6020830184610f49565b9392505050565b60006040820190506111196000830185610f49565b6111266020830184611020565b9392505050565b6000602082019050818103600083015261114681610f94565b9050919050565b6000602082019050611162600083018461102f565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561118b57600080fd5b8060405250919050565b600067ffffffffffffffff8211156111ac57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156111d457600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b6000819050919050565b61128d8161124a565b811461129857600080fd5b50565b6112a481611254565b81146112af57600080fd5b5056fea365627a7a723158200e73ea472019f2e36533b2b75cf3c34ba73f61ad26470f244539b885076c75fb6c6578706572696d656e74616cf564736f6c634300050f0040"

// DeployRollupMerkleUtils deploys a new Ethereum contract, binding an instance of RollupMerkleUtils to it.
func DeployRollupMerkleUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupMerkleUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupMerkleUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupMerkleUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupMerkleUtils{RollupMerkleUtilsCaller: RollupMerkleUtilsCaller{contract: contract}, RollupMerkleUtilsTransactor: RollupMerkleUtilsTransactor{contract: contract}, RollupMerkleUtilsFilterer: RollupMerkleUtilsFilterer{contract: contract}}, nil
}

// RollupMerkleUtils is an auto generated Go binding around an Ethereum contract.
type RollupMerkleUtils struct {
	RollupMerkleUtilsCaller     // Read-only binding to the contract
	RollupMerkleUtilsTransactor // Write-only binding to the contract
	RollupMerkleUtilsFilterer   // Log filterer for contract events
}

// RollupMerkleUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupMerkleUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMerkleUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupMerkleUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMerkleUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupMerkleUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMerkleUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupMerkleUtilsSession struct {
	Contract     *RollupMerkleUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RollupMerkleUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupMerkleUtilsCallerSession struct {
	Contract *RollupMerkleUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RollupMerkleUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupMerkleUtilsTransactorSession struct {
	Contract     *RollupMerkleUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RollupMerkleUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupMerkleUtilsRaw struct {
	Contract *RollupMerkleUtils // Generic contract binding to access the raw methods on
}

// RollupMerkleUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupMerkleUtilsCallerRaw struct {
	Contract *RollupMerkleUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// RollupMerkleUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupMerkleUtilsTransactorRaw struct {
	Contract *RollupMerkleUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupMerkleUtils creates a new instance of RollupMerkleUtils, bound to a specific deployed contract.
func NewRollupMerkleUtils(address common.Address, backend bind.ContractBackend) (*RollupMerkleUtils, error) {
	contract, err := bindRollupMerkleUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupMerkleUtils{RollupMerkleUtilsCaller: RollupMerkleUtilsCaller{contract: contract}, RollupMerkleUtilsTransactor: RollupMerkleUtilsTransactor{contract: contract}, RollupMerkleUtilsFilterer: RollupMerkleUtilsFilterer{contract: contract}}, nil
}

// NewRollupMerkleUtilsCaller creates a new read-only instance of RollupMerkleUtils, bound to a specific deployed contract.
func NewRollupMerkleUtilsCaller(address common.Address, caller bind.ContractCaller) (*RollupMerkleUtilsCaller, error) {
	contract, err := bindRollupMerkleUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMerkleUtilsCaller{contract: contract}, nil
}

// NewRollupMerkleUtilsTransactor creates a new write-only instance of RollupMerkleUtils, bound to a specific deployed contract.
func NewRollupMerkleUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupMerkleUtilsTransactor, error) {
	contract, err := bindRollupMerkleUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMerkleUtilsTransactor{contract: contract}, nil
}

// NewRollupMerkleUtilsFilterer creates a new log filterer instance of RollupMerkleUtils, bound to a specific deployed contract.
func NewRollupMerkleUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupMerkleUtilsFilterer, error) {
	contract, err := bindRollupMerkleUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupMerkleUtilsFilterer{contract: contract}, nil
}

// bindRollupMerkleUtils binds a generic wrapper to an already deployed contract.
func bindRollupMerkleUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupMerkleUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMerkleUtils *RollupMerkleUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupMerkleUtils.Contract.RollupMerkleUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMerkleUtils *RollupMerkleUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.RollupMerkleUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMerkleUtils *RollupMerkleUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.RollupMerkleUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMerkleUtils *RollupMerkleUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupMerkleUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.contract.Transact(opts, method, params...)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) DefaultHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "defaultHashes", arg0)
	return *ret0, err
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.DefaultHashes(&_RollupMerkleUtils.CallOpts, arg0)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.DefaultHashes(&_RollupMerkleUtils.CallOpts, arg0)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) constant returns(bytes32, bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) GetChildren(opts *bind.CallOpts, _parent [32]byte) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RollupMerkleUtils.contract.Call(opts, out, "getChildren", _parent)
	return *ret0, *ret1, err
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) constant returns(bytes32, bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _RollupMerkleUtils.Contract.GetChildren(&_RollupMerkleUtils.CallOpts, _parent)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) constant returns(bytes32, bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _RollupMerkleUtils.Contract.GetChildren(&_RollupMerkleUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) GetLeftSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "getLeftSiblingKey", _parent)
	return *ret0, err
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetLeftSiblingKey(&_RollupMerkleUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetLeftSiblingKey(&_RollupMerkleUtils.CallOpts, _parent)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) GetMerkleRoot(opts *bind.CallOpts, _dataBlocks [][]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "getMerkleRoot", _dataBlocks)
	return *ret0, err
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetMerkleRoot(&_RollupMerkleUtils.CallOpts, _dataBlocks)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetMerkleRoot(&_RollupMerkleUtils.CallOpts, _dataBlocks)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) constant returns(uint8)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) GetNthBitFromRight(opts *bind.CallOpts, _intVal *big.Int, _index *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "getNthBitFromRight", _intVal, _index)
	return *ret0, err
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) constant returns(uint8)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _RollupMerkleUtils.Contract.GetNthBitFromRight(&_RollupMerkleUtils.CallOpts, _intVal, _index)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) constant returns(uint8)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _RollupMerkleUtils.Contract.GetNthBitFromRight(&_RollupMerkleUtils.CallOpts, _intVal, _index)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) GetRightSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "getRightSiblingKey", _parent)
	return *ret0, err
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetRightSiblingKey(&_RollupMerkleUtils.CallOpts, _parent)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetRightSiblingKey(&_RollupMerkleUtils.CallOpts, _parent)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "getRoot")
	return *ret0, err
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) GetRoot() ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetRoot(&_RollupMerkleUtils.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() constant returns(bytes32)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) GetRoot() ([32]byte, error) {
	return _RollupMerkleUtils.Contract.GetRoot(&_RollupMerkleUtils.CallOpts)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) constant returns(bytes32[])
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) GetSiblings(opts *bind.CallOpts, _path *big.Int) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "getSiblings", _path)
	return *ret0, err
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) constant returns(bytes32[])
func (_RollupMerkleUtils *RollupMerkleUtilsSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _RollupMerkleUtils.Contract.GetSiblings(&_RollupMerkleUtils.CallOpts, _path)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) constant returns(bytes32[])
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _RollupMerkleUtils.Contract.GetSiblings(&_RollupMerkleUtils.CallOpts, _path)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() constant returns(bytes32 root, uint256 height)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) Tree(opts *bind.CallOpts) (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	ret := new(struct {
		Root   [32]byte
		Height *big.Int
	})
	out := ret
	err := _RollupMerkleUtils.contract.Call(opts, out, "tree")
	return *ret, err
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() constant returns(bytes32 root, uint256 height)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _RollupMerkleUtils.Contract.Tree(&_RollupMerkleUtils.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() constant returns(bytes32 root, uint256 height)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _RollupMerkleUtils.Contract.Tree(&_RollupMerkleUtils.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bool)
func (_RollupMerkleUtils *RollupMerkleUtilsCaller) Verify(opts *bind.CallOpts, _root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RollupMerkleUtils.contract.Call(opts, out, "verify", _root, _dataBlock, _path, _siblings)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bool)
func (_RollupMerkleUtils *RollupMerkleUtilsSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _RollupMerkleUtils.Contract.Verify(&_RollupMerkleUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bool)
func (_RollupMerkleUtils *RollupMerkleUtilsCallerSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _RollupMerkleUtils.Contract.Verify(&_RollupMerkleUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactor) SetMerkleRootAndHeight(opts *bind.TransactOpts, _root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.contract.Transact(opts, "setMerkleRootAndHeight", _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.SetMerkleRootAndHeight(&_RollupMerkleUtils.TransactOpts, _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.SetMerkleRootAndHeight(&_RollupMerkleUtils.TransactOpts, _root, _height)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactor) Store(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.contract.Transact(opts, "store", _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.Store(&_RollupMerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.Store(&_RollupMerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactor) StoreLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.contract.Transact(opts, "storeLeaf", _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.StoreLeaf(&_RollupMerkleUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.StoreLeaf(&_RollupMerkleUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactor) StoreNode(opts *bind.TransactOpts, _parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.contract.Transact(opts, "storeNode", _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.StoreNode(&_RollupMerkleUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.StoreNode(&_RollupMerkleUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactor) Update(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.contract.Transact(opts, "update", _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.Update(&_RollupMerkleUtils.TransactOpts, _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.Update(&_RollupMerkleUtils.TransactOpts, _dataBlock, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactor) UpdateLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.contract.Transact(opts, "updateLeaf", _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.UpdateLeaf(&_RollupMerkleUtils.TransactOpts, _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.UpdateLeaf(&_RollupMerkleUtils.TransactOpts, _leaf, _path)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactor) VerifyAndStore(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.contract.Transact(opts, "verifyAndStore", _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.VerifyAndStore(&_RollupMerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_RollupMerkleUtils *RollupMerkleUtilsTransactorSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RollupMerkleUtils.Contract.VerifyAndStore(&_RollupMerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}
