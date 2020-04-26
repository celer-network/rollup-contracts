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

// BlockCommitteeBlockProposal is an auto generated low-level Go binding around an user-defined struct.
type BlockCommitteeBlockProposal struct {
	BlockNumber *big.Int
	Transitions [][]byte
}

// BlockCommitteeABI is the input ABI used to generate the binding from.
const BlockCommitteeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transitions\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structBlockCommittee.BlockProposal\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"BlockConsensusReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transitions\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structBlockCommittee.BlockProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentCommitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_transitions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"signBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlockCommitteeBin is the compiled bytecode used for deploying new contracts.
var BlockCommitteeBin = "0x60806040523480156200001157600080fd5b506040516200239b3803806200239b833981810160405281019062000037919062000288565b6000620000496200010760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3508060019080519060200190620000ff9291906200010f565b505062000372565b600033905090565b8280548282559060005260206000209081019282156200018b579160200282015b828111156200018a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000130565b5b5090506200019a91906200019e565b5090565b620001e191905b80821115620001dd57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620001a5565b5090565b90565b600081519050620001f58162000358565b92915050565b600082601f8301126200020d57600080fd5b8151620002246200021e82620002fb565b620002cd565b915081818352602084019350602081019050838560208402820111156200024a57600080fd5b60005b838110156200027e5781620002638882620001e4565b8452602084019350602083019250506001810190506200024d565b5050505092915050565b6000602082840312156200029b57600080fd5b600082015167ffffffffffffffff811115620002b657600080fd5b620002c484828501620001fb565b91505092915050565b6000604051905081810181811067ffffffffffffffff82111715620002f157600080fd5b8060405250919050565b600067ffffffffffffffff8211156200031357600080fd5b602082029050602081019050919050565b6000620003318262000338565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620003638162000324565b81146200036f57600080fd5b50565b61201980620003826000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b1461011e5780639300c9261461013c578063d6119e7014610158578063d90d692714610174578063f2fde38b1461019257610093565b806335aa2e4414610098578063715018a6146100c85780637194d4b2146100d25780638be10194146100ee575b600080fd5b6100b260048036038101906100ad91906114af565b6101ae565b6040516100bf9190611b4e565b60405180910390f35b6100d06101ea565b005b6100ec60048036038101906100e791906114d8565b61033f565b005b610108600480360381019061010391906114af565b6105c3565b6040516101159190611bae565b60405180910390f35b61012661067c565b6040516101339190611b4e565b60405180910390f35b6101566004803603810190610151919061146a565b6106a5565b005b610172600480360381019061016d9190611416565b610818565b005b61017c610ac2565b6040516101899190611b4e565b60405180910390f35b6101ac60048036038101906101a791906113ed565b610ae8565b005b600181815481106101bb57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6101f2610cac565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027790611cd0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c690611cf0565b60405180910390fd5b6000801515600560009054906101000a900460ff16151514610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d90611c90565b60405180910390fd5b60608585905067ffffffffffffffff8111801561044257600080fd5b5060405190808252806020026020018201604052801561047657816020015b60608152602001906001900390816104615790505b50905060008090505b8686905081101561050f5786868281811061049657fe5b90506020028101906104a89190611d77565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508282815181106104f757fe5b6020026020010181905250808060010191505061047f565b50604051806040016040528088815260200182815250600660008201518160000155602082015181600101908051906020019061054d929190611014565b509050506001600560006101000a81548160ff0219169083151502179055506105ba3385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610818565b50505050505050565b600281815481106105d057fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106745780601f1061064957610100808354040283529160200191610674565b820191906000526020600020905b81548152906001019060200180831161065757829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106ad610cac565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461073b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073290611cd0565b60405180910390fd5b6000801515600560009054906101000a900460ff16151514610792576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078990611c90565b60405180910390fd5b8282600191906107a3929190611074565b5060018054905067ffffffffffffffff811180156107c057600080fd5b506040519080825280602002602001820160405280156107f457816020015b60608152602001906001900390816107df5790505b506002908051906020019061080a929190611014565b506000600481905550505050565b6001801515600560009054906101000a900460ff1615151461086f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086690611c90565b60405180910390fd5b60006001805490509050600080600090505b82811015610932576001818154811061089657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561092557846002828154811061090257fe5b90600052602060002001908051906020019061091f929190611114565b50610932565b8080600101915050610881565b5080610973576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096a90611c50565b60405180910390fd5b60006006600001546006600101604051602001610991929190611d47565b60405160208183030381529060405280519060200120905060006109b482610cb4565b90508673ffffffffffffffffffffffffffffffffffffffff166109d78288610ce4565b73ffffffffffffffffffffffffffffffffffffffff1614610a2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2490611bf0565b60405180910390fd5b6008600081548092919060010191905055506000600460018054905010610a5e576002850260036008540211610a64565b84600854145b90508015610ab8577fd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc125360066002604051610a9f929190611d10565b60405180910390a1610aaf610ed5565b610ab7610f7e565b5b5050505050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610af0610cac565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7590611cd0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610bee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be590611c30565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b600081604051602001610cc79190611b28565b604051602081830303815290604052805190602001209050919050565b60006041825114610d2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2190611c10565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610dad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da490611c70565b60405180910390fd5b601b8160ff1614158015610dc55750601c8160ff1614155b15610e05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfc90611cb0565b60405180910390fd5b600060018783868660405160008152602001604052604051610e2a9493929190611b69565b6020604051602081039080840390855afa158015610e4c573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ec8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ebf90611bd0565b60405180910390fd5b8094505050505092915050565b6000600560006101000a81548160ff02191690831515021790555060028054905067ffffffffffffffff81118015610f0c57600080fd5b50604051908082528060200260200182016040528015610f4057816020015b6060815260200190600190039081610f2b5790505b5060029080519060200190610f56929190611014565b506006600080820160009055600182016000610f729190611194565b50506000600881905550565b60018054905060016004540181610f9157fe5b06600481905550600160045481548110610fa757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b828054828255906000526020600020908101928215611063579160200282015b82811115611062578251829080519060200190611052929190611114565b5091602001919060010190611034565b5b50905061107091906111b5565b5090565b828054828255906000526020600020908101928215611103579160200282015b8281111561110257823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611094565b5b50905061111091906111e1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061115557805160ff1916838001178555611183565b82800160010185558215611183579182015b82811115611182578251825591602001919060010190611167565b5b5090506111909190611224565b5090565b50805460008255906000526020600020908101906111b291906111b5565b50565b6111de91905b808211156111da57600081816111d19190611249565b506001016111bb565b5090565b90565b61122191905b8082111561121d57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016111e7565b5090565b90565b61124691905b8082111561124257600081600090555060010161122a565b5090565b90565b50805460018160011615610100020316600290046000825580601f1061126f575061128e565b601f01602090049060005260206000209081019061128d9190611224565b5b50565b6000813590506112a081611fb5565b92915050565b60008083601f8401126112b857600080fd5b8235905067ffffffffffffffff8111156112d157600080fd5b6020830191508360208202830111156112e957600080fd5b9250929050565b60008083601f84011261130257600080fd5b8235905067ffffffffffffffff81111561131b57600080fd5b60208301915083602082028301111561133357600080fd5b9250929050565b60008083601f84011261134c57600080fd5b8235905067ffffffffffffffff81111561136557600080fd5b60208301915083600182028301111561137d57600080fd5b9250929050565b600082601f83011261139557600080fd5b81356113a86113a382611dfb565b611dce565b915080825260208301602083018583830111156113c457600080fd5b6113cf838284611f31565b50505092915050565b6000813590506113e781611fcc565b92915050565b6000602082840312156113ff57600080fd5b600061140d84828501611291565b91505092915050565b6000806040838503121561142957600080fd5b600061143785828601611291565b925050602083013567ffffffffffffffff81111561145457600080fd5b61146085828601611384565b9150509250929050565b6000806020838503121561147d57600080fd5b600083013567ffffffffffffffff81111561149757600080fd5b6114a3858286016112a6565b92509250509250929050565b6000602082840312156114c157600080fd5b60006114cf848285016113d8565b91505092915050565b6000806000806000606086880312156114f057600080fd5b60006114fe888289016113d8565b955050602086013567ffffffffffffffff81111561151b57600080fd5b611527888289016112f0565b9450945050604086013567ffffffffffffffff81111561154657600080fd5b6115528882890161133a565b92509250509295509295909350565b600061156d83836116cb565b905092915050565b61157e81611ede565b82525050565b600061158f82611e51565b6115998185611e74565b9350836020820285016115ab85611e27565b8060005b858110156115e6578484038952816115c78582611561565b94506115d283611e67565b925060208a019950506001810190506115af565b50829750879550505050505092915050565b600061160382611e51565b61160d8185611e85565b93508360208202850161161f85611e27565b8060005b8581101561165a5784840389528161163b8582611561565b945061164683611e67565b925060208a01995050600181019050611623565b50829750879550505050505092915050565b61167581611ef0565b82525050565b61168c61168782611ef0565b611f8d565b82525050565b600061169d82611e5c565b6116a78185611ea7565b93506116b7818560208601611f40565b6116c081611f97565b840191505092915050565b6000815460018116600081146116e8576001811461170e57611752565b607f60028304166116f98187611e96565b955060ff198316865260208601935050611752565b6002820461171c8187611e96565b955061172785611e3c565b60005b828110156117495781548189015260018201915060208101905061172a565b80880195505050505b505092915050565b6000611767601883611eb8565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b60006117a7601583611eb8565b91507f5369676e617475726520697320696e76616c69642100000000000000000000006000830152602082019050919050565b60006117e7601f83611eb8565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b6000611827601c83611ec9565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611867602683611eb8565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006118cd601a83611eb8565b91507f5369676e6572206d75737420626520612076616c696461746f720000000000006000830152602082019050919050565b600061190d602283611eb8565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611973601783611eb8565b91507f496e76616c69642070726f706f73616c207374617475730000000000000000006000830152602082019050919050565b60006119b3602283611eb8565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611a19602083611eb8565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000611a59602183611eb8565b91507f4f6e6c7920636f6d6d6974746572206d617920706572666f726d20616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000604083016000808401549050611ac981611f73565b611ad66000870182611afb565b50600184018583036020870152611aed8382611584565b925050819250505092915050565b611b0481611f1a565b82525050565b611b1381611f1a565b82525050565b611b2281611f24565b82525050565b6000611b338261181a565b9150611b3f828461167b565b60208201915081905092915050565b6000602082019050611b636000830184611575565b92915050565b6000608082019050611b7e600083018761166c565b611b8b6020830186611b19565b611b98604083018561166c565b611ba5606083018461166c565b95945050505050565b60006020820190508181036000830152611bc88184611692565b905092915050565b60006020820190508181036000830152611be98161175a565b9050919050565b60006020820190508181036000830152611c098161179a565b9050919050565b60006020820190508181036000830152611c29816117da565b9050919050565b60006020820190508181036000830152611c498161185a565b9050919050565b60006020820190508181036000830152611c69816118c0565b9050919050565b60006020820190508181036000830152611c8981611900565b9050919050565b60006020820190508181036000830152611ca981611966565b9050919050565b60006020820190508181036000830152611cc9816119a6565b9050919050565b60006020820190508181036000830152611ce981611a0c565b9050919050565b60006020820190508181036000830152611d0981611a4c565b9050919050565b60006040820190508181036000830152611d2a8185611ab2565b90508181036020830152611d3e81846115f8565b90509392505050565b6000604082019050611d5c6000830185611b0a565b8181036020830152611d6e81846115f8565b90509392505050565b60008083356001602003843603038112611d9057600080fd5b80840192508235915067ffffffffffffffff821115611dae57600080fd5b602083019250600182023603831315611dc657600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff82111715611df157600080fd5b8060405250919050565b600067ffffffffffffffff821115611e1257600080fd5b601f19601f8301169050602081019050919050565b60008190508160005260206000209050919050565b60008190508160005260206000209050919050565b600081549050919050565b600081519050919050565b6000600182019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000819050919050565b6000611ee982611efa565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611f5e578082015181840152602081019050611f43565b83811115611f6d576000848401525b50505050565b6000611f86611f8183611fa8565b611ed4565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160001c9050919050565b611fbe81611ede565b8114611fc957600080fd5b50565b611fd581611f1a565b8114611fe057600080fd5b5056fea26469706673582212209420cb97902ed647a4e5482239a25f4003f0b76d3936b86c67b2626a2f9aa1ef64736f6c63430006060033"

// DeployBlockCommittee deploys a new Ethereum contract, binding an instance of BlockCommittee to it.
func DeployBlockCommittee(auth *bind.TransactOpts, backend bind.ContractBackend, _validators []common.Address) (common.Address, *types.Transaction, *BlockCommittee, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockCommitteeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlockCommitteeBin), backend, _validators)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockCommittee{BlockCommitteeCaller: BlockCommitteeCaller{contract: contract}, BlockCommitteeTransactor: BlockCommitteeTransactor{contract: contract}, BlockCommitteeFilterer: BlockCommitteeFilterer{contract: contract}}, nil
}

// BlockCommittee is an auto generated Go binding around an Ethereum contract.
type BlockCommittee struct {
	BlockCommitteeCaller     // Read-only binding to the contract
	BlockCommitteeTransactor // Write-only binding to the contract
	BlockCommitteeFilterer   // Log filterer for contract events
}

// BlockCommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockCommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockCommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockCommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockCommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockCommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockCommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockCommitteeSession struct {
	Contract     *BlockCommittee   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockCommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockCommitteeCallerSession struct {
	Contract *BlockCommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BlockCommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockCommitteeTransactorSession struct {
	Contract     *BlockCommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BlockCommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockCommitteeRaw struct {
	Contract *BlockCommittee // Generic contract binding to access the raw methods on
}

// BlockCommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockCommitteeCallerRaw struct {
	Contract *BlockCommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// BlockCommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockCommitteeTransactorRaw struct {
	Contract *BlockCommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockCommittee creates a new instance of BlockCommittee, bound to a specific deployed contract.
func NewBlockCommittee(address common.Address, backend bind.ContractBackend) (*BlockCommittee, error) {
	contract, err := bindBlockCommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockCommittee{BlockCommitteeCaller: BlockCommitteeCaller{contract: contract}, BlockCommitteeTransactor: BlockCommitteeTransactor{contract: contract}, BlockCommitteeFilterer: BlockCommitteeFilterer{contract: contract}}, nil
}

// NewBlockCommitteeCaller creates a new read-only instance of BlockCommittee, bound to a specific deployed contract.
func NewBlockCommitteeCaller(address common.Address, caller bind.ContractCaller) (*BlockCommitteeCaller, error) {
	contract, err := bindBlockCommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeCaller{contract: contract}, nil
}

// NewBlockCommitteeTransactor creates a new write-only instance of BlockCommittee, bound to a specific deployed contract.
func NewBlockCommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockCommitteeTransactor, error) {
	contract, err := bindBlockCommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeTransactor{contract: contract}, nil
}

// NewBlockCommitteeFilterer creates a new log filterer instance of BlockCommittee, bound to a specific deployed contract.
func NewBlockCommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockCommitteeFilterer, error) {
	contract, err := bindBlockCommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeFilterer{contract: contract}, nil
}

// bindBlockCommittee binds a generic wrapper to an already deployed contract.
func bindBlockCommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockCommitteeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockCommittee *BlockCommitteeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockCommittee.Contract.BlockCommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockCommittee *BlockCommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockCommittee.Contract.BlockCommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockCommittee *BlockCommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockCommittee.Contract.BlockCommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockCommittee *BlockCommitteeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockCommittee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockCommittee *BlockCommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockCommittee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockCommittee *BlockCommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockCommittee.Contract.contract.Transact(opts, method, params...)
}

// CurrentCommitter is a free data retrieval call binding the contract method 0xd90d6927.
//
// Solidity: function currentCommitter() view returns(address)
func (_BlockCommittee *BlockCommitteeCaller) CurrentCommitter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockCommittee.contract.Call(opts, out, "currentCommitter")
	return *ret0, err
}

// CurrentCommitter is a free data retrieval call binding the contract method 0xd90d6927.
//
// Solidity: function currentCommitter() view returns(address)
func (_BlockCommittee *BlockCommitteeSession) CurrentCommitter() (common.Address, error) {
	return _BlockCommittee.Contract.CurrentCommitter(&_BlockCommittee.CallOpts)
}

// CurrentCommitter is a free data retrieval call binding the contract method 0xd90d6927.
//
// Solidity: function currentCommitter() view returns(address)
func (_BlockCommittee *BlockCommitteeCallerSession) CurrentCommitter() (common.Address, error) {
	return _BlockCommittee.Contract.CurrentCommitter(&_BlockCommittee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockCommittee *BlockCommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockCommittee.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockCommittee *BlockCommitteeSession) Owner() (common.Address, error) {
	return _BlockCommittee.Contract.Owner(&_BlockCommittee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockCommittee *BlockCommitteeCallerSession) Owner() (common.Address, error) {
	return _BlockCommittee.Contract.Owner(&_BlockCommittee.CallOpts)
}

// Signatures is a free data retrieval call binding the contract method 0x8be10194.
//
// Solidity: function signatures(uint256 ) view returns(bytes)
func (_BlockCommittee *BlockCommitteeCaller) Signatures(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _BlockCommittee.contract.Call(opts, out, "signatures", arg0)
	return *ret0, err
}

// Signatures is a free data retrieval call binding the contract method 0x8be10194.
//
// Solidity: function signatures(uint256 ) view returns(bytes)
func (_BlockCommittee *BlockCommitteeSession) Signatures(arg0 *big.Int) ([]byte, error) {
	return _BlockCommittee.Contract.Signatures(&_BlockCommittee.CallOpts, arg0)
}

// Signatures is a free data retrieval call binding the contract method 0x8be10194.
//
// Solidity: function signatures(uint256 ) view returns(bytes)
func (_BlockCommittee *BlockCommitteeCallerSession) Signatures(arg0 *big.Int) ([]byte, error) {
	return _BlockCommittee.Contract.Signatures(&_BlockCommittee.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_BlockCommittee *BlockCommitteeCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockCommittee.contract.Call(opts, out, "validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_BlockCommittee *BlockCommitteeSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _BlockCommittee.Contract.Validators(&_BlockCommittee.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_BlockCommittee *BlockCommitteeCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _BlockCommittee.Contract.Validators(&_BlockCommittee.CallOpts, arg0)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x7194d4b2.
//
// Solidity: function proposeBlock(uint256 _blockNumber, bytes[] _transitions, bytes _signature) returns()
func (_BlockCommittee *BlockCommitteeTransactor) ProposeBlock(opts *bind.TransactOpts, _blockNumber *big.Int, _transitions [][]byte, _signature []byte) (*types.Transaction, error) {
	return _BlockCommittee.contract.Transact(opts, "proposeBlock", _blockNumber, _transitions, _signature)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x7194d4b2.
//
// Solidity: function proposeBlock(uint256 _blockNumber, bytes[] _transitions, bytes _signature) returns()
func (_BlockCommittee *BlockCommitteeSession) ProposeBlock(_blockNumber *big.Int, _transitions [][]byte, _signature []byte) (*types.Transaction, error) {
	return _BlockCommittee.Contract.ProposeBlock(&_BlockCommittee.TransactOpts, _blockNumber, _transitions, _signature)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x7194d4b2.
//
// Solidity: function proposeBlock(uint256 _blockNumber, bytes[] _transitions, bytes _signature) returns()
func (_BlockCommittee *BlockCommitteeTransactorSession) ProposeBlock(_blockNumber *big.Int, _transitions [][]byte, _signature []byte) (*types.Transaction, error) {
	return _BlockCommittee.Contract.ProposeBlock(&_BlockCommittee.TransactOpts, _blockNumber, _transitions, _signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockCommittee *BlockCommitteeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockCommittee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockCommittee *BlockCommitteeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockCommittee.Contract.RenounceOwnership(&_BlockCommittee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockCommittee *BlockCommitteeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockCommittee.Contract.RenounceOwnership(&_BlockCommittee.TransactOpts)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_BlockCommittee *BlockCommitteeTransactor) SetValidators(opts *bind.TransactOpts, _validators []common.Address) (*types.Transaction, error) {
	return _BlockCommittee.contract.Transact(opts, "setValidators", _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_BlockCommittee *BlockCommitteeSession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _BlockCommittee.Contract.SetValidators(&_BlockCommittee.TransactOpts, _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_BlockCommittee *BlockCommitteeTransactorSession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _BlockCommittee.Contract.SetValidators(&_BlockCommittee.TransactOpts, _validators)
}

// SignBlock is a paid mutator transaction binding the contract method 0xd6119e70.
//
// Solidity: function signBlock(address _signer, bytes _signature) returns()
func (_BlockCommittee *BlockCommitteeTransactor) SignBlock(opts *bind.TransactOpts, _signer common.Address, _signature []byte) (*types.Transaction, error) {
	return _BlockCommittee.contract.Transact(opts, "signBlock", _signer, _signature)
}

// SignBlock is a paid mutator transaction binding the contract method 0xd6119e70.
//
// Solidity: function signBlock(address _signer, bytes _signature) returns()
func (_BlockCommittee *BlockCommitteeSession) SignBlock(_signer common.Address, _signature []byte) (*types.Transaction, error) {
	return _BlockCommittee.Contract.SignBlock(&_BlockCommittee.TransactOpts, _signer, _signature)
}

// SignBlock is a paid mutator transaction binding the contract method 0xd6119e70.
//
// Solidity: function signBlock(address _signer, bytes _signature) returns()
func (_BlockCommittee *BlockCommitteeTransactorSession) SignBlock(_signer common.Address, _signature []byte) (*types.Transaction, error) {
	return _BlockCommittee.Contract.SignBlock(&_BlockCommittee.TransactOpts, _signer, _signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockCommittee *BlockCommitteeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockCommittee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockCommittee *BlockCommitteeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockCommittee.Contract.TransferOwnership(&_BlockCommittee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockCommittee *BlockCommitteeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockCommittee.Contract.TransferOwnership(&_BlockCommittee.TransactOpts, newOwner)
}

// BlockCommitteeBlockConsensusReachedIterator is returned from FilterBlockConsensusReached and is used to iterate over the raw logs and unpacked data for BlockConsensusReached events raised by the BlockCommittee contract.
type BlockCommitteeBlockConsensusReachedIterator struct {
	Event *BlockCommitteeBlockConsensusReached // Event containing the contract specifics and raw log

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
func (it *BlockCommitteeBlockConsensusReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockCommitteeBlockConsensusReached)
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
		it.Event = new(BlockCommitteeBlockConsensusReached)
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
func (it *BlockCommitteeBlockConsensusReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockCommitteeBlockConsensusReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockCommitteeBlockConsensusReached represents a BlockConsensusReached event raised by the BlockCommittee contract.
type BlockCommitteeBlockConsensusReached struct {
	Proposal   BlockCommitteeBlockProposal
	Signatures [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockConsensusReached is a free log retrieval operation binding the contract event 0xd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc1253.
//
// Solidity: event BlockConsensusReached(BlockCommitteeBlockProposal proposal, bytes[] signatures)
func (_BlockCommittee *BlockCommitteeFilterer) FilterBlockConsensusReached(opts *bind.FilterOpts) (*BlockCommitteeBlockConsensusReachedIterator, error) {

	logs, sub, err := _BlockCommittee.contract.FilterLogs(opts, "BlockConsensusReached")
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeBlockConsensusReachedIterator{contract: _BlockCommittee.contract, event: "BlockConsensusReached", logs: logs, sub: sub}, nil
}

// WatchBlockConsensusReached is a free log subscription operation binding the contract event 0xd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc1253.
//
// Solidity: event BlockConsensusReached(BlockCommitteeBlockProposal proposal, bytes[] signatures)
func (_BlockCommittee *BlockCommitteeFilterer) WatchBlockConsensusReached(opts *bind.WatchOpts, sink chan<- *BlockCommitteeBlockConsensusReached) (event.Subscription, error) {

	logs, sub, err := _BlockCommittee.contract.WatchLogs(opts, "BlockConsensusReached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockCommitteeBlockConsensusReached)
				if err := _BlockCommittee.contract.UnpackLog(event, "BlockConsensusReached", log); err != nil {
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

// ParseBlockConsensusReached is a log parse operation binding the contract event 0xd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc1253.
//
// Solidity: event BlockConsensusReached(BlockCommitteeBlockProposal proposal, bytes[] signatures)
func (_BlockCommittee *BlockCommitteeFilterer) ParseBlockConsensusReached(log types.Log) (*BlockCommitteeBlockConsensusReached, error) {
	event := new(BlockCommitteeBlockConsensusReached)
	if err := _BlockCommittee.contract.UnpackLog(event, "BlockConsensusReached", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BlockCommitteeBlockProposedIterator is returned from FilterBlockProposed and is used to iterate over the raw logs and unpacked data for BlockProposed events raised by the BlockCommittee contract.
type BlockCommitteeBlockProposedIterator struct {
	Event *BlockCommitteeBlockProposed // Event containing the contract specifics and raw log

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
func (it *BlockCommitteeBlockProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockCommitteeBlockProposed)
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
		it.Event = new(BlockCommitteeBlockProposed)
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
func (it *BlockCommitteeBlockProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockCommitteeBlockProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockCommitteeBlockProposed represents a BlockProposed event raised by the BlockCommittee contract.
type BlockCommitteeBlockProposed struct {
	Proposal BlockCommitteeBlockProposal
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x8e859432a5d0af78f97935da86e1c192b0e88a3beeb6d5ba1cacf5ccff5b704b.
//
// Solidity: event BlockProposed(BlockCommitteeBlockProposal proposal)
func (_BlockCommittee *BlockCommitteeFilterer) FilterBlockProposed(opts *bind.FilterOpts) (*BlockCommitteeBlockProposedIterator, error) {

	logs, sub, err := _BlockCommittee.contract.FilterLogs(opts, "BlockProposed")
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeBlockProposedIterator{contract: _BlockCommittee.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0x8e859432a5d0af78f97935da86e1c192b0e88a3beeb6d5ba1cacf5ccff5b704b.
//
// Solidity: event BlockProposed(BlockCommitteeBlockProposal proposal)
func (_BlockCommittee *BlockCommitteeFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *BlockCommitteeBlockProposed) (event.Subscription, error) {

	logs, sub, err := _BlockCommittee.contract.WatchLogs(opts, "BlockProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockCommitteeBlockProposed)
				if err := _BlockCommittee.contract.UnpackLog(event, "BlockProposed", log); err != nil {
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

// ParseBlockProposed is a log parse operation binding the contract event 0x8e859432a5d0af78f97935da86e1c192b0e88a3beeb6d5ba1cacf5ccff5b704b.
//
// Solidity: event BlockProposed(BlockCommitteeBlockProposal proposal)
func (_BlockCommittee *BlockCommitteeFilterer) ParseBlockProposed(log types.Log) (*BlockCommitteeBlockProposed, error) {
	event := new(BlockCommitteeBlockProposed)
	if err := _BlockCommittee.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BlockCommitteeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockCommittee contract.
type BlockCommitteeOwnershipTransferredIterator struct {
	Event *BlockCommitteeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockCommitteeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockCommitteeOwnershipTransferred)
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
		it.Event = new(BlockCommitteeOwnershipTransferred)
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
func (it *BlockCommitteeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockCommitteeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockCommitteeOwnershipTransferred represents a OwnershipTransferred event raised by the BlockCommittee contract.
type BlockCommitteeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockCommittee *BlockCommitteeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockCommitteeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockCommittee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeOwnershipTransferredIterator{contract: _BlockCommittee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockCommittee *BlockCommitteeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockCommitteeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockCommittee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockCommitteeOwnershipTransferred)
				if err := _BlockCommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BlockCommittee *BlockCommitteeFilterer) ParseOwnershipTransferred(log types.Log) (*BlockCommitteeOwnershipTransferred, error) {
	event := new(BlockCommitteeOwnershipTransferred)
	if err := _BlockCommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
