#!/bin/sh

# Must be run from top-level directory

extract_abi_bin() {
  jq .abi artifacts/$1.json >genfiles/$1.abi
  jq -r .bytecode artifacts/$1.json >genfiles/$1.bin
}

run_abigen() {
  abigen -abi genfiles/$1.abi -bin genfiles/$1.bin -type $1 -pkg $2 -out $3.go
}

rm -rf genfiles
mkdir -p genfiles
extract_abi_bin AccountRegistry
extract_abi_bin DataTypes
extract_abi_bin DepositWithdrawManager
extract_abi_bin MerkleUtils
extract_abi_bin RollupChain
extract_abi_bin TokenRegistry
extract_abi_bin TransitionEvaluator
extract_abi_bin ValidatorRegistry

extract_abi_bin BlockCommittee
extract_abi_bin DummyApp
extract_abi_bin SidechainERC20
extract_abi_bin TokenMapper

rm -rf bindings/go/mainchain
mkdir -p bindings/go/mainchain
run_abigen AccountRegistry mainchain bindings/go/mainchain/account_registry
run_abigen DepositWithdrawManager mainchain bindings/go/mainchain/deposit_withdraw_manager
run_abigen MerkleUtils mainchain bindings/go/mainchain/merkle_utils
run_abigen RollupChain mainchain bindings/go/mainchain/rollup_chain
run_abigen TokenRegistry mainchain bindings/go/mainchain/token_registry
run_abigen TransitionEvaluator mainchain bindings/go/mainchain/transition_evaluator
run_abigen ValidatorRegistry mainchain bindings/go/mainchain/validator_registry

rm -rf bindings/go/sidechain
mkdir -p bindings/go/sidechain
run_abigen BlockCommittee sidechain bindings/go/sidechain/block_committee
run_abigen DummyApp sidechain bindings/go/sidechain/dummy_app
run_abigen SidechainERC20 sidechain bindings/go/sidechain/sidechain_erc20
run_abigen TokenMapper sidechain bindings/go/sidechain/token_mapper

# Hack until we figure out how to avoid duplicate declaration with abigen
sed -i '' -e '/^\/\/ DataTypesAccountInfo.*$/,/^}$/d' \
  -e '/^\/\/ DataTypesStorageSlot.*$/,/^}$/d' \
  bindings/go/mainchain/transition_evaluator.go
sed -i '' -e '/^\/\/ DataTypesAccountInfo.*$/,/^}$/d' \
  -e '/^\/\/ DataTypesIncludedStorageSlot.*$/,/^}$/d' \
  -e '/^\/\/ DataTypesIncludedTransition.*$/,/^}$/d' \
  -e '/^\/\/ DataTypesStorageSlot.*$/,/^}$/d' \
  -e '/^\/\/ DataTypesTransitionInclusionProof.*$/,/^}$/d' \
  bindings/go/mainchain/deposit_withdraw_manager.go
