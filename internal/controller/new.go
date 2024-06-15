package controller

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/the-witcher-knight/store-contract/internal/contracts"
)

type Controller interface {
	Store(ctx context.Context, value int64) error

	Retrieve(ctx context.Context) (int64, error)
}

type controller struct {
	storeContract contracts.Storage
	auth          bind.TransactOpts
	client        ethclient.Client
}

func New(storeContract contracts.Storage, auth bind.TransactOpts, client ethclient.Client) Controller {
	return &controller{
		storeContract: storeContract,
		auth:          auth,
		client:        client,
	}
}
