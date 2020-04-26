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

// MerkleUtilsABI is the input ABI used to generate the binding from.
const MerkleUtilsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"defaultHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getChildren\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getLeftSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_dataBlocks\",\"type\":\"bytes[]\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_intVal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNthBitFromRight\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getRightSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"getSiblings\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"setMerkleRootAndHeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"storeLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_leftChild\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rightChild\",\"type\":\"bytes32\"}],\"name\":\"storeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"updateLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyAndStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MerkleUtilsBin is the compiled bytecode used for deploying new contracts.
var MerkleUtilsBin = "0x60806040523480156200001157600080fd5b50620000226200002860201b60201c565b62000194565b60006040516020016200003c91906200014f565b6040516020818303038152906040528051906020012060008060a081106200006057fe5b01819055506000600190505b60a0811015620000e65760006001820360a081106200008757fe5b015460006001830360a081106200009a57fe5b0154604051602001620000af9291906200011f565b6040516020818303038152906040528051906020012060008260a08110620000d357fe5b018190555080806001019150506200006c565b50565b620000fe620000f8826200016c565b62000180565b82525050565b62000119620001138262000176565b6200018a565b82525050565b60006200012d8285620000e9565b6020820191506200013f8284620000e9565b6020820191508190509392505050565b60006200015d828462000104565b60208201915081905092915050565b6000819050919050565b6000819050919050565b6000819050919050565b6000819050919050565b61132b80620001a46000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80635ca1e165116100a2578063d37684ff11610071578063d37684ff146102b2578063db0787cb146102e3578063df7c726314610313578063e913e47f14610343578063fd54b228146103735761010b565b80635ca1e1651461024057806363327f891461025e5780639c0de5201461027a578063c3b45234146102965761010b565b806340ff34ef116100de57806340ff34ef146101a85780634359356d146101d857806348419ad8146101f45780635c22b6d9146102245761010b565b8063101b166c14610110578063158933ad14610140578063272684b51461015c57806330d90a7614610178575b600080fd5b61012a60048036038101906101259190610e3c565b610392565b604051610137919061106f565b60405180910390f35b61015a60048036038101906101559190610dbd565b610488565b005b61017660048036038101906101719190610cc6565b6104a5565b005b610192600480360381019061018d9190610c33565b6104c2565b60405161019f9190611091565b60405180910390f35b6101c260048036038101906101bd9190610b76565b61054c565b6040516101cf91906110ac565b60405180910390f35b6101f260048036038101906101ed9190610dbd565b610787565b005b61020e60048036038101906102099190610e3c565b6107e9565b60405161021b91906110ac565b60405180910390f35b61023e60048036038101906102399190610cc6565b610801565b005b610248610819565b60405161025591906110ac565b60405180910390f35b61027860048036038101906102739190610be4565b610826565b005b610294600480360381019061028f9190610d02565b610871565b005b6102b060048036038101906102ab9190610d69565b610910565b005b6102cc60048036038101906102c79190610bbb565b61092d565b6040516102da9291906110c7565b60405180910390f35b6102fd60048036038101906102f89190610e65565b610977565b60405161030a9190611139565b60405180910390f35b61032d60048036038101906103289190610bbb565b610988565b60405161033a91906110ac565b60405180910390f35b61035d60048036038101906103589190610bbb565b6109b7565b60405161036a91906110ac565b60405180910390f35b61037b6109e6565b6040516103899291906110f0565b60405180910390f35b60608060a06001015467ffffffffffffffff811180156103b157600080fd5b506040519080825280602002602001820160405280156103e05781602001602082028036833780820191505090505b509050600060a0600001549050600060a06001015490505b600081111561047d5760006001820390506000806104158561092d565b9150915060006104258985610977565b60ff16141561044f578194508086848151811061043e57fe5b60200260200101818152505061046c565b8094508186848151811061045f57fe5b6020026020010181815250505b5050508080600190039150506103f8565b508192505050919050565b60008380519060200120905061049f818484610871565b50505050565b60606104b082610392565b90506104bd838383610871565b505050565b6000808480519060200120905060008090505b835181101561053d5760008482815181106104ec57fe5b6020026020010151905060006105028784610977565b905060008160ff1614156105215761051a84836109f8565b935061052e565b61052b82856109f8565b93505b505080806001019150506104d5565b50858114915050949350505050565b600080838390509050600080905060606001830167ffffffffffffffff8111801561057657600080fd5b506040519080825280602002602001820160405280156105a55781602001602082028036833780820191505090505b50905060008090505b86869050811015610612578686828181106105c557fe5b90506020028101906105d79190611154565b6040516105e5929190611056565b60405180910390208282815181106105f957fe5b60200260200101818152505080806001019150506105ae565b50600186869050141561063e578060008151811061062c57fe5b60200260200101519350505050610781565b60016002848161064a57fe5b06141561067f5760008260a0811061065e57fe5b015481848151811061066c57fe5b6020026020010181815250506001830192505b5b60018311156107665760018201915060008090505b6002848161069f57fe5b04811015610705576106e08260028302815181106106b957fe5b60200260200101518360016002850201815181106106d357fe5b60200260200101516109f8565b8282815181106106ec57fe5b6020026020010181815250508080600101915050610695565b506002838161071057fe5b04925060016002848161071f57fe5b0614801561072e575060018314155b156107615760008260a0811061074057fe5b015481848151811061074e57fe5b6020026020010181815250506001830192505b610680565b8060008151811061077357fe5b602002602001015193505050505b92915050565b600060a060000154905061079c848484610488565b8060a060000154146107e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107da90611119565b60405180910390fd5b50505050565b60008160a081106107f657fe5b016000915090505481565b8160a0600001819055508060a0600101819055505050565b600060a060000154905090565b8160a0600201600061083786610988565b8152602001908152602001600020819055508060a0600201600061085a866109b7565b815260200190815260200160002081905550505050565b600083905060008090505b82518110156108ff5760008084838151811061089457fe5b6020026020010151905060006108aa8785610977565b905060008160ff1614156108d4576108c285836109f8565b92506108cf838684610826565b6108ec565b6108de82866109f8565b92506108eb838387610826565b5b829450505050808060010191505061087c565b508060a06000018190555050505050565b606061091b82610392565b9050610928838383610488565b505050565b60008060a0600201600061094085610988565b81526020019081526020016000205460a0600201600061095f866109b7565b81526020019081526020016000205491509150915091565b600060018284901c16905092915050565b60007f011111111111111111111111111111111111111111111111111111111111111160001b82169050919050565b60007f100000000000000000000000000000000000000000000000000000000000000060001b82179050919050565b60a08060000154908060010154905082565b60008282604051602001610a0d92919061102a565b60405160208183030381529060405280519060200120905092915050565b600082601f830112610a3c57600080fd5b8135610a4f610a4a826111d8565b6111ab565b91508181835260208401935060208101905083856020840282011115610a7457600080fd5b60005b83811015610aa45781610a8a8882610af8565b845260208401935060208301925050600181019050610a77565b5050505092915050565b60008083601f840112610ac057600080fd5b8235905067ffffffffffffffff811115610ad957600080fd5b602083019150836020820283011115610af157600080fd5b9250929050565b600081359050610b07816112c7565b92915050565b600082601f830112610b1e57600080fd5b8135610b31610b2c82611200565b6111ab565b91508082526020830160208301858383011115610b4d57600080fd5b610b588382846112ae565b50505092915050565b600081359050610b70816112de565b92915050565b60008060208385031215610b8957600080fd5b600083013567ffffffffffffffff811115610ba357600080fd5b610baf85828601610aae565b92509250509250929050565b600060208284031215610bcd57600080fd5b6000610bdb84828501610af8565b91505092915050565b600080600060608486031215610bf957600080fd5b6000610c0786828701610af8565b9350506020610c1886828701610af8565b9250506040610c2986828701610af8565b9150509250925092565b60008060008060808587031215610c4957600080fd5b6000610c5787828801610af8565b945050602085013567ffffffffffffffff811115610c7457600080fd5b610c8087828801610b0d565b9350506040610c9187828801610b61565b925050606085013567ffffffffffffffff811115610cae57600080fd5b610cba87828801610a2b565b91505092959194509250565b60008060408385031215610cd957600080fd5b6000610ce785828601610af8565b9250506020610cf885828601610b61565b9150509250929050565b600080600060608486031215610d1757600080fd5b6000610d2586828701610af8565b9350506020610d3686828701610b61565b925050604084013567ffffffffffffffff811115610d5357600080fd5b610d5f86828701610a2b565b9150509250925092565b60008060408385031215610d7c57600080fd5b600083013567ffffffffffffffff811115610d9657600080fd5b610da285828601610b0d565b9250506020610db385828601610b61565b9150509250929050565b600080600060608486031215610dd257600080fd5b600084013567ffffffffffffffff811115610dec57600080fd5b610df886828701610b0d565b9350506020610e0986828701610b61565b925050604084013567ffffffffffffffff811115610e2657600080fd5b610e3286828701610a2b565b9150509250925092565b600060208284031215610e4e57600080fd5b6000610e5c84828501610b61565b91505092915050565b60008060408385031215610e7857600080fd5b6000610e8685828601610b61565b9250506020610e9785828601610b61565b9150509250929050565b6000610ead8383610f26565b60208301905092915050565b6000610ec48261123c565b610ece8185611254565b9350610ed98361122c565b8060005b83811015610f0a578151610ef18882610ea1565b9750610efc83611247565b925050600181019050610edd565b5085935050505092915050565b610f2081611281565b82525050565b610f2f8161128d565b82525050565b610f3e8161128d565b82525050565b610f55610f508261128d565b6112bd565b82525050565b6000610f678385611265565b9350610f748385846112ae565b82840190509392505050565b6000610f8d605683611270565b91507f4661696c65642073616d6520726f6f7420766572696669636174696f6e20636860008301527f65636b2120546869732077617320616e20696e636c7573696f6e2070726f6f6660208301527f20666f72206120646966666572656e74207472656521000000000000000000006040830152606082019050919050565b61101581611297565b82525050565b611024816112a1565b82525050565b60006110368285610f44565b6020820191506110468284610f44565b6020820191508190509392505050565b6000611063828486610f5b565b91508190509392505050565b600060208201905081810360008301526110898184610eb9565b905092915050565b60006020820190506110a66000830184610f17565b92915050565b60006020820190506110c16000830184610f35565b92915050565b60006040820190506110dc6000830185610f35565b6110e96020830184610f35565b9392505050565b60006040820190506111056000830185610f35565b611112602083018461100c565b9392505050565b6000602082019050818103600083015261113281610f80565b9050919050565b600060208201905061114e600083018461101b565b92915050565b6000808335600160200384360303811261116d57600080fd5b80840192508235915067ffffffffffffffff82111561118b57600080fd5b6020830192506001820236038313156111a357600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff821117156111ce57600080fd5b8060405250919050565b600067ffffffffffffffff8211156111ef57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561121757600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b6000819050919050565b6112d08161128d565b81146112db57600080fd5b50565b6112e781611297565b81146112f257600080fd5b5056fea2646970667358221220101b678e592994820795c845941d04759a2941635b6c63474b8c27c2f3eb76f064736f6c63430006060033"

// DeployMerkleUtils deploys a new Ethereum contract, binding an instance of MerkleUtils to it.
func DeployMerkleUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleUtils{MerkleUtilsCaller: MerkleUtilsCaller{contract: contract}, MerkleUtilsTransactor: MerkleUtilsTransactor{contract: contract}, MerkleUtilsFilterer: MerkleUtilsFilterer{contract: contract}}, nil
}

// MerkleUtils is an auto generated Go binding around an Ethereum contract.
type MerkleUtils struct {
	MerkleUtilsCaller     // Read-only binding to the contract
	MerkleUtilsTransactor // Write-only binding to the contract
	MerkleUtilsFilterer   // Log filterer for contract events
}

// MerkleUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleUtilsSession struct {
	Contract     *MerkleUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleUtilsCallerSession struct {
	Contract *MerkleUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MerkleUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleUtilsTransactorSession struct {
	Contract     *MerkleUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MerkleUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleUtilsRaw struct {
	Contract *MerkleUtils // Generic contract binding to access the raw methods on
}

// MerkleUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleUtilsCallerRaw struct {
	Contract *MerkleUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleUtilsTransactorRaw struct {
	Contract *MerkleUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleUtils creates a new instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtils(address common.Address, backend bind.ContractBackend) (*MerkleUtils, error) {
	contract, err := bindMerkleUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleUtils{MerkleUtilsCaller: MerkleUtilsCaller{contract: contract}, MerkleUtilsTransactor: MerkleUtilsTransactor{contract: contract}, MerkleUtilsFilterer: MerkleUtilsFilterer{contract: contract}}, nil
}

// NewMerkleUtilsCaller creates a new read-only instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtilsCaller(address common.Address, caller bind.ContractCaller) (*MerkleUtilsCaller, error) {
	contract, err := bindMerkleUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilsCaller{contract: contract}, nil
}

// NewMerkleUtilsTransactor creates a new write-only instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleUtilsTransactor, error) {
	contract, err := bindMerkleUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilsTransactor{contract: contract}, nil
}

// NewMerkleUtilsFilterer creates a new log filterer instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleUtilsFilterer, error) {
	contract, err := bindMerkleUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilsFilterer{contract: contract}, nil
}

// bindMerkleUtils binds a generic wrapper to an already deployed contract.
func bindMerkleUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleUtils *MerkleUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleUtils.Contract.MerkleUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleUtils *MerkleUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleUtils.Contract.MerkleUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleUtils *MerkleUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleUtils.Contract.MerkleUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleUtils *MerkleUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleUtils *MerkleUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleUtils *MerkleUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleUtils.Contract.contract.Transact(opts, method, params...)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) DefaultHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "defaultHashes", arg0)
	return *ret0, err
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkleUtils.Contract.DefaultHashes(&_MerkleUtils.CallOpts, arg0)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkleUtils.Contract.DefaultHashes(&_MerkleUtils.CallOpts, arg0)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetChildren(opts *bind.CallOpts, _parent [32]byte) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MerkleUtils.contract.Call(opts, out, "getChildren", _parent)
	return *ret0, *ret1, err
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkleUtils.Contract.GetChildren(&_MerkleUtils.CallOpts, _parent)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkleUtils.Contract.GetChildren(&_MerkleUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetLeftSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "getLeftSiblingKey", _parent)
	return *ret0, err
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetLeftSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetLeftSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetMerkleRoot(opts *bind.CallOpts, _dataBlocks [][]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "getMerkleRoot", _dataBlocks)
	return *ret0, err
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetMerkleRoot(&_MerkleUtils.CallOpts, _dataBlocks)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetMerkleRoot(&_MerkleUtils.CallOpts, _dataBlocks)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkleUtils *MerkleUtilsCaller) GetNthBitFromRight(opts *bind.CallOpts, _intVal *big.Int, _index *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "getNthBitFromRight", _intVal, _index)
	return *ret0, err
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkleUtils *MerkleUtilsSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkleUtils.Contract.GetNthBitFromRight(&_MerkleUtils.CallOpts, _intVal, _index)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkleUtils *MerkleUtilsCallerSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkleUtils.Contract.GetNthBitFromRight(&_MerkleUtils.CallOpts, _intVal, _index)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetRightSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "getRightSiblingKey", _parent)
	return *ret0, err
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetRightSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetRightSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "getRoot")
	return *ret0, err
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetRoot() ([32]byte, error) {
	return _MerkleUtils.Contract.GetRoot(&_MerkleUtils.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetRoot() ([32]byte, error) {
	return _MerkleUtils.Contract.GetRoot(&_MerkleUtils.CallOpts)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkleUtils *MerkleUtilsCaller) GetSiblings(opts *bind.CallOpts, _path *big.Int) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "getSiblings", _path)
	return *ret0, err
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkleUtils *MerkleUtilsSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkleUtils.Contract.GetSiblings(&_MerkleUtils.CallOpts, _path)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkleUtils *MerkleUtilsCallerSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkleUtils.Contract.GetSiblings(&_MerkleUtils.CallOpts, _path)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkleUtils *MerkleUtilsCaller) Tree(opts *bind.CallOpts) (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	ret := new(struct {
		Root   [32]byte
		Height *big.Int
	})
	out := ret
	err := _MerkleUtils.contract.Call(opts, out, "tree")
	return *ret, err
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkleUtils *MerkleUtilsSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkleUtils.Contract.Tree(&_MerkleUtils.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkleUtils *MerkleUtilsCallerSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkleUtils.Contract.Tree(&_MerkleUtils.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkleUtils *MerkleUtilsCaller) Verify(opts *bind.CallOpts, _root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerkleUtils.contract.Call(opts, out, "verify", _root, _dataBlock, _path, _siblings)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkleUtils *MerkleUtilsSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkleUtils.Contract.Verify(&_MerkleUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkleUtils *MerkleUtilsCallerSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkleUtils.Contract.Verify(&_MerkleUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleUtils *MerkleUtilsTransactor) SetMerkleRootAndHeight(opts *bind.TransactOpts, _root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "setMerkleRootAndHeight", _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleUtils *MerkleUtilsSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.SetMerkleRootAndHeight(&_MerkleUtils.TransactOpts, _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.SetMerkleRootAndHeight(&_MerkleUtils.TransactOpts, _root, _height)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactor) Store(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "store", _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Store(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Store(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactor) StoreLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "storeLeaf", _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreLeaf(&_MerkleUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreLeaf(&_MerkleUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleUtils *MerkleUtilsTransactor) StoreNode(opts *bind.TransactOpts, _parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "storeNode", _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleUtils *MerkleUtilsSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreNode(&_MerkleUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreNode(&_MerkleUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactor) Update(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "update", _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Update(&_MerkleUtils.TransactOpts, _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Update(&_MerkleUtils.TransactOpts, _dataBlock, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactor) UpdateLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "updateLeaf", _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.UpdateLeaf(&_MerkleUtils.TransactOpts, _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.UpdateLeaf(&_MerkleUtils.TransactOpts, _leaf, _path)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactor) VerifyAndStore(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "verifyAndStore", _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.VerifyAndStore(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.VerifyAndStore(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}
