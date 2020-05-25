// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainchain

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorRegistryABI is the input ABI used to generate the binding from.
const ValidatorRegistryABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCommitter\",\"type\":\"address\"}],\"name\":\"CommitterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_transitions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentCommitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pickNextCommitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupChainAddress\",\"type\":\"address\"}],\"name\":\"setRollupChainAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorRegistryBin is the compiled bytecode used for deploying new contracts.
var ValidatorRegistryBin = "0x60806040523480156200001157600080fd5b5060405162001fdc38038062001fdc833981810160405281019062000037919062000288565b6000620000496200010760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3508060019080519060200190620000ff9291906200010f565b505062000372565b600033905090565b8280548282559060005260206000209081019282156200018b579160200282015b828111156200018a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000130565b5b5090506200019a91906200019e565b5090565b620001e191905b80821115620001dd57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620001a5565b5090565b90565b600081519050620001f58162000358565b92915050565b600082601f8301126200020d57600080fd5b8151620002246200021e82620002fb565b620002cd565b915081818352602084019350602081019050838560208402820111156200024a57600080fd5b60005b838110156200027e5781620002638882620001e4565b8452602084019350602083019250506001810190506200024d565b5050505092915050565b6000602082840312156200029b57600080fd5b600082015167ffffffffffffffff811115620002b657600080fd5b620002c484828501620001fb565b91505092915050565b6000604051905081810181811067ffffffffffffffff82111715620002f157600080fd5b8060405250919050565b600067ffffffffffffffff8211156200031357600080fd5b602082029050602081019050919050565b6000620003318262000338565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620003638162000324565b81146200036f57600080fd5b50565b611c5a80620003826000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80639300c926116100665780639300c926146100fa578063d354166514610116578063d90d692714610146578063f2fde38b14610164578063f690495e1461018057610093565b806335aa2e44146100985780634ef2b1be146100c8578063715018a6146100d25780638da5cb5b146100dc575b600080fd5b6100b260048036038101906100ad91906112d2565b61019c565b6040516100bf919061185d565b60405180910390f35b6100d06101d8565b005b6100da610406565b005b6100e461055b565b6040516100f1919061185d565b60405180910390f35b610114600480360381019061010f919061128d565b610584565b005b610130600480360381019061012b91906112fb565b610668565b60405161013d9190611878565b60405180910390f35b61014e6108e4565b60405161015b919061185d565b60405180910390f35b61017e60048036038101906101799190611264565b61090a565b005b61019a60048036038101906101959190611264565b610ace565b005b600181815481106101a957fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025f90611a18565b60405180910390fd5b6001805490506001600354018161027b57fe5b0660038190555060016003548154811061029157fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fb48ade409baa2182ab477a3be094650b21d9a9e71fc34e14129c0c94b6ea1970600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161034d919061185d565b60405180910390a1600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166313df8728600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016103d2919061185d565b600060405180830381600087803b1580156103ec57600080fd5b505af1158015610400573d6000803e3d6000fd5b50505050565b61040e610c37565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461049c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610493906119d8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61058c610c37565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461061a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610611906119d8565b60405180910390fd5b610664828280806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050610c3f565b5050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f190611a18565b60405180910390fd5b60006001805490509050600080600090505b8281101561087557600089898960405160200161072b93929190611a38565b604051602081830303815290604052805190602001209050600061074e82610eb8565b90506001838154811061075d57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610808828a8a878181106107ae57fe5b90506020028101906107c09190611a6a565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610ee8565b73ffffffffffffffffffffffffffffffffffffffff161461085e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610855906118f8565b60405180910390fd5b83806001019450505050808060010191505061070c565b5060006004831061088e57600283026003830211610892565b8282145b9050806108d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108cb90611958565b60405180910390fd5b6001935050505095945050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610912610c37565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146109a0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610997906119d8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0790611978565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610ad6610c37565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5b906119d8565b60405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610c346001805480602002602001604051908101604052809291908181526020018280548015610c2a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610be0575b5050505050610c3f565b50565b600033905090565b6000815111610c83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7a90611938565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610d15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0c906119f8565b60405180910390fd5b8060019080519060200190610d2b9291906110d9565b5060006003819055506001600081548110610d4257fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fb48ade409baa2182ab477a3be094650b21d9a9e71fc34e14129c0c94b6ea1970600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610dfe919061185d565b60405180910390a1600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166313df8728600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401610e83919061185d565b600060405180830381600087803b158015610e9d57600080fd5b505af1158015610eb1573d6000803e3d6000fd5b5050505050565b600081604051602001610ecb9190611837565b604051602081830303815290604052805190602001209050919050565b60006041825114610f2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2590611918565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610fb1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa890611998565b60405180910390fd5b601b8160ff1614158015610fc95750601c8160ff1614155b15611009576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611000906119b8565b60405180910390fd5b60006001878386866040516000815260200160405260405161102e9493929190611893565b6020604051602081039080840390855afa158015611050573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156110cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c3906118d8565b60405180910390fd5b8094505050505092915050565b828054828255906000526020600020908101928215611152579160200282015b828111156111515782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906110f9565b5b50905061115f9190611163565b5090565b6111a391905b8082111561119f57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611169565b5090565b90565b6000813590506111b581611bf6565b92915050565b60008083601f8401126111cd57600080fd5b8235905067ffffffffffffffff8111156111e657600080fd5b6020830191508360208202830111156111fe57600080fd5b9250929050565b60008083601f84011261121757600080fd5b8235905067ffffffffffffffff81111561123057600080fd5b60208301915083602082028301111561124857600080fd5b9250929050565b60008135905061125e81611c0d565b92915050565b60006020828403121561127657600080fd5b6000611284848285016111a6565b91505092915050565b600080602083850312156112a057600080fd5b600083013567ffffffffffffffff8111156112ba57600080fd5b6112c6858286016111bb565b92509250509250929050565b6000602082840312156112e457600080fd5b60006112f28482850161124f565b91505092915050565b60008060008060006060868803121561131357600080fd5b60006113218882890161124f565b955050602086013567ffffffffffffffff81111561133e57600080fd5b61134a88828901611205565b9450945050604086013567ffffffffffffffff81111561136957600080fd5b61137588828901611205565b92509250509295509295909350565b6000611391848484611454565b90509392505050565b6113a381611b6d565b82525050565b60006113b58385611ad8565b9350836020840285016113c784611ac1565b8060005b8781101561140d5784840389526113e28284611b16565b6113ed868284611384565b95506113f884611acb565b935060208b019a5050506001810190506113cb565b50829750879450505050509392505050565b61142881611b7f565b82525050565b61143781611b8b565b82525050565b61144e61144982611b8b565b611bdb565b82525050565b60006114608385611ae9565b935061146d838584611bcc565b61147683611be5565b840190509392505050565b600061148e601883611afa565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b60006114ce601583611afa565b91507f5369676e617475726520697320696e76616c69642100000000000000000000006000830152602082019050919050565b600061150e601f83611afa565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b600061154e601383611afa565b91507f456d7074792076616c696461746f7220736574000000000000000000000000006000830152602082019050919050565b600061158e601c83611b0b565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b60006115ce601583611afa565b91507f4e6f7420656e6f756768207369676e61747572657300000000000000000000006000830152602082019050919050565b600061160e602683611afa565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611674602283611afa565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006116da602283611afa565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611740602083611afa565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000611780601383611afa565b91507f526f6c6c7570436861696e206e6f7420736574000000000000000000000000006000830152602082019050919050565b60006117c0602383611afa565b91507f4f6e6c7920526f6c6c7570436861696e206d617920706572666f726d2061637460008301527f696f6e00000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b61182281611bb5565b82525050565b61183181611bbf565b82525050565b600061184282611581565b915061184e828461143d565b60208201915081905092915050565b6000602082019050611872600083018461139a565b92915050565b600060208201905061188d600083018461141f565b92915050565b60006080820190506118a8600083018761142e565b6118b56020830186611828565b6118c2604083018561142e565b6118cf606083018461142e565b95945050505050565b600060208201905081810360008301526118f181611481565b9050919050565b60006020820190508181036000830152611911816114c1565b9050919050565b6000602082019050818103600083015261193181611501565b9050919050565b6000602082019050818103600083015261195181611541565b9050919050565b60006020820190508181036000830152611971816115c1565b9050919050565b6000602082019050818103600083015261199181611601565b9050919050565b600060208201905081810360008301526119b181611667565b9050919050565b600060208201905081810360008301526119d1816116cd565b9050919050565b600060208201905081810360008301526119f181611733565b9050919050565b60006020820190508181036000830152611a1181611773565b9050919050565b60006020820190508181036000830152611a31816117b3565b9050919050565b6000604082019050611a4d6000830186611819565b8181036020830152611a608184866113a9565b9050949350505050565b60008083356001602003843603038112611a8357600080fd5b80840192508235915067ffffffffffffffff821115611aa157600080fd5b602083019250600182023603831315611ab957600080fd5b509250929050565b6000819050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60008083356001602003843603038112611b2f57600080fd5b83810192508235915060208301925067ffffffffffffffff821115611b5357600080fd5b600182023603841315611b6557600080fd5b509250929050565b6000611b7882611b95565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b6000819050919050565b6000601f19601f8301169050919050565b611bff81611b6d565b8114611c0a57600080fd5b50565b611c1681611bb5565b8114611c2157600080fd5b5056fea26469706673582212204d09bf2a0456f4deaae9c91ded8a42cf02c97f0e741f50600c526b786334d2fe64736f6c63430006060033"

// DeployValidatorRegistry deploys a new Ethereum contract, binding an instance of ValidatorRegistry to it.
func DeployValidatorRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _validators []common.Address) (common.Address, *types.Transaction, *ValidatorRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorRegistryBin), backend, _validators)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorRegistry{ValidatorRegistryCaller: ValidatorRegistryCaller{contract: contract}, ValidatorRegistryTransactor: ValidatorRegistryTransactor{contract: contract}, ValidatorRegistryFilterer: ValidatorRegistryFilterer{contract: contract}}, nil
}

// ValidatorRegistry is an auto generated Go binding around an Ethereum contract.
type ValidatorRegistry struct {
	ValidatorRegistryCaller     // Read-only binding to the contract
	ValidatorRegistryTransactor // Write-only binding to the contract
	ValidatorRegistryFilterer   // Log filterer for contract events
}

// ValidatorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorRegistrySession struct {
	Contract     *ValidatorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorRegistryCallerSession struct {
	Contract *ValidatorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorRegistryTransactorSession struct {
	Contract     *ValidatorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRegistryRaw struct {
	Contract *ValidatorRegistry // Generic contract binding to access the raw methods on
}

// ValidatorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorRegistryCallerRaw struct {
	Contract *ValidatorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorRegistryTransactorRaw struct {
	Contract *ValidatorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorRegistry creates a new instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistry(address common.Address, backend bind.ContractBackend) (*ValidatorRegistry, error) {
	contract, err := bindValidatorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistry{ValidatorRegistryCaller: ValidatorRegistryCaller{contract: contract}, ValidatorRegistryTransactor: ValidatorRegistryTransactor{contract: contract}, ValidatorRegistryFilterer: ValidatorRegistryFilterer{contract: contract}}, nil
}

// NewValidatorRegistryCaller creates a new read-only instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryCaller(address common.Address, caller bind.ContractCaller) (*ValidatorRegistryCaller, error) {
	contract, err := bindValidatorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryCaller{contract: contract}, nil
}

// NewValidatorRegistryTransactor creates a new write-only instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorRegistryTransactor, error) {
	contract, err := bindValidatorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryTransactor{contract: contract}, nil
}

// NewValidatorRegistryFilterer creates a new log filterer instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorRegistryFilterer, error) {
	contract, err := bindValidatorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryFilterer{contract: contract}, nil
}

// bindValidatorRegistry binds a generic wrapper to an already deployed contract.
func bindValidatorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRegistry *ValidatorRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorRegistry.Contract.ValidatorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRegistry *ValidatorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.ValidatorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRegistry *ValidatorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.ValidatorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRegistry *ValidatorRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRegistry *ValidatorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRegistry *ValidatorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.contract.Transact(opts, method, params...)
}

// CheckSignatures is a free data retrieval call binding the contract method 0xd3541665.
//
// Solidity: function checkSignatures(uint256 _blockNumber, bytes[] _transitions, bytes[] _signatures) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCaller) CheckSignatures(opts *bind.CallOpts, _blockNumber *big.Int, _transitions [][]byte, _signatures [][]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorRegistry.contract.Call(opts, out, "checkSignatures", _blockNumber, _transitions, _signatures)
	return *ret0, err
}

// CheckSignatures is a free data retrieval call binding the contract method 0xd3541665.
//
// Solidity: function checkSignatures(uint256 _blockNumber, bytes[] _transitions, bytes[] _signatures) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistrySession) CheckSignatures(_blockNumber *big.Int, _transitions [][]byte, _signatures [][]byte) (bool, error) {
	return _ValidatorRegistry.Contract.CheckSignatures(&_ValidatorRegistry.CallOpts, _blockNumber, _transitions, _signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0xd3541665.
//
// Solidity: function checkSignatures(uint256 _blockNumber, bytes[] _transitions, bytes[] _signatures) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) CheckSignatures(_blockNumber *big.Int, _transitions [][]byte, _signatures [][]byte) (bool, error) {
	return _ValidatorRegistry.Contract.CheckSignatures(&_ValidatorRegistry.CallOpts, _blockNumber, _transitions, _signatures)
}

// CurrentCommitter is a free data retrieval call binding the contract method 0xd90d6927.
//
// Solidity: function currentCommitter() view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCaller) CurrentCommitter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorRegistry.contract.Call(opts, out, "currentCommitter")
	return *ret0, err
}

// CurrentCommitter is a free data retrieval call binding the contract method 0xd90d6927.
//
// Solidity: function currentCommitter() view returns(address)
func (_ValidatorRegistry *ValidatorRegistrySession) CurrentCommitter() (common.Address, error) {
	return _ValidatorRegistry.Contract.CurrentCommitter(&_ValidatorRegistry.CallOpts)
}

// CurrentCommitter is a free data retrieval call binding the contract method 0xd90d6927.
//
// Solidity: function currentCommitter() view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) CurrentCommitter() (common.Address, error) {
	return _ValidatorRegistry.Contract.CurrentCommitter(&_ValidatorRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorRegistry *ValidatorRegistrySession) Owner() (common.Address, error) {
	return _ValidatorRegistry.Contract.Owner(&_ValidatorRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) Owner() (common.Address, error) {
	return _ValidatorRegistry.Contract.Owner(&_ValidatorRegistry.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorRegistry.contract.Call(opts, out, "validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorRegistry *ValidatorRegistrySession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorRegistry.Contract.Validators(&_ValidatorRegistry.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorRegistry.Contract.Validators(&_ValidatorRegistry.CallOpts, arg0)
}

// PickNextCommitter is a paid mutator transaction binding the contract method 0x4ef2b1be.
//
// Solidity: function pickNextCommitter() returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) PickNextCommitter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "pickNextCommitter")
}

// PickNextCommitter is a paid mutator transaction binding the contract method 0x4ef2b1be.
//
// Solidity: function pickNextCommitter() returns()
func (_ValidatorRegistry *ValidatorRegistrySession) PickNextCommitter() (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.PickNextCommitter(&_ValidatorRegistry.TransactOpts)
}

// PickNextCommitter is a paid mutator transaction binding the contract method 0x4ef2b1be.
//
// Solidity: function pickNextCommitter() returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) PickNextCommitter() (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.PickNextCommitter(&_ValidatorRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorRegistry *ValidatorRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.RenounceOwnership(&_ValidatorRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.RenounceOwnership(&_ValidatorRegistry.TransactOpts)
}

// SetRollupChainAddress is a paid mutator transaction binding the contract method 0xf690495e.
//
// Solidity: function setRollupChainAddress(address _rollupChainAddress) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) SetRollupChainAddress(opts *bind.TransactOpts, _rollupChainAddress common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "setRollupChainAddress", _rollupChainAddress)
}

// SetRollupChainAddress is a paid mutator transaction binding the contract method 0xf690495e.
//
// Solidity: function setRollupChainAddress(address _rollupChainAddress) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) SetRollupChainAddress(_rollupChainAddress common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetRollupChainAddress(&_ValidatorRegistry.TransactOpts, _rollupChainAddress)
}

// SetRollupChainAddress is a paid mutator transaction binding the contract method 0xf690495e.
//
// Solidity: function setRollupChainAddress(address _rollupChainAddress) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) SetRollupChainAddress(_rollupChainAddress common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetRollupChainAddress(&_ValidatorRegistry.TransactOpts, _rollupChainAddress)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) SetValidators(opts *bind.TransactOpts, _validators []common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "setValidators", _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetValidators(&_ValidatorRegistry.TransactOpts, _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetValidators(&_ValidatorRegistry.TransactOpts, _validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.TransferOwnership(&_ValidatorRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.TransferOwnership(&_ValidatorRegistry.TransactOpts, newOwner)
}

// ValidatorRegistryCommitterChangedIterator is returned from FilterCommitterChanged and is used to iterate over the raw logs and unpacked data for CommitterChanged events raised by the ValidatorRegistry contract.
type ValidatorRegistryCommitterChangedIterator struct {
	Event *ValidatorRegistryCommitterChanged // Event containing the contract specifics and raw log

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
func (it *ValidatorRegistryCommitterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRegistryCommitterChanged)
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
		it.Event = new(ValidatorRegistryCommitterChanged)
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
func (it *ValidatorRegistryCommitterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRegistryCommitterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRegistryCommitterChanged represents a CommitterChanged event raised by the ValidatorRegistry contract.
type ValidatorRegistryCommitterChanged struct {
	NewCommitter common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCommitterChanged is a free log retrieval operation binding the contract event 0xb48ade409baa2182ab477a3be094650b21d9a9e71fc34e14129c0c94b6ea1970.
//
// Solidity: event CommitterChanged(address newCommitter)
func (_ValidatorRegistry *ValidatorRegistryFilterer) FilterCommitterChanged(opts *bind.FilterOpts) (*ValidatorRegistryCommitterChangedIterator, error) {

	logs, sub, err := _ValidatorRegistry.contract.FilterLogs(opts, "CommitterChanged")
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryCommitterChangedIterator{contract: _ValidatorRegistry.contract, event: "CommitterChanged", logs: logs, sub: sub}, nil
}

// WatchCommitterChanged is a free log subscription operation binding the contract event 0xb48ade409baa2182ab477a3be094650b21d9a9e71fc34e14129c0c94b6ea1970.
//
// Solidity: event CommitterChanged(address newCommitter)
func (_ValidatorRegistry *ValidatorRegistryFilterer) WatchCommitterChanged(opts *bind.WatchOpts, sink chan<- *ValidatorRegistryCommitterChanged) (event.Subscription, error) {

	logs, sub, err := _ValidatorRegistry.contract.WatchLogs(opts, "CommitterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRegistryCommitterChanged)
				if err := _ValidatorRegistry.contract.UnpackLog(event, "CommitterChanged", log); err != nil {
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

// ParseCommitterChanged is a log parse operation binding the contract event 0xb48ade409baa2182ab477a3be094650b21d9a9e71fc34e14129c0c94b6ea1970.
//
// Solidity: event CommitterChanged(address newCommitter)
func (_ValidatorRegistry *ValidatorRegistryFilterer) ParseCommitterChanged(log types.Log) (*ValidatorRegistryCommitterChanged, error) {
	event := new(ValidatorRegistryCommitterChanged)
	if err := _ValidatorRegistry.contract.UnpackLog(event, "CommitterChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorRegistry contract.
type ValidatorRegistryOwnershipTransferredIterator struct {
	Event *ValidatorRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRegistryOwnershipTransferred)
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
		it.Event = new(ValidatorRegistryOwnershipTransferred)
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
func (it *ValidatorRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ValidatorRegistry contract.
type ValidatorRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorRegistry *ValidatorRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryOwnershipTransferredIterator{contract: _ValidatorRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorRegistry *ValidatorRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRegistryOwnershipTransferred)
				if err := _ValidatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ValidatorRegistry *ValidatorRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorRegistryOwnershipTransferred, error) {
	event := new(ValidatorRegistryOwnershipTransferred)
	if err := _ValidatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
