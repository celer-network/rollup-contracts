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
const DepositWithdrawManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupChainAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_transitionEvaluatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_registerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_depositSignature\",\"type\":\"bytes\"}],\"name\":\"registerAndDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_preStateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transition\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transitionIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.TransitionInclusionProof\",\"name\":\"inclusionProof\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.IncludedTransition\",\"name\":\"_includedTransition\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slotIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"transferNonces\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"withdrawNonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataTypes.AccountInfo\",\"name\":\"value\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.StorageSlot\",\"name\":\"storageSlot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.IncludedStorageSlot\",\"name\":\"_transitionStorageSlot\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositWithdrawManagerBin is the compiled bytecode used for deploying new contracts.
var DepositWithdrawManagerBin = "0x60806040523480156200001157600080fd5b50604051620023e3380380620023e383398181016040528101906200003791906200015c565b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000216565b6000815190506200015681620001fc565b92915050565b600080600080608085870312156200017357600080fd5b6000620001838782880162000145565b9450506020620001968782880162000145565b9350506040620001a98782880162000145565b9250506060620001bc8782880162000145565b91505092959194509250565b6000620001d582620001dc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200020781620001c8565b81146200021357600080fd5b50565b6121bd80620002266000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635753321d14610046578063e799caac14610062578063e9c670ad1461007e575b600080fd5b610060600480360381019061005b91906111f0565b61009a565b005b61007c6004803603810190610077919061131b565b610184565b005b610098600480360381019061009391906112a0565b6106af565b005b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f2cd61c98886866040518463ffffffff1660e01b81526004016100f993929190611cd3565b600060405180830381600087803b15801561011357600080fd5b505af1158015610127573d6000803e3d6000fd5b5050505061017b87878785858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506106af565b50505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638d12aeff868686866040518563ffffffff1660e01b81526004016101e59493929190611c80565b600060405180830381600087803b1580156101ff57600080fd5b505af1158015610213573d6000803e3d6000fd5b5050505061021f610ba4565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc354f6785600001516040518263ffffffff1660e01b815260040161027e9190611d73565b60006040518083038186803b15801561029657600080fd5b505afa1580156102aa573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102d39190611403565b90506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639bf53f1883606001516040518263ffffffff1660e01b81526004016103369190611eb5565b60206040518083038186803b15801561034e57600080fd5b505afa158015610362573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038691906111c7565b905060008260a0015190506000836080015190506000308a8584866040516020016103b5959493929190611bb9565b60405160208183030381529060405280519060200120905060006103d882610983565b90508a73ffffffffffffffffffffffffffffffffffffffff166103fb82896109b3565b73ffffffffffffffffffffffffffffffffffffffff1614610451576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044890611e95565b60405180910390fd5b600160008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054841461050f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050690611e75565b60405180910390fd5b600160008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919060010191905055508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8c856040518363ffffffff1660e01b81526004016105d6929190611d05565b602060405180830381600087803b1580156105f057600080fd5b505af1158015610604573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062891906113da565b610667576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065e90611df5565b60405180910390fd5b7f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e56208b868560405161069a93929190611c49565b60405180910390a15050505050505050505050565b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000308686868560405160200161074a959493929190611b4f565b604051602081830303815290604052805190602001209050600061076d82610983565b90508673ffffffffffffffffffffffffffffffffffffffff1661079082866109b3565b73ffffffffffffffffffffffffffffffffffffffff16146107e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107dd90611e35565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff166323b872dd8830886040518463ffffffff1660e01b815260040161082393929190611c49565b602060405180830381600087803b15801561083d57600080fd5b505af1158015610851573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087591906113da565b6108b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ab90611dd5565b60405180910390fd5b7ff1444b5cad7ce70cb018d1b8edc8618fe303f3c7f034d8d572a6e27facbf2bef8787876040516108e793929190611c49565b60405180910390a16000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050555050505050505050565b6000816040516020016109969190611c23565b604051602081830303815290604052805190602001209050919050565b600060418251146109f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f090611db5565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610a7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7390611e15565b60405180910390fd5b601b8160ff1614158015610a945750601c8160ff1614155b15610ad4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acb90611e55565b60405180910390fd5b600060018783868660405160008152602001604052604051610af99493929190611d2e565b6020604051602081039080840390855afa158015610b1b573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610b97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8e90611d95565b60405180910390fd5b8094505050505092915050565b6040518060e00160405280600081526020016000801916815260200160008152602001600081526020016000815260200160008152602001606081525090565b600081359050610bf38161212b565b92915050565b600081519050610c088161212b565b92915050565b600082601f830112610c1f57600080fd5b8135610c32610c2d82611efd565b611ed0565b91508181835260208401935060208101905083856020840282011115610c5757600080fd5b60005b83811015610c875781610c6d8882610d29565b845260208401935060208301925050600181019050610c5a565b5050505092915050565b600082601f830112610ca257600080fd5b8135610cb5610cb082611f25565b611ed0565b91508181835260208401935060208101905083856020840282011115610cda57600080fd5b60005b83811015610d0a5781610cf0888261119d565b845260208401935060208301925050600181019050610cdd565b5050505092915050565b600081519050610d2381612142565b92915050565b600081359050610d3881612159565b92915050565b600081519050610d4d81612159565b92915050565b60008083601f840112610d6557600080fd5b8235905067ffffffffffffffff811115610d7e57600080fd5b602083019150836001820283011115610d9657600080fd5b9250929050565b600082601f830112610dae57600080fd5b8135610dc1610dbc82611f4d565b611ed0565b91508082526020830160208301858383011115610ddd57600080fd5b610de8838284612093565b50505092915050565b600082601f830112610e0257600080fd5b8151610e15610e1082611f4d565b611ed0565b91508082526020830160208301858383011115610e3157600080fd5b610e3c8382846120a2565b50505092915050565b600060808284031215610e5757600080fd5b610e616080611ed0565b90506000610e7184828501610be4565b600083015250602082013567ffffffffffffffff811115610e9157600080fd5b610e9d84828501610c91565b602083015250604082013567ffffffffffffffff811115610ebd57600080fd5b610ec984828501610c91565b604083015250606082013567ffffffffffffffff811115610ee957600080fd5b610ef584828501610c91565b60608301525092915050565b600060408284031215610f1357600080fd5b610f1d6040611ed0565b9050600082013567ffffffffffffffff811115610f3957600080fd5b610f4584828501610ff9565b600083015250602082013567ffffffffffffffff811115610f6557600080fd5b610f7184828501610c0e565b60208301525092915050565b600060408284031215610f8f57600080fd5b610f996040611ed0565b9050600082013567ffffffffffffffff811115610fb557600080fd5b610fc184828501610d9d565b600083015250602082013567ffffffffffffffff811115610fe157600080fd5b610fed8482850161105d565b60208301525092915050565b60006040828403121561100b57600080fd5b6110156040611ed0565b905060006110258482850161119d565b600083015250602082013567ffffffffffffffff81111561104557600080fd5b61105184828501610e45565b60208301525092915050565b60006060828403121561106f57600080fd5b6110796060611ed0565b905060006110898482850161119d565b600083015250602061109d8482850161119d565b602083015250604082013567ffffffffffffffff8111156110bd57600080fd5b6110c984828501610c0e565b60408301525092915050565b600060e082840312156110e757600080fd5b6110f160e0611ed0565b90506000611101848285016111b2565b600083015250602061111584828501610d3e565b6020830152506040611129848285016111b2565b604083015250606061113d848285016111b2565b6060830152506080611151848285016111b2565b60808301525060a0611165848285016111b2565b60a08301525060c082015167ffffffffffffffff81111561118557600080fd5b61119184828501610df1565b60c08301525092915050565b6000813590506111ac81612170565b92915050565b6000815190506111c181612170565b92915050565b6000602082840312156111d957600080fd5b60006111e784828501610bf9565b91505092915050565b600080600080600080600060a0888a03121561120b57600080fd5b60006112198a828b01610be4565b975050602061122a8a828b01610be4565b965050604061123b8a828b0161119d565b955050606088013567ffffffffffffffff81111561125857600080fd5b6112648a828b01610d53565b9450945050608088013567ffffffffffffffff81111561128357600080fd5b61128f8a828b01610d53565b925092505092959891949750929550565b600080600080608085870312156112b657600080fd5b60006112c487828801610be4565b94505060206112d587828801610be4565b93505060406112e68782880161119d565b925050606085013567ffffffffffffffff81111561130357600080fd5b61130f87828801610d9d565b91505092959194509250565b600080600080600060a0868803121561133357600080fd5b600061134188828901610be4565b955050602061135288828901610d29565b945050604086013567ffffffffffffffff81111561136f57600080fd5b61137b88828901610f7d565b935050606086013567ffffffffffffffff81111561139857600080fd5b6113a488828901610f01565b925050608086013567ffffffffffffffff8111156113c157600080fd5b6113cd88828901610d9d565b9150509295509295909350565b6000602082840312156113ec57600080fd5b60006113fa84828501610d14565b91505092915050565b60006020828403121561141557600080fd5b600082015167ffffffffffffffff81111561142f57600080fd5b61143b848285016110d5565b91505092915050565b60006114508383611565565b60208301905092915050565b60006114688383611b0b565b60208301905092915050565b61147d81612034565b82525050565b61148c81612034565b82525050565b6114a361149e82612034565b6120d5565b82525050565b60006114b482611f99565b6114be8185611fd4565b93506114c983611f79565b8060005b838110156114fa5781516114e18882611444565b97506114ec83611fba565b9250506001810190506114cd565b5085935050505092915050565b600061151282611fa4565b61151c8185611fe5565b935061152783611f89565b8060005b8381101561155857815161153f888261145c565b975061154a83611fc7565b92505060018101905061152b565b5085935050505092915050565b61156e81612052565b82525050565b61157d81612052565b82525050565b61159461158f82612052565b6120e7565b82525050565b60006115a68385612007565b93506115b3838584612093565b6115bc8361210d565b840190509392505050565b60006115d282611faf565b6115dc8185611ff6565b93506115ec8185602086016120a2565b6115f58161210d565b840191505092915050565b600061160b82611faf565b6116158185612007565b93506116258185602086016120a2565b61162e8161210d565b840191505092915050565b6000611646601883612018565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b6000611686601f83612018565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b60006116c6601c83612029565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611706600e83612018565b91507f4465706f736974206661696c65640000000000000000000000000000000000006000830152602082019050919050565b6000611746600f83612018565b91507f5769746864726177206661696c656400000000000000000000000000000000006000830152602082019050919050565b6000611786600783612029565b91507f6465706f736974000000000000000000000000000000000000000000000000006000830152600782019050919050565b60006117c6602283612018565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061182c601d83612018565b91507f4465706f736974207369676e617475726520697320696e76616c6964210000006000830152602082019050919050565b600061186c602283612018565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006118d2600883612029565b91507f77697468647261770000000000000000000000000000000000000000000000006000830152600882019050919050565b6000611912601483612018565b91507f57726f6e67207769746864726177206e6f6e63650000000000000000000000006000830152602082019050919050565b6000611952601e83612018565b91507f5769746864726177207369676e617475726520697320696e76616c69642100006000830152602082019050919050565b600060808301600083015161199d6000860182611474565b50602083015184820360208601526119b58282611507565b915050604083015184820360408601526119cf8282611507565b915050606083015184820360608601526119e98282611507565b9150508091505092915050565b60006040830160008301518482036000860152611a138282611a7e565b91505060208301518482036020860152611a2d82826114a9565b9150508091505092915050565b60006040830160008301518482036000860152611a5782826115c7565b91505060208301518482036020860152611a718282611abb565b9150508091505092915050565b6000604083016000830151611a966000860182611b0b565b5060208301518482036020860152611aae8282611985565b9150508091505092915050565b6000606083016000830151611ad36000860182611b0b565b506020830151611ae66020860182611b0b565b5060408301518482036040860152611afe82826114a9565b9150508091505092915050565b611b148161207c565b82525050565b611b238161207c565b82525050565b611b3a611b358261207c565b612103565b82525050565b611b4981612086565b82525050565b6000611b5b8288611492565b601482019150611b6a82611779565b9150611b768287611492565b601482019150611b868286611492565b601482019150611b968285611b29565b602082019150611ba68284611b29565b6020820191508190509695505050505050565b6000611bc58288611492565b601482019150611bd4826118c5565b9150611be08287611492565b601482019150611bf08286611492565b601482019150611c008285611b29565b602082019150611c108284611b29565b6020820191508190509695505050505050565b6000611c2e826116b9565b9150611c3a8284611583565b60208201915081905092915050565b6000606082019050611c5e6000830186611483565b611c6b6020830185611483565b611c786040830184611b1a565b949350505050565b6000608082019050611c956000830187611483565b611ca26020830186611574565b8181036040830152611cb48185611a3a565b90508181036060830152611cc881846119f6565b905095945050505050565b6000604082019050611ce86000830186611483565b8181036020830152611cfb81848661159a565b9050949350505050565b6000604082019050611d1a6000830185611483565b611d276020830184611b1a565b9392505050565b6000608082019050611d436000830187611574565b611d506020830186611b40565b611d5d6040830185611574565b611d6a6060830184611574565b95945050505050565b60006020820190508181036000830152611d8d8184611600565b905092915050565b60006020820190508181036000830152611dae81611639565b9050919050565b60006020820190508181036000830152611dce81611679565b9050919050565b60006020820190508181036000830152611dee816116f9565b9050919050565b60006020820190508181036000830152611e0e81611739565b9050919050565b60006020820190508181036000830152611e2e816117b9565b9050919050565b60006020820190508181036000830152611e4e8161181f565b9050919050565b60006020820190508181036000830152611e6e8161185f565b9050919050565b60006020820190508181036000830152611e8e81611905565b9050919050565b60006020820190508181036000830152611eae81611945565b9050919050565b6000602082019050611eca6000830184611b1a565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611ef357600080fd5b8060405250919050565b600067ffffffffffffffff821115611f1457600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115611f3c57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115611f6457600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061203f8261205c565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156120c05780820151818401526020810190506120a5565b838111156120cf576000848401525b50505050565b60006120e0826120f1565b9050919050565b6000819050919050565b60006120fc8261211e565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b61213481612034565b811461213f57600080fd5b50565b61214b81612046565b811461215657600080fd5b50565b61216281612052565b811461216d57600080fd5b50565b6121798161207c565b811461218457600080fd5b5056fea2646970667358221220a6996dd14758806af66eb8528b57b5d14f057fb54c849b6d9f4ac72268b7cba664736f6c63430006060033"

// DeployDepositWithdrawManager deploys a new Ethereum contract, binding an instance of DepositWithdrawManager to it.
func DeployDepositWithdrawManager(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupChainAddress common.Address, _transitionEvaluatorAddress common.Address, _accountRegistryAddress common.Address, _tokenRegistryAddress common.Address) (common.Address, *types.Transaction, *DepositWithdrawManager, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositWithdrawManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositWithdrawManagerBin), backend, _rollupChainAddress, _transitionEvaluatorAddress, _accountRegistryAddress, _tokenRegistryAddress)
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

// Deposit is a paid mutator transaction binding the contract method 0xe9c670ad.
//
// Solidity: function deposit(address _account, address _token, uint256 _amount, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactor) Deposit(opts *bind.TransactOpts, _account common.Address, _token common.Address, _amount *big.Int, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.contract.Transact(opts, "deposit", _account, _token, _amount, _signature)
}

// Deposit is a paid mutator transaction binding the contract method 0xe9c670ad.
//
// Solidity: function deposit(address _account, address _token, uint256 _amount, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerSession) Deposit(_account common.Address, _token common.Address, _amount *big.Int, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Deposit(&_DepositWithdrawManager.TransactOpts, _account, _token, _amount, _signature)
}

// Deposit is a paid mutator transaction binding the contract method 0xe9c670ad.
//
// Solidity: function deposit(address _account, address _token, uint256 _amount, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorSession) Deposit(_account common.Address, _token common.Address, _amount *big.Int, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Deposit(&_DepositWithdrawManager.TransactOpts, _account, _token, _amount, _signature)
}

// RegisterAndDeposit is a paid mutator transaction binding the contract method 0x5753321d.
//
// Solidity: function registerAndDeposit(address _account, address _token, uint256 _amount, bytes _registerSignature, bytes _depositSignature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactor) RegisterAndDeposit(opts *bind.TransactOpts, _account common.Address, _token common.Address, _amount *big.Int, _registerSignature []byte, _depositSignature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.contract.Transact(opts, "registerAndDeposit", _account, _token, _amount, _registerSignature, _depositSignature)
}

// RegisterAndDeposit is a paid mutator transaction binding the contract method 0x5753321d.
//
// Solidity: function registerAndDeposit(address _account, address _token, uint256 _amount, bytes _registerSignature, bytes _depositSignature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerSession) RegisterAndDeposit(_account common.Address, _token common.Address, _amount *big.Int, _registerSignature []byte, _depositSignature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.RegisterAndDeposit(&_DepositWithdrawManager.TransactOpts, _account, _token, _amount, _registerSignature, _depositSignature)
}

// RegisterAndDeposit is a paid mutator transaction binding the contract method 0x5753321d.
//
// Solidity: function registerAndDeposit(address _account, address _token, uint256 _amount, bytes _registerSignature, bytes _depositSignature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorSession) RegisterAndDeposit(_account common.Address, _token common.Address, _amount *big.Int, _registerSignature []byte, _depositSignature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.RegisterAndDeposit(&_DepositWithdrawManager.TransactOpts, _account, _token, _amount, _registerSignature, _depositSignature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe799caac.
//
// Solidity: function withdraw(address _account, bytes32 _preStateRoot, DataTypesIncludedTransition _includedTransition, DataTypesIncludedStorageSlot _transitionStorageSlot, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactor) Withdraw(opts *bind.TransactOpts, _account common.Address, _preStateRoot [32]byte, _includedTransition DataTypesIncludedTransition, _transitionStorageSlot DataTypesIncludedStorageSlot, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.contract.Transact(opts, "withdraw", _account, _preStateRoot, _includedTransition, _transitionStorageSlot, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe799caac.
//
// Solidity: function withdraw(address _account, bytes32 _preStateRoot, DataTypesIncludedTransition _includedTransition, DataTypesIncludedStorageSlot _transitionStorageSlot, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerSession) Withdraw(_account common.Address, _preStateRoot [32]byte, _includedTransition DataTypesIncludedTransition, _transitionStorageSlot DataTypesIncludedStorageSlot, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Withdraw(&_DepositWithdrawManager.TransactOpts, _account, _preStateRoot, _includedTransition, _transitionStorageSlot, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe799caac.
//
// Solidity: function withdraw(address _account, bytes32 _preStateRoot, DataTypesIncludedTransition _includedTransition, DataTypesIncludedStorageSlot _transitionStorageSlot, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorSession) Withdraw(_account common.Address, _preStateRoot [32]byte, _includedTransition DataTypesIncludedTransition, _transitionStorageSlot DataTypesIncludedStorageSlot, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Withdraw(&_DepositWithdrawManager.TransactOpts, _account, _preStateRoot, _includedTransition, _transitionStorageSlot, _signature)
}

// DepositWithdrawManagerTokenDepositedIterator is returned from FilterTokenDeposited and is used to iterate over the raw logs and unpacked data for TokenDeposited events raised by the DepositWithdrawManager contract.
type DepositWithdrawManagerTokenDepositedIterator struct {
	Event *DepositWithdrawManagerTokenDeposited // Event containing the contract specifics and raw log

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
func (it *DepositWithdrawManagerTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositWithdrawManagerTokenDeposited)
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
		it.Event = new(DepositWithdrawManagerTokenDeposited)
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
func (it *DepositWithdrawManagerTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositWithdrawManagerTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositWithdrawManagerTokenDeposited represents a TokenDeposited event raised by the DepositWithdrawManager contract.
type DepositWithdrawManagerTokenDeposited struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposited is a free log retrieval operation binding the contract event 0xf1444b5cad7ce70cb018d1b8edc8618fe303f3c7f034d8d572a6e27facbf2bef.
//
// Solidity: event TokenDeposited(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) FilterTokenDeposited(opts *bind.FilterOpts) (*DepositWithdrawManagerTokenDepositedIterator, error) {

	logs, sub, err := _DepositWithdrawManager.contract.FilterLogs(opts, "TokenDeposited")
	if err != nil {
		return nil, err
	}
	return &DepositWithdrawManagerTokenDepositedIterator{contract: _DepositWithdrawManager.contract, event: "TokenDeposited", logs: logs, sub: sub}, nil
}

// WatchTokenDeposited is a free log subscription operation binding the contract event 0xf1444b5cad7ce70cb018d1b8edc8618fe303f3c7f034d8d572a6e27facbf2bef.
//
// Solidity: event TokenDeposited(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) WatchTokenDeposited(opts *bind.WatchOpts, sink chan<- *DepositWithdrawManagerTokenDeposited) (event.Subscription, error) {

	logs, sub, err := _DepositWithdrawManager.contract.WatchLogs(opts, "TokenDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositWithdrawManagerTokenDeposited)
				if err := _DepositWithdrawManager.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
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

// ParseTokenDeposited is a log parse operation binding the contract event 0xf1444b5cad7ce70cb018d1b8edc8618fe303f3c7f034d8d572a6e27facbf2bef.
//
// Solidity: event TokenDeposited(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) ParseTokenDeposited(log types.Log) (*DepositWithdrawManagerTokenDeposited, error) {
	event := new(DepositWithdrawManagerTokenDeposited)
	if err := _DepositWithdrawManager.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositWithdrawManagerTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the DepositWithdrawManager contract.
type DepositWithdrawManagerTokenWithdrawnIterator struct {
	Event *DepositWithdrawManagerTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *DepositWithdrawManagerTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositWithdrawManagerTokenWithdrawn)
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
		it.Event = new(DepositWithdrawManagerTokenWithdrawn)
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
func (it *DepositWithdrawManagerTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositWithdrawManagerTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositWithdrawManagerTokenWithdrawn represents a TokenWithdrawn event raised by the DepositWithdrawManager contract.
type DepositWithdrawManagerTokenWithdrawn struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts) (*DepositWithdrawManagerTokenWithdrawnIterator, error) {

	logs, sub, err := _DepositWithdrawManager.contract.FilterLogs(opts, "TokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return &DepositWithdrawManagerTokenWithdrawnIterator{contract: _DepositWithdrawManager.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DepositWithdrawManagerTokenWithdrawn) (event.Subscription, error) {

	logs, sub, err := _DepositWithdrawManager.contract.WatchLogs(opts, "TokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositWithdrawManagerTokenWithdrawn)
				if err := _DepositWithdrawManager.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address account, address token, uint256 amount)
func (_DepositWithdrawManager *DepositWithdrawManagerFilterer) ParseTokenWithdrawn(log types.Log) (*DepositWithdrawManagerTokenWithdrawn, error) {
	event := new(DepositWithdrawManagerTokenWithdrawn)
	if err := _DepositWithdrawManager.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
