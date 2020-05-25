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



// DepositWithdrawManagerABI is the input ABI used to generate the binding from.
const DepositWithdrawManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupChainAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_transitionEvaluatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_registerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_depositSignature\",\"type\":\"bytes\"}],\"name\":\"registerAndDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transition\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transitionIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.TransitionInclusionProof\",\"name\":\"inclusionProof\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.IncludedTransition\",\"name\":\"_includedTransition\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DepositWithdrawManagerBin is the compiled bytecode used for deploying new contracts.
var DepositWithdrawManagerBin = "0x60806040523480156200001157600080fd5b506040516200213c3803806200213c83398181016040528101906200003791906200015c565b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000216565b6000815190506200015681620001fc565b92915050565b600080600080608085870312156200017357600080fd5b6000620001838782880162000145565b9450506020620001968782880162000145565b9350506040620001a98782880162000145565b9250506060620001bc8782880162000145565b91505092959194509250565b6000620001d582620001dc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200020781620001c8565b81146200021357600080fd5b50565b611f1680620002266000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633f0b37d91461005c5780634d4fc8dd1461008c5780635753321d146100a8578063647bbada146100c4578063e9c670ad146100f4575b600080fd5b610076600480360381019061007191906110e8565b610110565b6040516100839190611c6f565b60405180910390f35b6100a660048036038101906100a1919061124f565b610135565b005b6100c260048036038101906100bd9190611124565b6106b7565b005b6100de60048036038101906100d991906110e8565b6107a1565b6040516100eb9190611c6f565b60405180910390f35b61010e600480360381019061010991906111d4565b6107c6565b005b6000602052816000526040600020602052806000526040600020600091509150505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ac820cb884846040518363ffffffff1660e01b8152600401610192929190611a6f565b60206040518083038186803b1580156101aa57600080fd5b505afa1580156101be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e291906112ce565b610221576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021890611c2f565b60405180910390fd5b610229610cbb565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc354f6784600001516040518263ffffffff1660e01b81526004016102889190611b0d565b60006040518083038186803b1580156102a057600080fd5b505afa1580156102b4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102dd91906112f7565b90506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639bf53f1883606001516040518263ffffffff1660e01b81526004016103409190611c6f565b60206040518083038186803b15801561035857600080fd5b505afa15801561036c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039091906110bf565b905060008260a001519050600083608001519050600030888584866040516020016103bf959493929190611976565b60405160208183030381529060405280519060200120905060006103e282610a9a565b90508873ffffffffffffffffffffffffffffffffffffffff166104058289610aca565b73ffffffffffffffffffffffffffffffffffffffff161461045b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045290611c4f565b60405180910390fd5b600160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548414610519576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051090611c0f565b60405180910390fd5b600160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919060010191905055508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8a856040518363ffffffff1660e01b81526004016105e0929190611a9f565b602060405180830381600087803b1580156105fa57600080fd5b505af115801561060e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063291906112ce565b610671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066890611b8f565b60405180910390fd5b7f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e56208986856040516106a493929190611a06565b60405180910390a1505050505050505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f2cd61c98886866040518463ffffffff1660e01b815260040161071693929190611a3d565b600060405180830381600087803b15801561073057600080fd5b505af1158015610744573d6000803e3d6000fd5b5050505061079887878785858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506107c6565b50505050505050565b6001602052816000526040600020602052806000526040600020600091509150505481565b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000308686868560405160200161086195949392919061190c565b604051602081830303815290604052805190602001209050600061088482610a9a565b90508673ffffffffffffffffffffffffffffffffffffffff166108a78286610aca565b73ffffffffffffffffffffffffffffffffffffffff16146108fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f490611bcf565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff166323b872dd8830886040518463ffffffff1660e01b815260040161093a93929190611a06565b602060405180830381600087803b15801561095457600080fd5b505af1158015610968573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098c91906112ce565b6109cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c290611b6f565b60405180910390fd5b7ff1444b5cad7ce70cb018d1b8edc8618fe303f3c7f034d8d572a6e27facbf2bef8787876040516109fe93929190611a06565b60405180910390a16000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050555050505050505050565b600081604051602001610aad91906119e0565b604051602081830303815290604052805190602001209050919050565b60006041825114610b10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0790611b4f565b60405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610b93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8a90611baf565b60405180910390fd5b601b8160ff1614158015610bab5750601c8160ff1614155b15610beb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be290611bef565b60405180910390fd5b600060018783868660405160008152602001604052604051610c109493929190611ac8565b6020604051602081039080840390855afa158015610c32573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610cae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ca590611b2f565b60405180910390fd5b8094505050505092915050565b6040518060e00160405280600081526020016000801916815260200160008152602001600081526020016000815260200160008152602001606081525090565b600081359050610d0a81611e84565b92915050565b600081519050610d1f81611e84565b92915050565b600082601f830112610d3657600080fd5b8135610d49610d4482611cb7565b611c8a565b91508181835260208401935060208101905083856020840282011115610d6e57600080fd5b60005b83811015610d9e5781610d848882610dbd565b845260208401935060208301925050600181019050610d71565b5050505092915050565b600081519050610db781611e9b565b92915050565b600081359050610dcc81611eb2565b92915050565b600081519050610de181611eb2565b92915050565b60008083601f840112610df957600080fd5b8235905067ffffffffffffffff811115610e1257600080fd5b602083019150836001820283011115610e2a57600080fd5b9250929050565b600082601f830112610e4257600080fd5b8135610e55610e5082611cdf565b611c8a565b91508082526020830160208301858383011115610e7157600080fd5b610e7c838284611dec565b50505092915050565b600082601f830112610e9657600080fd5b8151610ea9610ea482611cdf565b611c8a565b91508082526020830160208301858383011115610ec557600080fd5b610ed0838284611dfb565b50505092915050565b600060408284031215610eeb57600080fd5b610ef56040611c8a565b9050600082013567ffffffffffffffff811115610f1157600080fd5b610f1d84828501610e31565b600083015250602082013567ffffffffffffffff811115610f3d57600080fd5b610f4984828501610f55565b60208301525092915050565b600060608284031215610f6757600080fd5b610f716060611c8a565b90506000610f8184828501611095565b6000830152506020610f9584828501611095565b602083015250604082013567ffffffffffffffff811115610fb557600080fd5b610fc184828501610d25565b60408301525092915050565b600060e08284031215610fdf57600080fd5b610fe960e0611c8a565b90506000610ff9848285016110aa565b600083015250602061100d84828501610dd2565b6020830152506040611021848285016110aa565b6040830152506060611035848285016110aa565b6060830152506080611049848285016110aa565b60808301525060a061105d848285016110aa565b60a08301525060c082015167ffffffffffffffff81111561107d57600080fd5b61108984828501610e85565b60c08301525092915050565b6000813590506110a481611ec9565b92915050565b6000815190506110b981611ec9565b92915050565b6000602082840312156110d157600080fd5b60006110df84828501610d10565b91505092915050565b600080604083850312156110fb57600080fd5b600061110985828601610cfb565b925050602061111a85828601610cfb565b9150509250929050565b600080600080600080600060a0888a03121561113f57600080fd5b600061114d8a828b01610cfb565b975050602061115e8a828b01610cfb565b965050604061116f8a828b01611095565b955050606088013567ffffffffffffffff81111561118c57600080fd5b6111988a828b01610de7565b9450945050608088013567ffffffffffffffff8111156111b757600080fd5b6111c38a828b01610de7565b925092505092959891949750929550565b600080600080608085870312156111ea57600080fd5b60006111f887828801610cfb565b945050602061120987828801610cfb565b935050604061121a87828801611095565b925050606085013567ffffffffffffffff81111561123757600080fd5b61124387828801610e31565b91505092959194509250565b60008060006060848603121561126457600080fd5b600061127286828701610cfb565b935050602084013567ffffffffffffffff81111561128f57600080fd5b61129b86828701610ed9565b925050604084013567ffffffffffffffff8111156112b857600080fd5b6112c486828701610e31565b9150509250925092565b6000602082840312156112e057600080fd5b60006112ee84828501610da8565b91505092915050565b60006020828403121561130957600080fd5b600082015167ffffffffffffffff81111561132357600080fd5b61132f84828501610fcd565b91505092915050565b600061134483836113d4565b60208301905092915050565b61135981611d8d565b82525050565b61137061136b82611d8d565b611e2e565b82525050565b600061138182611d1b565b61138b8185611d3e565b935061139683611d0b565b8060005b838110156113c75781516113ae8882611338565b97506113b983611d31565b92505060018101905061139a565b5085935050505092915050565b6113dd81611dab565b82525050565b6113ec81611dab565b82525050565b6114036113fe82611dab565b611e40565b82525050565b60006114158385611d60565b9350611422838584611dec565b61142b83611e66565b840190509392505050565b600061144182611d26565b61144b8185611d4f565b935061145b818560208601611dfb565b61146481611e66565b840191505092915050565b600061147a82611d26565b6114848185611d60565b9350611494818560208601611dfb565b61149d81611e66565b840191505092915050565b60006114b5601883611d71565b91507f45434453413a20696e76616c6964207369676e617475726500000000000000006000830152602082019050919050565b60006114f5601f83611d71565b91507f45434453413a20696e76616c6964207369676e6174757265206c656e677468006000830152602082019050919050565b6000611535601c83611d82565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611575600e83611d71565b91507f4465706f736974206661696c65640000000000000000000000000000000000006000830152602082019050919050565b60006115b5600f83611d71565b91507f5769746864726177206661696c656400000000000000000000000000000000006000830152602082019050919050565b60006115f5600783611d82565b91507f6465706f736974000000000000000000000000000000000000000000000000006000830152600782019050919050565b6000611635602283611d71565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061169b601d83611d71565b91507f4465706f736974207369676e617475726520697320696e76616c6964210000006000830152602082019050919050565b60006116db602283611d71565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f75650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611741600883611d82565b91507f77697468647261770000000000000000000000000000000000000000000000006000830152600882019050919050565b6000611781601483611d71565b91507f57726f6e67207769746864726177206e6f6e63650000000000000000000000006000830152602082019050919050565b60006117c1601b83611d71565b91507f5769746864726177207472616e736974696f6e20696e76616c696400000000006000830152602082019050919050565b6000611801601e83611d71565b91507f5769746864726177207369676e617475726520697320696e76616c69642100006000830152602082019050919050565b600060408301600083015184820360008601526118518282611436565b9150506020830151848203602086015261186b8282611878565b9150508091505092915050565b600060608301600083015161189060008601826118c8565b5060208301516118a360208601826118c8565b50604083015184820360408601526118bb8282611376565b9150508091505092915050565b6118d181611dd5565b82525050565b6118e081611dd5565b82525050565b6118f76118f282611dd5565b611e5c565b82525050565b61190681611ddf565b82525050565b6000611918828861135f565b601482019150611927826115e8565b9150611933828761135f565b601482019150611943828661135f565b60148201915061195382856118e6565b60208201915061196382846118e6565b6020820191508190509695505050505050565b6000611982828861135f565b60148201915061199182611734565b915061199d828761135f565b6014820191506119ad828661135f565b6014820191506119bd82856118e6565b6020820191506119cd82846118e6565b6020820191508190509695505050505050565b60006119eb82611528565b91506119f782846113f2565b60208201915081905092915050565b6000606082019050611a1b6000830186611350565b611a286020830185611350565b611a3560408301846118d7565b949350505050565b6000604082019050611a526000830186611350565b8181036020830152611a65818486611409565b9050949350505050565b6000604082019050611a846000830185611350565b8181036020830152611a968184611834565b90509392505050565b6000604082019050611ab46000830185611350565b611ac160208301846118d7565b9392505050565b6000608082019050611add60008301876113e3565b611aea60208301866118fd565b611af760408301856113e3565b611b0460608301846113e3565b95945050505050565b60006020820190508181036000830152611b27818461146f565b905092915050565b60006020820190508181036000830152611b48816114a8565b9050919050565b60006020820190508181036000830152611b68816114e8565b9050919050565b60006020820190508181036000830152611b8881611568565b9050919050565b60006020820190508181036000830152611ba8816115a8565b9050919050565b60006020820190508181036000830152611bc881611628565b9050919050565b60006020820190508181036000830152611be88161168e565b9050919050565b60006020820190508181036000830152611c08816116ce565b9050919050565b60006020820190508181036000830152611c2881611774565b9050919050565b60006020820190508181036000830152611c48816117b4565b9050919050565b60006020820190508181036000830152611c68816117f4565b9050919050565b6000602082019050611c8460008301846118d7565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611cad57600080fd5b8060405250919050565b600067ffffffffffffffff821115611cce57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115611cf657600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611d9882611db5565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611e19578082015181840152602081019050611dfe565b83811115611e28576000848401525b50505050565b6000611e3982611e4a565b9050919050565b6000819050919050565b6000611e5582611e77565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b611e8d81611d8d565b8114611e9857600080fd5b50565b611ea481611d9f565b8114611eaf57600080fd5b50565b611ebb81611dab565b8114611ec657600080fd5b50565b611ed281611dd5565b8114611edd57600080fd5b5056fea2646970667358221220fd3153f9997567dec9828b0849aedcf787f0fdd1219059fa53a180541eefbd1864736f6c63430006060033"

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

// DepositNonces is a free data retrieval call binding the contract method 0x3f0b37d9.
//
// Solidity: function depositNonces(address , address ) view returns(uint256)
func (_DepositWithdrawManager *DepositWithdrawManagerCaller) DepositNonces(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositWithdrawManager.contract.Call(opts, out, "depositNonces", arg0, arg1)
	return *ret0, err
}

// DepositNonces is a free data retrieval call binding the contract method 0x3f0b37d9.
//
// Solidity: function depositNonces(address , address ) view returns(uint256)
func (_DepositWithdrawManager *DepositWithdrawManagerSession) DepositNonces(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DepositWithdrawManager.Contract.DepositNonces(&_DepositWithdrawManager.CallOpts, arg0, arg1)
}

// DepositNonces is a free data retrieval call binding the contract method 0x3f0b37d9.
//
// Solidity: function depositNonces(address , address ) view returns(uint256)
func (_DepositWithdrawManager *DepositWithdrawManagerCallerSession) DepositNonces(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DepositWithdrawManager.Contract.DepositNonces(&_DepositWithdrawManager.CallOpts, arg0, arg1)
}

// WithdrawNonces is a free data retrieval call binding the contract method 0x647bbada.
//
// Solidity: function withdrawNonces(address , address ) view returns(uint256)
func (_DepositWithdrawManager *DepositWithdrawManagerCaller) WithdrawNonces(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositWithdrawManager.contract.Call(opts, out, "withdrawNonces", arg0, arg1)
	return *ret0, err
}

// WithdrawNonces is a free data retrieval call binding the contract method 0x647bbada.
//
// Solidity: function withdrawNonces(address , address ) view returns(uint256)
func (_DepositWithdrawManager *DepositWithdrawManagerSession) WithdrawNonces(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DepositWithdrawManager.Contract.WithdrawNonces(&_DepositWithdrawManager.CallOpts, arg0, arg1)
}

// WithdrawNonces is a free data retrieval call binding the contract method 0x647bbada.
//
// Solidity: function withdrawNonces(address , address ) view returns(uint256)
func (_DepositWithdrawManager *DepositWithdrawManagerCallerSession) WithdrawNonces(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DepositWithdrawManager.Contract.WithdrawNonces(&_DepositWithdrawManager.CallOpts, arg0, arg1)
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

// Withdraw is a paid mutator transaction binding the contract method 0x4d4fc8dd.
//
// Solidity: function withdraw(address _account, (bytes,(uint256,uint256,bytes32[])) _includedTransition, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactor) Withdraw(opts *bind.TransactOpts, _account common.Address, _includedTransition DataTypesIncludedTransition, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.contract.Transact(opts, "withdraw", _account, _includedTransition, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4d4fc8dd.
//
// Solidity: function withdraw(address _account, (bytes,(uint256,uint256,bytes32[])) _includedTransition, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerSession) Withdraw(_account common.Address, _includedTransition DataTypesIncludedTransition, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Withdraw(&_DepositWithdrawManager.TransactOpts, _account, _includedTransition, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4d4fc8dd.
//
// Solidity: function withdraw(address _account, (bytes,(uint256,uint256,bytes32[])) _includedTransition, bytes _signature) returns()
func (_DepositWithdrawManager *DepositWithdrawManagerTransactorSession) Withdraw(_account common.Address, _includedTransition DataTypesIncludedTransition, _signature []byte) (*types.Transaction, error) {
	return _DepositWithdrawManager.Contract.Withdraw(&_DepositWithdrawManager.TransactOpts, _account, _includedTransition, _signature)
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
