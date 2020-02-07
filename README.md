# Celer Optimistic Rollup Contracts

## Mainchain Contracts

Optimistic Rollup for deposit, withdraw and transfer of registered and mapped
ERC-20 tokens. The rollup chain tracks a users balances and nonces for a list of
tokens. An aggregator is expected to periodically submit the rollup blocks to
the contract.

## Sidechain Contracts

Modified ERC-20 contract to emit necessary data for rollup. In particular,
we augmented the transfer() function to include signature as an extra parameter
for rollup fraud proofs.
