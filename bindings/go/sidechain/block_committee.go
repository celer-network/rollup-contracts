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
const BlockCommitteeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transitions\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structBlockCommittee.BlockProposal\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"BlockConsensusReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"transitions\",\"type\":\"bytes[]\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProposer\",\"type\":\"address\"}],\"name\":\"ProposerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_transitions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"signBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlockCommitteeBin is the compiled bytecode used for deploying new contracts.
var BlockCommitteeBin = "0x60806040523480156200001157600080fd5b5060405162002ac438038062002ac4833981810160405281019062000037919062000636565b600062000049620000ff60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350620000f8816200010760201b60201c565b50620007c3565b600033905090565b60008151116200014e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200014590620006eb565b60405180910390fd5b80600190805190602001906200016692919062000307565b50600060048190555060016000815481106200017e57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516200023c9190620006ce565b60405180910390a1620002546200025760201b60201c565b50565b6000600560006101000a81548160ff02191690831515021790555060018054905067ffffffffffffffff811180156200028f57600080fd5b50604051908082528060200260200182016040528015620002c557816020015b6060815260200190600190039081620002af5790505b5060029080519060200190620002dd92919062000396565b506006600080820160009055600182016000620002fb9190620003fd565b50506000600881905550565b82805482825590600052602060002090810192821562000383579160200282015b82811115620003825782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000328565b5b50905062000392919062000420565b5090565b828054828255906000526020600020908101928215620003ea579160200282015b82811115620003e9578251829080519060200190620003d892919062000466565b5091602001919060010190620003b7565b5b509050620003f99190620004ed565b5090565b50805460008255906000526020600020908101906200041d9190620004ed565b50565b6200046391905b808211156200045f57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010162000427565b5090565b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004a957805160ff1916838001178555620004da565b82800160010185558215620004da579182015b82811115620004d9578251825591602001919060010190620004bc565b5b509050620004e991906200051e565b5090565b6200051b91905b808211156200051757600081816200050d919062000546565b50600101620004f4565b5090565b90565b6200054391905b808211156200053f57600081600090555060010162000525565b5090565b90565b50805460018160011615610100020316600290046000825580601f106200056e57506200058f565b601f0160209004906000526020600020908101906200058e91906200051e565b5b50565b600081519050620005a381620007a9565b92915050565b600082601f830112620005bb57600080fd5b8151620005d2620005cc826200073b565b6200070d565b91508181835260208401935060208101905083856020840282011115620005f857600080fd5b60005b838110156200062c578162000611888262000592565b845260208401935060208301925050600181019050620005fb565b5050505092915050565b6000602082840312156200064957600080fd5b600082015167ffffffffffffffff8111156200066457600080fd5b6200067284828501620005a9565b91505092915050565b620006868162000775565b82525050565b60006200069b60138362000764565b91507f456d7074792076616c696461746f7220736574000000000000000000000000006000830152602082019050919050565b6000602082019050620006e560008301846200067b565b92915050565b6000602082019050818103600083015262000706816200068c565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200073157600080fd5b8060405250919050565b600067ffffffffffffffff8211156200075357600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620007828262000789565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620007b48162000775565b8114620007c057600080fd5b50565b6122f180620007d36000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b1461011e5780639300c9261461013c578063d6119e7014610158578063f2fde38b14610174578063fe82d0821461019057610093565b806335aa2e4414610098578063715018a6146100c85780637194d4b2146100d25780638be10194146100ee575b600080fd5b6100b260048036038101906100ad9190611633565b6101ae565b6040516100bf9190611dae565b60405180910390f35b6100d06101ea565b005b6100ec60048036038101906100e7919061165c565b61033f565b005b61010860048036038101906101039190611633565b6105fc565b6040516101159190611e0e565b60405180910390f35b6101266106b5565b6040516101339190611dae565b60405180910390f35b610156600480360381019061015191906115ee565b6106de565b005b610172600480360381019061016d919061159a565b61081a565b005b61018e60048036038101906101899190611571565b610ac3565b005b610198610c87565b6040516101a59190611dae565b60405180910390f35b600181815481106101bb57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6101f2610cad565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027790611f70565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c690611f30565b60405180910390fd5b6000801515600560009054906101000a900460ff16151514610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d90611f10565b60405180910390fd5b60608585905067ffffffffffffffff8111801561044257600080fd5b5060405190808252806020026020018201604052801561047657816020015b60608152602001906001900390816104615790505b50905060008090505b8686905081101561050f5786868281811061049657fe5b90506020028101906104a89190612027565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508282815181106104f757fe5b6020026020010181905250808060010191505061047f565b50604051806040016040528088815260200182815250600660008201518160000155602082015181600101908051906020019061054d9291906111ae565b509050506001600560006101000a81548160ff0219169083151502179055507f08f3dc9ea534e005aacb952cad2a7d85b6d631123cba3ee204bcc4ddb0aa5d4e878260405161059d929190611fc7565b60405180910390a16105f33385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061081a565b50505050505050565b6002818154811061060957fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106ad5780601f10610682576101008083540402835291602001916106ad565b820191906000526020600020905b81548152906001019060200180831161069057829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106e6610cad565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610774576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076b90611f70565b60405180910390fd5b6000801515600560009054906101000a900460ff161515146107cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c290611f10565b60405180910390fd5b610815838380806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050610cb5565b505050565b6001801515600560009054906101000a900460ff16151514610871576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086890611f10565b60405180910390fd5b60006001805490509050600080600090505b82811015610938576001818154811061089857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561092b57846002828154811061090457fe5b90600052602060002001908051906020019061092192919061120e565b5060019150610938565b8080600101915050610883565b5080610979576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097090611ed0565b60405180910390fd5b60006006600001546006600101604051602001610997929190611ff7565b60405160208183030381529060405280519060200120905060006109ba82610df5565b90508673ffffffffffffffffffffffffffffffffffffffff166109dd8288610e25565b73ffffffffffffffffffffffffffffffffffffffff1614610a33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2a90611e50565b60405180910390fd5b600860008154809291906001019190505550600060048510610a5f576002850260036008540211610a65565b84600854145b90508015610ab9577fd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc125360066002604051610aa0929190611f90565b60405180910390a1610ab0611016565b610ab86110bf565b5b5050505050505050565b610acb610cad565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b59576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5090611f70565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610bc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc090611eb0565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b6000815111610cf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf090611e90565b60405180910390fd5b8060019080519060200190610d0f92919061128e565b5060006004819055506001600081548110610d2657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610de29190611dae565b60405180910390a1610df2611016565b50565b600081604051602001610e089190611d88565b604051602081830303815290604052805190602001209050919050565b60006041825114610e6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6290611e70565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610eee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee590611ef0565b60405180910390fd5b601b8160ff1614158015610f065750601c8160ff1614155b15610f46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3d90611f50565b60405180910390fd5b600060018783868660405160008152602001604052604051610f6b9493929190611dc9565b6020604051602081039080840390855afa158015610f8d573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611009576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100090611e30565b60405180910390fd5b8094505050505092915050565b6000600560006101000a81548160ff02191690831515021790555060018054905067ffffffffffffffff8111801561104d57600080fd5b5060405190808252806020026020018201604052801561108157816020015b606081526020019060019003908161106c5790505b50600290805190602001906110979291906111ae565b5060066000808201600090556001820160006110b39190611318565b50506000600881905550565b600180549050600160045401816110d257fe5b066004819055506001600454815481106110e857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2249b0c82a2ebea913220f5abb96055181481cd617cd96e91b957c25bc570693600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516111a49190611dae565b60405180910390a1565b8280548282559060005260206000209081019282156111fd579160200282015b828111156111fc5782518290805190602001906111ec92919061120e565b50916020019190600101906111ce565b5b50905061120a9190611339565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061124f57805160ff191683800117855561127d565b8280016001018555821561127d579182015b8281111561127c578251825591602001919060010190611261565b5b50905061128a9190611365565b5090565b828054828255906000526020600020908101928215611307579160200282015b828111156113065782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906112ae565b5b509050611314919061138a565b5090565b50805460008255906000526020600020908101906113369190611339565b50565b61136291905b8082111561135e576000818161135591906113cd565b5060010161133f565b5090565b90565b61138791905b8082111561138357600081600090555060010161136b565b5090565b90565b6113ca91905b808211156113c657600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611390565b5090565b90565b50805460018160011615610100020316600290046000825580601f106113f35750611412565b601f0160209004906000526020600020908101906114119190611365565b5b50565b6000813590506114248161228d565b92915050565b60008083601f84011261143c57600080fd5b8235905067ffffffffffffffff81111561145557600080fd5b60208301915083602082028301111561146d57600080fd5b9250929050565b60008083601f84011261148657600080fd5b8235905067ffffffffffffffff81111561149f57600080fd5b6020830191508360208202830111156114b757600080fd5b9250929050565b60008083601f8401126114d057600080fd5b8235905067ffffffffffffffff8111156114e957600080fd5b60208301915083600182028301111561150157600080fd5b9250929050565b600082601f83011261151957600080fd5b813561152c611527826120ab565b61207e565b9150808252602083016020830185838301111561154857600080fd5b611553838284612209565b50505092915050565b60008135905061156b816122a4565b92915050565b60006020828403121561158357600080fd5b600061159184828501611415565b91505092915050565b600080604083850312156115ad57600080fd5b60006115bb85828601611415565b925050602083013567ffffffffffffffff8111156115d857600080fd5b6115e485828601611508565b9150509250929050565b6000806020838503121561160157600080fd5b600083013567ffffffffffffffff81111561161b57600080fd5b6116278582860161142a565b92509250509250929050565b60006020828403121561164557600080fd5b60006116538482850161155c565b91505092915050565b60008060008060006060868803121561167457600080fd5b60006116828882890161155c565b955050602086013567ffffffffffffffff81111561169f57600080fd5b6116ab88828901611474565b9450945050604086013567ffffffffffffffff8111156116ca57600080fd5b6116d6888289016114be565b92509250509295509295909350565b60006116f1838361189f565b905092915050565b60006117058383611911565b905092915050565b611716816121b6565b82525050565b600061172782612111565b611731818561215d565b935083602082028501611743856120d7565b8060005b8581101561177f578484038952815161176085826116e5565b945061176b83612132565b925060208a01995050600181019050611747565b50829750879550505050505092915050565b600061179c8261211c565b6117a6818561214c565b9350836020820285016117b8856120e7565b8060005b858110156117f3578484038952816117d485826116f9565b94506117df8361213f565b925060208a019950506001810190506117bc565b50829750879550505050505092915050565b60006118108261211c565b61181a818561215d565b93508360208202850161182c856120e7565b8060005b858110156118675784840389528161184885826116f9565b94506118538361213f565b925060208a01995050600181019050611830565b50829750879550505050505092915050565b611882816121c8565b82525050565b611899611894826121c8565b612265565b82525050565b60006118aa82612127565b6118b4818561216e565b93506118c4818560208601612218565b6118cd8161226f565b840191505092915050565b60006118e382612127565b6118ed818561217f565b93506118fd818560208601612218565b6119068161226f565b840191505092915050565b60008154600181166000811461192e576001811461195457611998565b607f600283041661193f818761216e565b955060ff198316865260208601935050611998565b60028204611962818761216e565b955061196d856120fc565b60005b8281101561198f57815481890152600182019150602081019050611970565b80880195505050505b505092915050565b60006119ad601883612190565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b60006119ed601583612190565b91507f5369676e617475726520697320696e76616c69642100000000000000000000006000830152602082019050919050565b6000611a2d601f83612190565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b6000611a6d601383612190565b91507f456d7074792076616c696461746f7220736574000000000000000000000000006000830152602082019050919050565b6000611aad601c836121a1565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611aed602683612190565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611b53601a83612190565b91507f5369676e6572206d75737420626520612076616c696461746f720000000000006000830152602082019050919050565b6000611b93602283612190565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611bf9601783612190565b91507f496e76616c69642070726f706f73616c207374617475730000000000000000006000830152602082019050919050565b6000611c39602083612190565b91507f4f6e6c792070726f706f736572206d617920706572666f726d20616374696f6e6000830152602082019050919050565b6000611c79602283612190565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611cdf602083612190565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000604083016000808401549050611d298161224b565b611d366000870182611d5b565b50600184018583036020870152611d4d8382611791565b925050819250505092915050565b611d64816121f2565b82525050565b611d73816121f2565b82525050565b611d82816121fc565b82525050565b6000611d9382611aa0565b9150611d9f8284611888565b60208201915081905092915050565b6000602082019050611dc3600083018461170d565b92915050565b6000608082019050611dde6000830187611879565b611deb6020830186611d79565b611df86040830185611879565b611e056060830184611879565b95945050505050565b60006020820190508181036000830152611e2881846118d8565b905092915050565b60006020820190508181036000830152611e49816119a0565b9050919050565b60006020820190508181036000830152611e69816119e0565b9050919050565b60006020820190508181036000830152611e8981611a20565b9050919050565b60006020820190508181036000830152611ea981611a60565b9050919050565b60006020820190508181036000830152611ec981611ae0565b9050919050565b60006020820190508181036000830152611ee981611b46565b9050919050565b60006020820190508181036000830152611f0981611b86565b9050919050565b60006020820190508181036000830152611f2981611bec565b9050919050565b60006020820190508181036000830152611f4981611c2c565b9050919050565b60006020820190508181036000830152611f6981611c6c565b9050919050565b60006020820190508181036000830152611f8981611cd2565b9050919050565b60006040820190508181036000830152611faa8185611d12565b90508181036020830152611fbe8184611805565b90509392505050565b6000604082019050611fdc6000830185611d6a565b8181036020830152611fee818461171c565b90509392505050565b600060408201905061200c6000830185611d6a565b818103602083015261201e8184611805565b90509392505050565b6000808335600160200384360303811261204057600080fd5b80840192508235915067ffffffffffffffff82111561205e57600080fd5b60208301925060018202360383131561207657600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff821117156120a157600080fd5b8060405250919050565b600067ffffffffffffffff8211156120c257600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081549050919050565b600081519050919050565b6000602082019050919050565b6000600182019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000819050919050565b60006121c1826121d2565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561223657808201518184015260208101905061221b565b83811115612245576000848401525b50505050565b600061225e61225983612280565b6121ac565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160001c9050919050565b612296816121b6565b81146122a157600080fd5b50565b6122ad816121f2565b81146122b857600080fd5b5056fea26469706673582212203194a2b78b522b34f189ae3ee9dfb91773af59e863789ab8b788cf7bf08f154964736f6c63430006060033"

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
// Solidity: event BlockConsensusReached((uint256,bytes[]) proposal, bytes[] signatures)
func (_BlockCommittee *BlockCommitteeFilterer) FilterBlockConsensusReached(opts *bind.FilterOpts) (*BlockCommitteeBlockConsensusReachedIterator, error) {

	logs, sub, err := _BlockCommittee.contract.FilterLogs(opts, "BlockConsensusReached")
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeBlockConsensusReachedIterator{contract: _BlockCommittee.contract, event: "BlockConsensusReached", logs: logs, sub: sub}, nil
}

// WatchBlockConsensusReached is a free log subscription operation binding the contract event 0xd7caf9f2fc3f80c9c139f02c1fe8d4dea9f3e20d4afbca1df37f080630cc1253.
//
// Solidity: event BlockConsensusReached((uint256,bytes[]) proposal, bytes[] signatures)
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
// Solidity: event BlockConsensusReached((uint256,bytes[]) proposal, bytes[] signatures)
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
	BlockNumber *big.Int
	Transitions [][]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x08f3dc9ea534e005aacb952cad2a7d85b6d631123cba3ee204bcc4ddb0aa5d4e.
//
// Solidity: event BlockProposed(uint256 blockNumber, bytes[] transitions)
func (_BlockCommittee *BlockCommitteeFilterer) FilterBlockProposed(opts *bind.FilterOpts) (*BlockCommitteeBlockProposedIterator, error) {

	logs, sub, err := _BlockCommittee.contract.FilterLogs(opts, "BlockProposed")
	if err != nil {
		return nil, err
	}
	return &BlockCommitteeBlockProposedIterator{contract: _BlockCommittee.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0x08f3dc9ea534e005aacb952cad2a7d85b6d631123cba3ee204bcc4ddb0aa5d4e.
//
// Solidity: event BlockProposed(uint256 blockNumber, bytes[] transitions)
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

// ParseBlockProposed is a log parse operation binding the contract event 0x08f3dc9ea534e005aacb952cad2a7d85b6d631123cba3ee204bcc4ddb0aa5d4e.
//
// Solidity: event BlockProposed(uint256 blockNumber, bytes[] transitions)
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
