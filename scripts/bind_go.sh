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
extract_abi_bin RollupTokenRegistry
extract_abi_bin RollupChain
extract_abi_bin RollupMerkleUtils
extract_abi_bin SidechainERC20
extract_abi_bin TokenMapper

rm -rf bindings/go/mainchain/rollup
mkdir -p bindings/go/mainchain/rollup
run_abigen RollupTokenRegistry rollup bindings/go/mainchain/rollup/rollup_token_registry
run_abigen RollupChain rollup bindings/go/mainchain/rollup/rollup_chain
run_abigen RollupMerkleUtils rollup bindings/go/mainchain/rollup/rollup_merkle_utils

rm -rf bindings/go/sidechain
mkdir -p bindings/go/sidechain
run_abigen SidechainERC20 sidechain bindings/go/sidechain/sidechain_erc20
run_abigen TokenMapper sidechain bindings/go/sidechain/token_mapper
