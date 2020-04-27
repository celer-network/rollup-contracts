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
var BlockCommitteeBin = "0x60806040523480156200001157600080fd5b506040516200239a3803806200239a833981810160405281019062000037919062000288565b6000620000496200010760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3508060019080519060200190620000ff9291906200010f565b505062000372565b600033905090565b8280548282559060005260206000209081019282156200018b579160200282015b828111156200018a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000130565b5b5090506200019a91906200019e565b5090565b620001e191905b80821115620001dd57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620001a5565b5090565b90565b600081519050620001f58162000358565b92915050565b600082601f8301126200020d57600080fd5b8151620002246200021e82620002fb565b620002cd565b915081818352602084019350602081019050838560208402820111156200024a57600080fd5b60005b838110156200027e5781620002638882620001e4565b8452602084019350602083019250506001810190506200024d565b5050505092915050565b6000602082840312156200029b57600080fd5b600082015167ffffffffffffffff811115620002b657600080fd5b620002c484828501620001fb565b91505092915050565b6000604051905081810181811067ffffffffffffffff82111715620002f157600080fd5b8060405250919050565b600067ffffffffffffffff8211156200031357600080fd5b602082029050602081019050919050565b6000620003318262000338565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620003638162000324565b81146200036f57600080fd5b50565b61201880620003826000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b1461011e5780639300c9261461013c578063d6119e7014610158578063d90d692714610174578063f2fde38b1461019257610093565b806335aa2e4414610098578063715018a6146100c85780637194d4b2146100d25780638be10194146100ee575b600080fd5b6100b260048036038101906100ad91906114ae565b6101ae565b6040516100bf9190611b4d565b60405180910390f35b6100d06101ea565b005b6100ec60048036038101906100e791906114d7565b61033f565b005b610108600480360381019061010391906114ae565b6105c3565b6040516101159190611bad565b60405180910390f35b61012661067c565b6040516101339190611b4d565b60405180910390f35b61015660048036038101906101519190611469565b6106a5565b005b610172600480360381019061016d9190611415565b610818565b005b61017c610ac1565b6040516101899190611b4d565b60405180910390f35b6101ac60048036038101906101a791906113ec565b610ae7565b005b600181815481106101bb57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6101f2610cab565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027790611ccf565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c690611cef565b60405180910390fd5b6000801515600560009054906101000a900460ff16151514610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d90611c8f565b60405180910390fd5b60608585905067ffffffffffffffff8111801561044257600080fd5b5060405190808252806020026020018201604052801561047657816020015b60608152602001906001900390816104615790505b50905060008090505b8686905081101561050f5786868281811061049657fe5b90506020028101906104a89190611d76565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508282815181106104f757fe5b6020026020010181905250808060010191505061047f565b50604051806040016040528088815260200182815250600660008201518160000155602082015181600101908051906020019061054d929190611013565b509050506001600560006101000a81548160ff0219169083151502179055506105ba3385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610818565b50505050505050565b600281815481106105d057fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106745780601f1061064957610100808354040283529160200191610674565b820191906000526020600020905b81548152906001019060200180831161065757829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106ad610cab565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461073b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073290611ccf565b60405180910390fd5b6000801515600560009054906101000a900460ff16151514610792576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078990611c8f565b60405180910390fd5b8282600191906107a3929190611073565b5060018054905067ffffffffffffffff811180156107c057600080fd5b506040519080825280602002602001820160405280156107f457816020015b60608152602001906001900390816107df5790505b506002908051906020019061080a929190611013565b506000600481905550505050565b6001801515600560009054906101000a900460ff1615151461086f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086690611c8f565b60405180910390fd5b60006001805490509050600080600090505b82811015610936576001818154811061089657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561092957846002828154811061090257fe5b90600052602060002001908051906020019061091f929190611113565b5060019150610936565b8080600101915050610881565b5080610977576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096e90611c4f565b60405180910390fd5b60006006600001546006600101604051602001610995929190611d46565b60405160208183030381529060405280519060200120905060006109b882610cb3565b90508673ffffffffffffffffffffffffffffffffffffffff166109db8288610ce3565b73ffffffffffffffffffffffffffffffffffffffff1614610a31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2890611bef565b60405180910390fd5b600860008154809291906001019190505550600060048510610a5d576002850260036008540211610a63565b84600854145b90508015610ab7577fd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc125360066002604051610a9e929190611d0f565b60405180910390a1610aae610ed4565b610ab6610f7d565b5b5050505050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610aef610cab565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7490611ccf565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610bed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be490611c2f565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b600081604051602001610cc69190611b27565b604051602081830303815290604052805190602001209050919050565b60006041825114610d29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2090611c0f565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610dac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da390611c6f565b60405180910390fd5b601b8160ff1614158015610dc45750601c8160ff1614155b15610e04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfb90611caf565b60405180910390fd5b600060018783868660405160008152602001604052604051610e299493929190611b68565b6020604051602081039080840390855afa158015610e4b573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ec7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ebe90611bcf565b60405180910390fd5b8094505050505092915050565b6000600560006101000a81548160ff02191690831515021790555060028054905067ffffffffffffffff81118015610f0b57600080fd5b50604051908082528060200260200182016040528015610f3f57816020015b6060815260200190600190039081610f2a5790505b5060029080519060200190610f55929190611013565b506006600080820160009055600182016000610f719190611193565b50506000600881905550565b60018054905060016004540181610f9057fe5b06600481905550600160045481548110610fa657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b828054828255906000526020600020908101928215611062579160200282015b82811115611061578251829080519060200190611051929190611113565b5091602001919060010190611033565b5b50905061106f91906111b4565b5090565b828054828255906000526020600020908101928215611102579160200282015b8281111561110157823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611093565b5b50905061110f91906111e0565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061115457805160ff1916838001178555611182565b82800160010185558215611182579182015b82811115611181578251825591602001919060010190611166565b5b50905061118f9190611223565b5090565b50805460008255906000526020600020908101906111b191906111b4565b50565b6111dd91905b808211156111d957600081816111d09190611248565b506001016111ba565b5090565b90565b61122091905b8082111561121c57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016111e6565b5090565b90565b61124591905b80821115611241576000816000905550600101611229565b5090565b90565b50805460018160011615610100020316600290046000825580601f1061126e575061128d565b601f01602090049060005260206000209081019061128c9190611223565b5b50565b60008135905061129f81611fb4565b92915050565b60008083601f8401126112b757600080fd5b8235905067ffffffffffffffff8111156112d057600080fd5b6020830191508360208202830111156112e857600080fd5b9250929050565b60008083601f84011261130157600080fd5b8235905067ffffffffffffffff81111561131a57600080fd5b60208301915083602082028301111561133257600080fd5b9250929050565b60008083601f84011261134b57600080fd5b8235905067ffffffffffffffff81111561136457600080fd5b60208301915083600182028301111561137c57600080fd5b9250929050565b600082601f83011261139457600080fd5b81356113a76113a282611dfa565b611dcd565b915080825260208301602083018583830111156113c357600080fd5b6113ce838284611f30565b50505092915050565b6000813590506113e681611fcb565b92915050565b6000602082840312156113fe57600080fd5b600061140c84828501611290565b91505092915050565b6000806040838503121561142857600080fd5b600061143685828601611290565b925050602083013567ffffffffffffffff81111561145357600080fd5b61145f85828601611383565b9150509250929050565b6000806020838503121561147c57600080fd5b600083013567ffffffffffffffff81111561149657600080fd5b6114a2858286016112a5565b92509250509250929050565b6000602082840312156114c057600080fd5b60006114ce848285016113d7565b91505092915050565b6000806000806000606086880312156114ef57600080fd5b60006114fd888289016113d7565b955050602086013567ffffffffffffffff81111561151a57600080fd5b611526888289016112ef565b9450945050604086013567ffffffffffffffff81111561154557600080fd5b61155188828901611339565b92509250509295509295909350565b600061156c83836116ca565b905092915050565b61157d81611edd565b82525050565b600061158e82611e50565b6115988185611e73565b9350836020820285016115aa85611e26565b8060005b858110156115e5578484038952816115c68582611560565b94506115d183611e66565b925060208a019950506001810190506115ae565b50829750879550505050505092915050565b600061160282611e50565b61160c8185611e84565b93508360208202850161161e85611e26565b8060005b858110156116595784840389528161163a8582611560565b945061164583611e66565b925060208a01995050600181019050611622565b50829750879550505050505092915050565b61167481611eef565b82525050565b61168b61168682611eef565b611f8c565b82525050565b600061169c82611e5b565b6116a68185611ea6565b93506116b6818560208601611f3f565b6116bf81611f96565b840191505092915050565b6000815460018116600081146116e7576001811461170d57611751565b607f60028304166116f88187611e95565b955060ff198316865260208601935050611751565b6002820461171b8187611e95565b955061172685611e3b565b60005b8281101561174857815481890152600182019150602081019050611729565b80880195505050505b505092915050565b6000611766601883611eb7565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b60006117a6601583611eb7565b91507f5369676e617475726520697320696e76616c69642100000000000000000000006000830152602082019050919050565b60006117e6601f83611eb7565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b6000611826601c83611ec8565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611866602683611eb7565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006118cc601a83611eb7565b91507f5369676e6572206d75737420626520612076616c696461746f720000000000006000830152602082019050919050565b600061190c602283611eb7565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611972601783611eb7565b91507f496e76616c69642070726f706f73616c207374617475730000000000000000006000830152602082019050919050565b60006119b2602283611eb7565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611a18602083611eb7565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000611a58602183611eb7565b91507f4f6e6c7920636f6d6d6974746572206d617920706572666f726d20616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000604083016000808401549050611ac881611f72565b611ad56000870182611afa565b50600184018583036020870152611aec8382611583565b925050819250505092915050565b611b0381611f19565b82525050565b611b1281611f19565b82525050565b611b2181611f23565b82525050565b6000611b3282611819565b9150611b3e828461167a565b60208201915081905092915050565b6000602082019050611b626000830184611574565b92915050565b6000608082019050611b7d600083018761166b565b611b8a6020830186611b18565b611b97604083018561166b565b611ba4606083018461166b565b95945050505050565b60006020820190508181036000830152611bc78184611691565b905092915050565b60006020820190508181036000830152611be881611759565b9050919050565b60006020820190508181036000830152611c0881611799565b9050919050565b60006020820190508181036000830152611c28816117d9565b9050919050565b60006020820190508181036000830152611c4881611859565b9050919050565b60006020820190508181036000830152611c68816118bf565b9050919050565b60006020820190508181036000830152611c88816118ff565b9050919050565b60006020820190508181036000830152611ca881611965565b9050919050565b60006020820190508181036000830152611cc8816119a5565b9050919050565b60006020820190508181036000830152611ce881611a0b565b9050919050565b60006020820190508181036000830152611d0881611a4b565b9050919050565b60006040820190508181036000830152611d298185611ab1565b90508181036020830152611d3d81846115f7565b90509392505050565b6000604082019050611d5b6000830185611b09565b8181036020830152611d6d81846115f7565b90509392505050565b60008083356001602003843603038112611d8f57600080fd5b80840192508235915067ffffffffffffffff821115611dad57600080fd5b602083019250600182023603831315611dc557600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff82111715611df057600080fd5b8060405250919050565b600067ffffffffffffffff821115611e1157600080fd5b601f19601f8301169050602081019050919050565b60008190508160005260206000209050919050565b60008190508160005260206000209050919050565b600081549050919050565b600081519050919050565b6000600182019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000819050919050565b6000611ee882611ef9565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611f5d578082015181840152602081019050611f42565b83811115611f6c576000848401525b50505050565b6000611f85611f8083611fa7565b611ed3565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160001c9050919050565b611fbd81611edd565b8114611fc857600080fd5b50565b611fd481611f19565b8114611fdf57600080fd5b5056fea264697066735822122095b7617b03ff933154570793d794fec06810a7bffc9401ae2e9d034bf06b990364736f6c63430006060033"

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
