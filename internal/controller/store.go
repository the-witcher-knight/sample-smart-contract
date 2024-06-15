package controller

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	pkgerrors "github.com/pkg/errors"
)

func (ctrl controller) Store(ctx context.Context, value int64) error {
	tx, err := ctrl.storeContract.Store(&ctrl.auth, big.NewInt(value))
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	_, err = bind.WaitMined(ctx, &ctrl.client, tx)
	if err != nil {
		return err
	}

	return nil
}
