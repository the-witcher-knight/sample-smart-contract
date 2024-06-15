package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/the-witcher-knight/store-contract/internal/config"
	"github.com/the-witcher-knight/store-contract/internal/contracts"
	"github.com/the-witcher-knight/store-contract/internal/system"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg := config.ReadConfigFromEnv()

	s, err := system.New(cfg)
	if err != nil {
		return err
	}

	chainID, err := s.EthClient().ChainID(ctx)
	if err != nil {
		return err
	}

	auth, err := authByPrivateKey(cfg, *chainID)
	if err != nil {
		return err
	}

	addr, tx, _, err := contracts.DeployStorage(&auth, s.EthClient())
	if err != nil {
		return err
	}

	fmt.Printf("Contract address: %s\n", addr.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
	fmt.Printf("Transaction nonce: %d\n", tx.Nonce())
	fmt.Printf("Transaction cost: %d\n", tx.Cost())
	fmt.Printf("Transaction timestamp: %s\n", tx.Time().Format(time.RFC3339))

	return nil
}

func authByPrivateKey(cfg config.Config, chainID big.Int) (bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(cfg.SecretKey)
	if err != nil {
		return bind.TransactOpts{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, &chainID)
	if err != nil {
		log.Fatalf("Failed to create keyed transactor: %v", err)
	}

	return *auth, nil
}
