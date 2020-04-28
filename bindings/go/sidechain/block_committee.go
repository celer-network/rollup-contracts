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
const BlockCommitteeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transitions\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structBlockCommittee.BlockProposal\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"BlockConsensusReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transitions\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structBlockCommittee.BlockProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProposer\",\"type\":\"address\"}],\"name\":\"ProposerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_transitions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"signBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlockCommitteeBin is the compiled bytecode used for deploying new contracts.
var BlockCommitteeBin = "0x60806040523480156200001157600080fd5b50604051620029b3380380620029b38339818101604052810190620000379190620005bf565b600062000049620000ff60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350620000f8816200010760201b60201c565b506200074c565b600033905090565b60008151116200014e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001459062000674565b60405180910390fd5b806001908051906020019062000166929190620002b3565b5060018054905067ffffffffffffffff811180156200018457600080fd5b50604051908082528060200260200182016040528015620001ba57816020015b6060815260200190600190039081620001a45790505b5060029080519060200190620001d292919062000342565b5060006004819055506001600081548110620001ea57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620002a8919062000657565b60405180910390a150565b8280548282559060005260206000209081019282156200032f579160200282015b828111156200032e5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190620002d4565b5b5090506200033e9190620003a9565b5090565b82805482825590600052602060002090810192821562000396579160200282015b828111156200039557825182908051906020019062000384929190620003ef565b509160200191906001019062000363565b5b509050620003a5919062000476565b5090565b620003ec91905b80821115620003e857600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620003b0565b5090565b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200043257805160ff191683800117855562000463565b8280016001018555821562000463579182015b828111156200046257825182559160200191906001019062000445565b5b509050620004729190620004a7565b5090565b620004a491905b80821115620004a05760008181620004969190620004cf565b506001016200047d565b5090565b90565b620004cc91905b80821115620004c8576000816000905550600101620004ae565b5090565b90565b50805460018160011615610100020316600290046000825580601f10620004f7575062000518565b601f016020900490600052602060002090810190620005179190620004a7565b5b50565b6000815190506200052c8162000732565b92915050565b600082601f8301126200054457600080fd5b81516200055b6200055582620006c4565b62000696565b915081818352602084019350602081019050838560208402820111156200058157600080fd5b60005b83811015620005b557816200059a88826200051b565b84526020840193506020830192505060018101905062000584565b5050505092915050565b600060208284031215620005d257600080fd5b600082015167ffffffffffffffff811115620005ed57600080fd5b620005fb8482850162000532565b91505092915050565b6200060f81620006fe565b82525050565b600062000624601383620006ed565b91507f456d7074792076616c696461746f7220736574000000000000000000000000006000830152602082019050919050565b60006020820190506200066e600083018462000604565b92915050565b600060208201905081810360008301526200068f8162000615565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620006ba57600080fd5b8060405250919050565b600067ffffffffffffffff821115620006dc57600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b60006200070b8262000712565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200073d81620006fe565b81146200074957600080fd5b50565b612257806200075c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b1461011e5780639300c9261461013c578063d6119e7014610158578063f2fde38b14610174578063fe82d0821461019057610093565b806335aa2e4414610098578063715018a6146100c85780637194d4b2146100d25780638be10194146100ee575b600080fd5b6100b260048036038101906100ad9190611691565b6101ae565b6040516100bf9190611d4a565b60405180910390f35b6100d06101ea565b005b6100ec60048036038101906100e791906116ba565b61033f565b005b61010860048036038101906101039190611691565b6105fb565b6040516101159190611daa565b60405180910390f35b6101266106b4565b6040516101339190611d4a565b60405180910390f35b6101566004803603810190610151919061164c565b6106dd565b005b610172600480360381019061016d91906115f8565b610819565b005b61018e600480360381019061018991906115cf565b610ac2565b005b610198610c86565b6040516101a59190611d4a565b60405180910390f35b600181815481106101bb57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6101f2610cac565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027790611f0c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c690611ecc565b60405180910390fd5b6000801515600560009054906101000a900460ff16151514610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d90611eac565b60405180910390fd5b60608585905067ffffffffffffffff8111801561044257600080fd5b5060405190808252806020026020018201604052801561047657816020015b60608152602001906001900390816104615790505b50905060008090505b8686905081101561050f5786868281811061049657fe5b90506020028101906104a89190611fb5565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508282815181106104f757fe5b6020026020010181905250808060010191505061047f565b50604051806040016040528088815260200182815250600660008201518160000155602082015181600101908051906020019061054d92919061120c565b509050506001600560006101000a81548160ff0219169083151502179055507f8e859432a5d0af78f97935da86e1c192b0e88a3beeb6d5ba1cacf5ccff5b704b600660405161059c9190611f2c565b60405180910390a16105f23385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610819565b50505050505050565b6002818154811061060857fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106ac5780601f10610681576101008083540402835291602001916106ac565b820191906000526020600020905b81548152906001019060200180831161068f57829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106e5610cac565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610773576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076a90611f0c565b60405180910390fd5b6000801515600560009054906101000a900460ff161515146107ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c190611eac565b60405180910390fd5b610814838380806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050610cb4565b505050565b6001801515600560009054906101000a900460ff16151514610870576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086790611eac565b60405180910390fd5b60006001805490509050600080600090505b82811015610937576001818154811061089757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561092a57846002828154811061090357fe5b90600052602060002001908051906020019061092092919061126c565b5060019150610937565b8080600101915050610882565b5080610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096f90611e6c565b60405180910390fd5b60006006600001546006600101604051602001610996929190611f85565b60405160208183030381529060405280519060200120905060006109b982610e53565b90508673ffffffffffffffffffffffffffffffffffffffff166109dc8288610e83565b73ffffffffffffffffffffffffffffffffffffffff1614610a32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2990611dec565b60405180910390fd5b600860008154809291906001019190505550600060048510610a5e576002850260036008540211610a64565b84600854145b90508015610ab8577fd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc125360066002604051610a9f929190611f4e565b60405180910390a1610aaf611074565b610ab761111d565b5b5050505050505050565b610aca610cac565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4f90611f0c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610bc8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bbf90611e4c565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b6000815111610cf8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cef90611e2c565b60405180910390fd5b8060019080519060200190610d0e9291906112ec565b5060018054905067ffffffffffffffff81118015610d2b57600080fd5b50604051908082528060200260200182016040528015610d5f57816020015b6060815260200190600190039081610d4a5790505b5060029080519060200190610d7592919061120c565b5060006004819055506001600081548110610d8c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610e489190611d4a565b60405180910390a150565b600081604051602001610e669190611d24565b604051602081830303815290604052805190602001209050919050565b60006041825114610ec9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec090611e0c565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610f4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4390611e8c565b60405180910390fd5b601b8160ff1614158015610f645750601c8160ff1614155b15610fa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9b90611eec565b60405180910390fd5b600060018783868660405160008152602001604052604051610fc99493929190611d65565b6020604051602081039080840390855afa158015610feb573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611067576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105e90611dcc565b60405180910390fd5b8094505050505092915050565b6000600560006101000a81548160ff02191690831515021790555060028054905067ffffffffffffffff811180156110ab57600080fd5b506040519080825280602002602001820160405280156110df57816020015b60608152602001906001900390816110ca5790505b50600290805190602001906110f592919061120c565b5060066000808201600090556001820160006111119190611376565b50506000600881905550565b6001805490506001600454018161113057fe5b0660048190555060016004548154811061114657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516112029190611d4a565b60405180910390a1565b82805482825590600052602060002090810192821561125b579160200282015b8281111561125a57825182908051906020019061124a92919061126c565b509160200191906001019061122c565b5b5090506112689190611397565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112ad57805160ff19168380011785556112db565b828001600101855582156112db579182015b828111156112da5782518255916020019190600101906112bf565b5b5090506112e891906113c3565b5090565b828054828255906000526020600020908101928215611365579160200282015b828111156113645782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061130c565b5b50905061137291906113e8565b5090565b50805460008255906000526020600020908101906113949190611397565b50565b6113c091905b808211156113bc57600081816113b3919061142b565b5060010161139d565b5090565b90565b6113e591905b808211156113e15760008160009055506001016113c9565b5090565b90565b61142891905b8082111561142457600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016113ee565b5090565b90565b50805460018160011615610100020316600290046000825580601f106114515750611470565b601f01602090049060005260206000209081019061146f91906113c3565b5b50565b600081359050611482816121f3565b92915050565b60008083601f84011261149a57600080fd5b8235905067ffffffffffffffff8111156114b357600080fd5b6020830191508360208202830111156114cb57600080fd5b9250929050565b60008083601f8401126114e457600080fd5b8235905067ffffffffffffffff8111156114fd57600080fd5b60208301915083602082028301111561151557600080fd5b9250929050565b60008083601f84011261152e57600080fd5b8235905067ffffffffffffffff81111561154757600080fd5b60208301915083600182028301111561155f57600080fd5b9250929050565b600082601f83011261157757600080fd5b813561158a61158582612039565b61200c565b915080825260208301602083018583830111156115a657600080fd5b6115b183828461216f565b50505092915050565b6000813590506115c98161220a565b92915050565b6000602082840312156115e157600080fd5b60006115ef84828501611473565b91505092915050565b6000806040838503121561160b57600080fd5b600061161985828601611473565b925050602083013567ffffffffffffffff81111561163657600080fd5b61164285828601611566565b9150509250929050565b6000806020838503121561165f57600080fd5b600083013567ffffffffffffffff81111561167957600080fd5b61168585828601611488565b92509250509250929050565b6000602082840312156116a357600080fd5b60006116b1848285016115ba565b91505092915050565b6000806000806000606086880312156116d257600080fd5b60006116e0888289016115ba565b955050602086013567ffffffffffffffff8111156116fd57600080fd5b611709888289016114d2565b9450945050604086013567ffffffffffffffff81111561172857600080fd5b6117348882890161151c565b92509250509295509295909350565b600061174f83836118ad565b905092915050565b6117608161211c565b82525050565b60006117718261208f565b61177b81856120b2565b93508360208202850161178d85612065565b8060005b858110156117c8578484038952816117a98582611743565b94506117b4836120a5565b925060208a01995050600181019050611791565b50829750879550505050505092915050565b60006117e58261208f565b6117ef81856120c3565b93508360208202850161180185612065565b8060005b8581101561183c5784840389528161181d8582611743565b9450611828836120a5565b925060208a01995050600181019050611805565b50829750879550505050505092915050565b6118578161212e565b82525050565b61186e6118698261212e565b6121cb565b82525050565b600061187f8261209a565b61188981856120e5565b935061189981856020860161217e565b6118a2816121d5565b840191505092915050565b6000815460018116600081146118ca57600181146118f057611934565b607f60028304166118db81876120d4565b955060ff198316865260208601935050611934565b600282046118fe81876120d4565b95506119098561207a565b60005b8281101561192b5781548189015260018201915060208101905061190c565b80880195505050505b505092915050565b60006119496018836120f6565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b60006119896015836120f6565b91507f5369676e617475726520697320696e76616c69642100000000000000000000006000830152602082019050919050565b60006119c9601f836120f6565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b6000611a096013836120f6565b91507f456d7074792076616c696461746f7220736574000000000000000000000000006000830152602082019050919050565b6000611a49601c83612107565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611a896026836120f6565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611aef601a836120f6565b91507f5369676e6572206d75737420626520612076616c696461746f720000000000006000830152602082019050919050565b6000611b2f6022836120f6565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611b956017836120f6565b91507f496e76616c69642070726f706f73616c207374617475730000000000000000006000830152602082019050919050565b6000611bd56020836120f6565b91507f4f6e6c792070726f706f736572206d617920706572666f726d20616374696f6e6000830152602082019050919050565b6000611c156022836120f6565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611c7b6020836120f6565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000604083016000808401549050611cc5816121b1565b611cd26000870182611cf7565b50600184018583036020870152611ce98382611766565b925050819250505092915050565b611d0081612158565b82525050565b611d0f81612158565b82525050565b611d1e81612162565b82525050565b6000611d2f82611a3c565b9150611d3b828461185d565b60208201915081905092915050565b6000602082019050611d5f6000830184611757565b92915050565b6000608082019050611d7a600083018761184e565b611d876020830186611d15565b611d94604083018561184e565b611da1606083018461184e565b95945050505050565b60006020820190508181036000830152611dc48184611874565b905092915050565b60006020820190508181036000830152611de58161193c565b9050919050565b60006020820190508181036000830152611e058161197c565b9050919050565b60006020820190508181036000830152611e25816119bc565b9050919050565b60006020820190508181036000830152611e45816119fc565b9050919050565b60006020820190508181036000830152611e6581611a7c565b9050919050565b60006020820190508181036000830152611e8581611ae2565b9050919050565b60006020820190508181036000830152611ea581611b22565b9050919050565b60006020820190508181036000830152611ec581611b88565b9050919050565b60006020820190508181036000830152611ee581611bc8565b9050919050565b60006020820190508181036000830152611f0581611c08565b9050919050565b60006020820190508181036000830152611f2581611c6e565b9050919050565b60006020820190508181036000830152611f468184611cae565b905092915050565b60006040820190508181036000830152611f688185611cae565b90508181036020830152611f7c81846117da565b90509392505050565b6000604082019050611f9a6000830185611d06565b8181036020830152611fac81846117da565b90509392505050565b60008083356001602003843603038112611fce57600080fd5b80840192508235915067ffffffffffffffff821115611fec57600080fd5b60208301925060018202360383131561200457600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff8211171561202f57600080fd5b8060405250919050565b600067ffffffffffffffff82111561205057600080fd5b601f19601f8301169050602081019050919050565b60008190508160005260206000209050919050565b60008190508160005260206000209050919050565b600081549050919050565b600081519050919050565b6000600182019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000819050919050565b600061212782612138565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561219c578082015181840152602081019050612181565b838111156121ab576000848401525b50505050565b60006121c46121bf836121e6565b612112565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160001c9050919050565b6121fc8161211c565b811461220757600080fd5b50565b61221381612158565b811461221e57600080fd5b5056fea2646970667358221220013fff91c10098f62977a53a7f779ea2f8e9a84113f5fbf76de062fe10e65ee464736f6c63430006060033"

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

// CurrentProposer is a free data retrieval call binding the contract method 0xfe82d082.
//
// Solidity: function currentProposer() view returns(address)
func (_BlockCommittee *BlockCommitteeCaller) CurrentProposer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockCommittee.contract.Call(opts, out, "currentProposer")
	return *ret0, err
}

// CurrentProposer is a free data retrieval call binding the contract method 0xfe82d082.
//
// Solidity: function currentProposer() view returns(address)
func (_BlockCommittee *BlockCommitteeSession) CurrentProposer() (common.Address, error) {
	return _BlockCommittee.Contract.CurrentProposer(&_BlockCommittee.CallOpts)
}

// CurrentProposer is a free data retrieval call binding the contract method 0xfe82d082.
//
// Solidity: function currentProposer() view returns(address)
func (_BlockCommittee *BlockCommitteeCallerSession) CurrentProposer() (common.Address, error) {
	return _BlockCommittee.Contract.CurrentProposer(&_BlockCommittee.CallOpts)
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

// BlockCommitteeProposerChangedIterator is returned from FilterProposerChanged and is used to iterate over the raw logs and unpacked data for ProposerChanged events raised by the BlockCommittee contract.
type BlockCommitteeProposerChangedIterator struct {
	Event *BlockCommitteeProposerChanged // Event containing the contract specifics and raw log

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
func (it *BlockCommitteeProposerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockCommitteeProposerChanged)
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
		it.Event = new(BlockCommitteeProposerChanged)
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
func (it *BlockCommitteeProposerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockCommitteeProposerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockCommitteeProposerChanged represents a ProposerChanged event raised by the BlockCommittee contract.
type BlockCommitteeProposerChanged struct {
	NewProposer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposerChanged is a free log retrieval operation binding the contract event 0x2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693.
//
// Solidity: event ProposerChanged(address newProposer)
func (_BlockCommittee *BlockCommitteeFilterer) FilterProposerChanged(opts *bind.FilterOpts) (*BlockCommitteeProposerChangedIterator, error) {

	logs, sub, err := _BlockCommittee.contract.FilterLogs(opts, "ProposerChanged")
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeProposerChangedIterator{contract: _BlockCommittee.contract, event: "ProposerChanged", logs: logs, sub: sub}, nil
}

// WatchProposerChanged is a free log subscription operation binding the contract event 0x2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693.
//
// Solidity: event ProposerChanged(address newProposer)
func (_BlockCommittee *BlockCommitteeFilterer) WatchProposerChanged(opts *bind.WatchOpts, sink chan<- *BlockCommitteeProposerChanged) (event.Subscription, error) {

	logs, sub, err := _BlockCommittee.contract.WatchLogs(opts, "ProposerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockCommitteeProposerChanged)
				if err := _BlockCommittee.contract.UnpackLog(event, "ProposerChanged", log); err != nil {
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

// ParseProposerChanged is a log parse operation binding the contract event 0x2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693.
//
// Solidity: event ProposerChanged(address newProposer)
func (_BlockCommittee *BlockCommitteeFilterer) ParseProposerChanged(log types.Log) (*BlockCommitteeProposerChanged, error) {
	event := new(BlockCommitteeProposerChanged)
	if err := _BlockCommittee.contract.UnpackLog(event, "ProposerChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
