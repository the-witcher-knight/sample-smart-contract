package controller

import (
	"context"
)

func (ctrl controller) Retrieve(ctx context.Context) (int64, error) {
	val, err := ctrl.storeContract.Retrieve(nil)
	if err != nil {
		return 0, err
	}

	return val.Int64(), nil
}
