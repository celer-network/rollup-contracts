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

// SidechainERC20ABI is the input ABI used to generate the binding from.
const SidechainERC20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainchainToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainchainToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainchainToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainchainToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mainchainToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SidechainERC20Bin is the compiled bytecode used for deploying new contracts.
var SidechainERC20Bin = "0x60806040523480156200001157600080fd5b506040516200277338038062002773833981810160405260808110156200003757600080fd5b8101908080519060200190929190805160405193929190846401000000008211156200006257600080fd5b838201915060208201858111156200007957600080fd5b82518660018202830111640100000000821117156200009757600080fd5b8083526020830192505050908051906020019080838360005b83811015620000cd578082015181840152602081019050620000b0565b50505050905090810190601f168015620000fb5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011f57600080fd5b838201915060208201858111156200013657600080fd5b82518660018202830111640100000000821117156200015457600080fd5b8083526020830192505050908051906020019080838360005b838110156200018a5780820151818401526020810190506200016d565b50505050905090810190601f168015620001b85780820380516001836020036101000a031916815260200191505b50604052602001805190602001909291905050508282828260039080519060200190620001e79291906200035e565b508160049080519060200190620002009291906200035e565b5080600560006101000a81548160ff021916908360ff1602179055505050506000620002316200035660201b60201c565b905080600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156200030b57600080fd5b83600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506200040d565b600033905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003a157805160ff1916838001178555620003d2565b82800160010185558215620003d2579182015b82811115620003d1578251825591602001919060010190620003b4565b5b509050620003e19190620003e5565b5090565b6200040a91905b8082111562000406576000816000905550600101620003ec565b5090565b90565b612356806200041d6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80637ecebe00116100ad578063a457c2d711610071578063a457c2d7146106e5578063a9059cbb1461074b578063be45fd62146107b1578063dd62ed3e146108ae578063f2fde38b146109265761012c565b80637ecebe00146105545780638da5cb5b146105ac5780638f32d59b146105f657806395d89b41146106185780639cdcac9c1461069b5761012c565b8063313ce567116100f4578063313ce5671461038357806339509351146103a757806349bdc2b81461040d57806370a08231146104f2578063715018a61461054a5761012c565b8063030ba25d1461013157806306fdde03146101f6578063095ea7b31461027957806318160ddd146102df57806323b872dd146102fd575b600080fd5b6101f46004803603604081101561014757600080fd5b81019080803590602001909291908035906020019064010000000081111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460018302840111640100000000831117156101a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061096a565b005b6101fe610b37565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561023e578082015181840152602081019050610223565b50505050905090810190601f16801561026b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102c56004803603604081101561028f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bd9565b604051808215151515815260200191505060405180910390f35b6102e7610c49565b6040518082815260200191505060405180910390f35b6103696004803603606081101561031357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c53565b604051808215151515815260200191505060405180910390f35b61038b610cc3565b604051808260ff1660ff16815260200191505060405180910390f35b6103f3600480360360408110156103bd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cda565b604051808215151515815260200191505060405180910390f35b6104f06004803603606081101561042357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561046a57600080fd5b82018360208201111561047c57600080fd5b8035906020019184600183028401116401000000008311171561049e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610d8d565b005b6105346004803603602081101561050857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ff4565b6040518082815260200191505060405180910390f35b61055261103c565b005b6105966004803603602081101561056a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611177565b6040518082815260200191505060405180910390f35b6105b461118f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105fe6111b9565b604051808215151515815260200191505060405180910390f35b610620611218565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610660578082015181840152602081019050610645565b50505050905090810190601f16801561068d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106a36112ba565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610731600480360360408110156106fb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506112e0565b604051808215151515815260200191505060405180910390f35b6107976004803603604081101561076157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506113ad565b604051808215151515815260200191505060405180910390f35b610894600480360360608110156107c757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561080e57600080fd5b82018360208201111561082057600080fd5b8035906020019184600183028401116401000000008311171561084257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061141d565b604051808215151515815260200191505060405180910390f35b610910600480360360408110156108c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115e9565b6040518082815260200191505060405180910390f35b6109686004803603602081101561093c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611659565b005b600033905060008311801561098757508261098482610ff4565b10155b61099057600080fd5b61099a81846116df565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506109f260018261189790919063ffffffff16565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f33be7eabd8ed368ca1aa14ce2ad1e90a0c9bf21edbb3820d5591546e4eb841578684876040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610af5578082015181840152602081019050610ada565b50505050905090810190601f168015610b225780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a350505050565b606060038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bcf5780601f10610ba457610100808354040283529160200191610bcf565b820191906000526020600020905b815481529060010190602001808311610bb257829003601f168201915b5050505050905090565b60006040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f44697361626c656420666561747572650000000000000000000000000000000081525060200191505060405180910390fd5b6000600254905090565b60006040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f44697361626c656420666561747572650000000000000000000000000000000081525060200191505060405180910390fd5b6000600560009054906101000a900460ff16905090565b6000610d83610ce761191f565b84610d7e8560016000610cf861191f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461189790919063ffffffff16565b611927565b6001905092915050565b610d956111b9565b610e07576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600082118015610e445750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614155b610e4d57600080fd5b610e578383611b1e565b6000600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050610eaf60018261189790919063ffffffff16565b600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f99d353966fe3c3445ee3f115b9ce5bc579b0e19a76b2484f79ae9d42e45fe4ce8584866040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610fb2578082015181840152602081019050610f97565b50505050905090810190601f168015610fdf5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a350505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6110446111b9565b6110b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60076020528060005260406000206000915090505481565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166111fc61191f565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112b05780601f10611285576101008083540402835291602001916112b0565b820191906000526020600020905b81548152906001019060200180831161129357829003601f168201915b5050505050905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006113a36112ed61191f565b8461139e856040518060600160405280602581526020016122fd602591396001600061131761191f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cd99092919063ffffffff16565b611927565b6001905092915050565b60006040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f44697361626c656420666561747572650000000000000000000000000000000081525060200191505060405180910390fd5b60008033905061142e818686611d99565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905061148660018261189790919063ffffffff16565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f17058394898febb944da4757aec20238c33e9566d2f7b14b2289184e107741b98885896040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156115a0578082015181840152602081019050611585565b50505050905090810190601f1680156115cd5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a46001925050509392505050565b60006040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f44697361626c656420666561747572650000000000000000000000000000000081525060200191505060405180910390fd5b6116616111b9565b6116d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6116dc8161204f565b50565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611765576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806122936021913960400191505060405180910390fd5b6117d081604051806060016040528060228152602001612203602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cd99092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506118278160025461219590919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600080828401905083811015611915576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156119ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806122d96024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611a33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061224b6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611bc1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611bd68160025461189790919063ffffffff16565b600281905550611c2d816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461189790919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000838311158290611d86576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611d4b578082015181840152602081019050611d30565b50505050905090810190601f168015611d785780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611e1f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806122b46025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611ea5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806121e06023913960400191505060405180910390fd5b611f108160405180606001604052806026815260200161226d602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cd99092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611fa3816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461189790919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156120d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806122256026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006121d783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611cd9565b90509291505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a723158202027bb343024b2cd8a06a8d1c3161eb39ece15b7c1cfa3241ca7a887743a065064736f6c634300050f0032"

// DeploySidechainERC20 deploys a new Ethereum contract, binding an instance of SidechainERC20 to it.
func DeploySidechainERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _mainchainToken common.Address, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *SidechainERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SidechainERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SidechainERC20Bin), backend, _mainchainToken, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SidechainERC20{SidechainERC20Caller: SidechainERC20Caller{contract: contract}, SidechainERC20Transactor: SidechainERC20Transactor{contract: contract}, SidechainERC20Filterer: SidechainERC20Filterer{contract: contract}}, nil
}

// SidechainERC20 is an auto generated Go binding around an Ethereum contract.
type SidechainERC20 struct {
	SidechainERC20Caller     // Read-only binding to the contract
	SidechainERC20Transactor // Write-only binding to the contract
	SidechainERC20Filterer   // Log filterer for contract events
}

// SidechainERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SidechainERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SidechainERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SidechainERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SidechainERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SidechainERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SidechainERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SidechainERC20Session struct {
	Contract     *SidechainERC20   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SidechainERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SidechainERC20CallerSession struct {
	Contract *SidechainERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SidechainERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SidechainERC20TransactorSession struct {
	Contract     *SidechainERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SidechainERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SidechainERC20Raw struct {
	Contract *SidechainERC20 // Generic contract binding to access the raw methods on
}

// SidechainERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SidechainERC20CallerRaw struct {
	Contract *SidechainERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SidechainERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SidechainERC20TransactorRaw struct {
	Contract *SidechainERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSidechainERC20 creates a new instance of SidechainERC20, bound to a specific deployed contract.
func NewSidechainERC20(address common.Address, backend bind.ContractBackend) (*SidechainERC20, error) {
	contract, err := bindSidechainERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20{SidechainERC20Caller: SidechainERC20Caller{contract: contract}, SidechainERC20Transactor: SidechainERC20Transactor{contract: contract}, SidechainERC20Filterer: SidechainERC20Filterer{contract: contract}}, nil
}

// NewSidechainERC20Caller creates a new read-only instance of SidechainERC20, bound to a specific deployed contract.
func NewSidechainERC20Caller(address common.Address, caller bind.ContractCaller) (*SidechainERC20Caller, error) {
	contract, err := bindSidechainERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20Caller{contract: contract}, nil
}

// NewSidechainERC20Transactor creates a new write-only instance of SidechainERC20, bound to a specific deployed contract.
func NewSidechainERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SidechainERC20Transactor, error) {
	contract, err := bindSidechainERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20Transactor{contract: contract}, nil
}

// NewSidechainERC20Filterer creates a new log filterer instance of SidechainERC20, bound to a specific deployed contract.
func NewSidechainERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SidechainERC20Filterer, error) {
	contract, err := bindSidechainERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20Filterer{contract: contract}, nil
}

// bindSidechainERC20 binds a generic wrapper to an already deployed contract.
func bindSidechainERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SidechainERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SidechainERC20 *SidechainERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SidechainERC20.Contract.SidechainERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SidechainERC20 *SidechainERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SidechainERC20.Contract.SidechainERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SidechainERC20 *SidechainERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SidechainERC20.Contract.SidechainERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SidechainERC20 *SidechainERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SidechainERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SidechainERC20 *SidechainERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SidechainERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SidechainERC20 *SidechainERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SidechainERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SidechainERC20.Contract.Allowance(&_SidechainERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SidechainERC20.Contract.Allowance(&_SidechainERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _SidechainERC20.Contract.BalanceOf(&_SidechainERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SidechainERC20.Contract.BalanceOf(&_SidechainERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SidechainERC20 *SidechainERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SidechainERC20 *SidechainERC20Session) Decimals() (uint8, error) {
	return _SidechainERC20.Contract.Decimals(&_SidechainERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SidechainERC20 *SidechainERC20CallerSession) Decimals() (uint8, error) {
	return _SidechainERC20.Contract.Decimals(&_SidechainERC20.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SidechainERC20 *SidechainERC20Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SidechainERC20 *SidechainERC20Session) IsOwner() (bool, error) {
	return _SidechainERC20.Contract.IsOwner(&_SidechainERC20.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SidechainERC20 *SidechainERC20CallerSession) IsOwner() (bool, error) {
	return _SidechainERC20.Contract.IsOwner(&_SidechainERC20.CallOpts)
}

// MainchainToken is a free data retrieval call binding the contract method 0x9cdcac9c.
//
// Solidity: function mainchainToken() constant returns(address)
func (_SidechainERC20 *SidechainERC20Caller) MainchainToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "mainchainToken")
	return *ret0, err
}

// MainchainToken is a free data retrieval call binding the contract method 0x9cdcac9c.
//
// Solidity: function mainchainToken() constant returns(address)
func (_SidechainERC20 *SidechainERC20Session) MainchainToken() (common.Address, error) {
	return _SidechainERC20.Contract.MainchainToken(&_SidechainERC20.CallOpts)
}

// MainchainToken is a free data retrieval call binding the contract method 0x9cdcac9c.
//
// Solidity: function mainchainToken() constant returns(address)
func (_SidechainERC20 *SidechainERC20CallerSession) MainchainToken() (common.Address, error) {
	return _SidechainERC20.Contract.MainchainToken(&_SidechainERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SidechainERC20 *SidechainERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SidechainERC20 *SidechainERC20Session) Name() (string, error) {
	return _SidechainERC20.Contract.Name(&_SidechainERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SidechainERC20 *SidechainERC20CallerSession) Name() (string, error) {
	return _SidechainERC20.Contract.Name(&_SidechainERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SidechainERC20.Contract.Nonces(&_SidechainERC20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) constant returns(uint256)
func (_SidechainERC20 *SidechainERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SidechainERC20.Contract.Nonces(&_SidechainERC20.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SidechainERC20 *SidechainERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SidechainERC20 *SidechainERC20Session) Owner() (common.Address, error) {
	return _SidechainERC20.Contract.Owner(&_SidechainERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SidechainERC20 *SidechainERC20CallerSession) Owner() (common.Address, error) {
	return _SidechainERC20.Contract.Owner(&_SidechainERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SidechainERC20 *SidechainERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SidechainERC20 *SidechainERC20Session) Symbol() (string, error) {
	return _SidechainERC20.Contract.Symbol(&_SidechainERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SidechainERC20 *SidechainERC20CallerSession) Symbol() (string, error) {
	return _SidechainERC20.Contract.Symbol(&_SidechainERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SidechainERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SidechainERC20 *SidechainERC20Session) TotalSupply() (*big.Int, error) {
	return _SidechainERC20.Contract.TotalSupply(&_SidechainERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SidechainERC20 *SidechainERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _SidechainERC20.Contract.TotalSupply(&_SidechainERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20Transactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20Session) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Approve(&_SidechainERC20.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20TransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Approve(&_SidechainERC20.TransactOpts, arg0, arg1)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SidechainERC20 *SidechainERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SidechainERC20 *SidechainERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.DecreaseAllowance(&_SidechainERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SidechainERC20 *SidechainERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.DecreaseAllowance(&_SidechainERC20.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x49bdc2b8.
//
// Solidity: function deposit(address account, uint256 amount, bytes signature) returns()
func (_SidechainERC20 *SidechainERC20Transactor) Deposit(opts *bind.TransactOpts, account common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "deposit", account, amount, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x49bdc2b8.
//
// Solidity: function deposit(address account, uint256 amount, bytes signature) returns()
func (_SidechainERC20 *SidechainERC20Session) Deposit(account common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Deposit(&_SidechainERC20.TransactOpts, account, amount, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x49bdc2b8.
//
// Solidity: function deposit(address account, uint256 amount, bytes signature) returns()
func (_SidechainERC20 *SidechainERC20TransactorSession) Deposit(account common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Deposit(&_SidechainERC20.TransactOpts, account, amount, signature)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SidechainERC20 *SidechainERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SidechainERC20 *SidechainERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.IncreaseAllowance(&_SidechainERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SidechainERC20 *SidechainERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.IncreaseAllowance(&_SidechainERC20.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SidechainERC20 *SidechainERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SidechainERC20 *SidechainERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _SidechainERC20.Contract.RenounceOwnership(&_SidechainERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SidechainERC20 *SidechainERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SidechainERC20.Contract.RenounceOwnership(&_SidechainERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20Transactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20Session) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Transfer(&_SidechainERC20.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20TransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Transfer(&_SidechainERC20.TransactOpts, arg0, arg1)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address recipient, uint256 amount, bytes signature) returns(bool)
func (_SidechainERC20 *SidechainERC20Transactor) Transfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "transfer0", recipient, amount, signature)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address recipient, uint256 amount, bytes signature) returns(bool)
func (_SidechainERC20 *SidechainERC20Session) Transfer0(recipient common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Transfer0(&_SidechainERC20.TransactOpts, recipient, amount, signature)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address recipient, uint256 amount, bytes signature) returns(bool)
func (_SidechainERC20 *SidechainERC20TransactorSession) Transfer0(recipient common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Transfer0(&_SidechainERC20.TransactOpts, recipient, amount, signature)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20Transactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20Session) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.TransferFrom(&_SidechainERC20.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_SidechainERC20 *SidechainERC20TransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SidechainERC20.Contract.TransferFrom(&_SidechainERC20.TransactOpts, arg0, arg1, arg2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SidechainERC20 *SidechainERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SidechainERC20 *SidechainERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SidechainERC20.Contract.TransferOwnership(&_SidechainERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SidechainERC20 *SidechainERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SidechainERC20.Contract.TransferOwnership(&_SidechainERC20.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 amount, bytes signature) returns()
func (_SidechainERC20 *SidechainERC20Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.contract.Transact(opts, "withdraw", amount, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 amount, bytes signature) returns()
func (_SidechainERC20 *SidechainERC20Session) Withdraw(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Withdraw(&_SidechainERC20.TransactOpts, amount, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 amount, bytes signature) returns()
func (_SidechainERC20 *SidechainERC20TransactorSession) Withdraw(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SidechainERC20.Contract.Withdraw(&_SidechainERC20.TransactOpts, amount, signature)
}

// SidechainERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SidechainERC20 contract.
type SidechainERC20ApprovalIterator struct {
	Event *SidechainERC20Approval // Event containing the contract specifics and raw log

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
func (it *SidechainERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SidechainERC20Approval)
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
		it.Event = new(SidechainERC20Approval)
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
func (it *SidechainERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SidechainERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SidechainERC20Approval represents a Approval event raised by the SidechainERC20 contract.
type SidechainERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SidechainERC20 *SidechainERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SidechainERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SidechainERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20ApprovalIterator{contract: _SidechainERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SidechainERC20 *SidechainERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SidechainERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SidechainERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SidechainERC20Approval)
				if err := _SidechainERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SidechainERC20 *SidechainERC20Filterer) ParseApproval(log types.Log) (*SidechainERC20Approval, error) {
	event := new(SidechainERC20Approval)
	if err := _SidechainERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SidechainERC20DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SidechainERC20 contract.
type SidechainERC20DepositIterator struct {
	Event *SidechainERC20Deposit // Event containing the contract specifics and raw log

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
func (it *SidechainERC20DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SidechainERC20Deposit)
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
		it.Event = new(SidechainERC20Deposit)
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
func (it *SidechainERC20DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SidechainERC20DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SidechainERC20Deposit represents a Deposit event raised by the SidechainERC20 contract.
type SidechainERC20Deposit struct {
	MainchainToken common.Address
	Account        common.Address
	Amount         *big.Int
	Nonce          *big.Int
	Signature      []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x99d353966fe3c3445ee3f115b9ce5bc579b0e19a76b2484f79ae9d42e45fe4ce.
//
// Solidity: event Deposit(address indexed mainchainToken, address indexed account, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) FilterDeposit(opts *bind.FilterOpts, mainchainToken []common.Address, account []common.Address) (*SidechainERC20DepositIterator, error) {

	var mainchainTokenRule []interface{}
	for _, mainchainTokenItem := range mainchainToken {
		mainchainTokenRule = append(mainchainTokenRule, mainchainTokenItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SidechainERC20.contract.FilterLogs(opts, "Deposit", mainchainTokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20DepositIterator{contract: _SidechainERC20.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x99d353966fe3c3445ee3f115b9ce5bc579b0e19a76b2484f79ae9d42e45fe4ce.
//
// Solidity: event Deposit(address indexed mainchainToken, address indexed account, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SidechainERC20Deposit, mainchainToken []common.Address, account []common.Address) (event.Subscription, error) {

	var mainchainTokenRule []interface{}
	for _, mainchainTokenItem := range mainchainToken {
		mainchainTokenRule = append(mainchainTokenRule, mainchainTokenItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SidechainERC20.contract.WatchLogs(opts, "Deposit", mainchainTokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SidechainERC20Deposit)
				if err := _SidechainERC20.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x99d353966fe3c3445ee3f115b9ce5bc579b0e19a76b2484f79ae9d42e45fe4ce.
//
// Solidity: event Deposit(address indexed mainchainToken, address indexed account, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) ParseDeposit(log types.Log) (*SidechainERC20Deposit, error) {
	event := new(SidechainERC20Deposit)
	if err := _SidechainERC20.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SidechainERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SidechainERC20 contract.
type SidechainERC20OwnershipTransferredIterator struct {
	Event *SidechainERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SidechainERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SidechainERC20OwnershipTransferred)
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
		it.Event = new(SidechainERC20OwnershipTransferred)
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
func (it *SidechainERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SidechainERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SidechainERC20OwnershipTransferred represents a OwnershipTransferred event raised by the SidechainERC20 contract.
type SidechainERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SidechainERC20 *SidechainERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SidechainERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SidechainERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20OwnershipTransferredIterator{contract: _SidechainERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SidechainERC20 *SidechainERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SidechainERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SidechainERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SidechainERC20OwnershipTransferred)
				if err := _SidechainERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SidechainERC20 *SidechainERC20Filterer) ParseOwnershipTransferred(log types.Log) (*SidechainERC20OwnershipTransferred, error) {
	event := new(SidechainERC20OwnershipTransferred)
	if err := _SidechainERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SidechainERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SidechainERC20 contract.
type SidechainERC20TransferIterator struct {
	Event *SidechainERC20Transfer // Event containing the contract specifics and raw log

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
func (it *SidechainERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SidechainERC20Transfer)
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
		it.Event = new(SidechainERC20Transfer)
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
func (it *SidechainERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SidechainERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SidechainERC20Transfer represents a Transfer event raised by the SidechainERC20 contract.
type SidechainERC20Transfer struct {
	MainchainToken common.Address
	Sender         common.Address
	Recipient      common.Address
	Amount         *big.Int
	Nonce          *big.Int
	Signature      []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x17058394898febb944da4757aec20238c33e9566d2f7b14b2289184e107741b9.
//
// Solidity: event Transfer(address indexed mainchainToken, address indexed sender, address indexed recipient, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) FilterTransfer(opts *bind.FilterOpts, mainchainToken []common.Address, sender []common.Address, recipient []common.Address) (*SidechainERC20TransferIterator, error) {

	var mainchainTokenRule []interface{}
	for _, mainchainTokenItem := range mainchainToken {
		mainchainTokenRule = append(mainchainTokenRule, mainchainTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SidechainERC20.contract.FilterLogs(opts, "Transfer", mainchainTokenRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20TransferIterator{contract: _SidechainERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x17058394898febb944da4757aec20238c33e9566d2f7b14b2289184e107741b9.
//
// Solidity: event Transfer(address indexed mainchainToken, address indexed sender, address indexed recipient, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SidechainERC20Transfer, mainchainToken []common.Address, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var mainchainTokenRule []interface{}
	for _, mainchainTokenItem := range mainchainToken {
		mainchainTokenRule = append(mainchainTokenRule, mainchainTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SidechainERC20.contract.WatchLogs(opts, "Transfer", mainchainTokenRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SidechainERC20Transfer)
				if err := _SidechainERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x17058394898febb944da4757aec20238c33e9566d2f7b14b2289184e107741b9.
//
// Solidity: event Transfer(address indexed mainchainToken, address indexed sender, address indexed recipient, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) ParseTransfer(log types.Log) (*SidechainERC20Transfer, error) {
	event := new(SidechainERC20Transfer)
	if err := _SidechainERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SidechainERC20Transfer0Iterator is returned from FilterTransfer0 and is used to iterate over the raw logs and unpacked data for Transfer0 events raised by the SidechainERC20 contract.
type SidechainERC20Transfer0Iterator struct {
	Event *SidechainERC20Transfer0 // Event containing the contract specifics and raw log

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
func (it *SidechainERC20Transfer0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SidechainERC20Transfer0)
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
		it.Event = new(SidechainERC20Transfer0)
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
func (it *SidechainERC20Transfer0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SidechainERC20Transfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SidechainERC20Transfer0 represents a Transfer0 event raised by the SidechainERC20 contract.
type SidechainERC20Transfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer0 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SidechainERC20 *SidechainERC20Filterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SidechainERC20Transfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SidechainERC20.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20Transfer0Iterator{contract: _SidechainERC20.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

// WatchTransfer0 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SidechainERC20 *SidechainERC20Filterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *SidechainERC20Transfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SidechainERC20.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SidechainERC20Transfer0)
				if err := _SidechainERC20.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

// ParseTransfer0 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SidechainERC20 *SidechainERC20Filterer) ParseTransfer0(log types.Log) (*SidechainERC20Transfer0, error) {
	event := new(SidechainERC20Transfer0)
	if err := _SidechainERC20.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SidechainERC20WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SidechainERC20 contract.
type SidechainERC20WithdrawIterator struct {
	Event *SidechainERC20Withdraw // Event containing the contract specifics and raw log

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
func (it *SidechainERC20WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SidechainERC20Withdraw)
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
		it.Event = new(SidechainERC20Withdraw)
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
func (it *SidechainERC20WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SidechainERC20WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SidechainERC20Withdraw represents a Withdraw event raised by the SidechainERC20 contract.
type SidechainERC20Withdraw struct {
	MainchainToken common.Address
	Account        common.Address
	Amount         *big.Int
	Nonce          *big.Int
	Signature      []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x33be7eabd8ed368ca1aa14ce2ad1e90a0c9bf21edbb3820d5591546e4eb84157.
//
// Solidity: event Withdraw(address indexed mainchainToken, address indexed account, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) FilterWithdraw(opts *bind.FilterOpts, mainchainToken []common.Address, account []common.Address) (*SidechainERC20WithdrawIterator, error) {

	var mainchainTokenRule []interface{}
	for _, mainchainTokenItem := range mainchainToken {
		mainchainTokenRule = append(mainchainTokenRule, mainchainTokenItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SidechainERC20.contract.FilterLogs(opts, "Withdraw", mainchainTokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SidechainERC20WithdrawIterator{contract: _SidechainERC20.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x33be7eabd8ed368ca1aa14ce2ad1e90a0c9bf21edbb3820d5591546e4eb84157.
//
// Solidity: event Withdraw(address indexed mainchainToken, address indexed account, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SidechainERC20Withdraw, mainchainToken []common.Address, account []common.Address) (event.Subscription, error) {

	var mainchainTokenRule []interface{}
	for _, mainchainTokenItem := range mainchainToken {
		mainchainTokenRule = append(mainchainTokenRule, mainchainTokenItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SidechainERC20.contract.WatchLogs(opts, "Withdraw", mainchainTokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SidechainERC20Withdraw)
				if err := _SidechainERC20.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x33be7eabd8ed368ca1aa14ce2ad1e90a0c9bf21edbb3820d5591546e4eb84157.
//
// Solidity: event Withdraw(address indexed mainchainToken, address indexed account, uint256 amount, uint256 nonce, bytes signature)
func (_SidechainERC20 *SidechainERC20Filterer) ParseWithdraw(log types.Log) (*SidechainERC20Withdraw, error) {
	event := new(SidechainERC20Withdraw)
	if err := _SidechainERC20.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
