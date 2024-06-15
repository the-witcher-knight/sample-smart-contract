# Storage Contract

> Sample smart contract integration written in Go

## Prerequisite

- Setup `docker` on you local machine.
- Install `make` to use `Makefile`.

## How to run

1. Run `make setup`.
2. Run `docker logs ganache` to view the accounts and private key hashes.
3. Copy one of the keypair accounts and private keys, and add them to `.env.local`.
4. Run `make deploy-contract`.
5. You will see the contract address. Copy it and add it to `.env.local`.
6. Run `make server`.

> Note: Please remove the prefix '0x' from every configuration before adding them to `.env.local`.

> If you don't use macOS, you can either purchase one or refer to 
> the Makefile to find the commands used in each step. 
> Then, type those commands into your command line interface (`cmd`).

## How to contribute

### To update contracts

- Modify smart contract at `data/contracts`.
- Run `make solc` to generate `abi` and `bin` files in `build/contracts`.
- Run `make abigen` to generate go package, this action will update `internal/contracts` package.

> This application use [abigen](https://geth.ethereum.org/docs/tools/abigen) to generate Go from `abi` and `bin` file.

### To update APIs

- That quite easy...

## References

- https://geth.ethereum.org/docs/developers
- https://archive.trufflesuite.com/ganache/
