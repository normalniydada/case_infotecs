package service

import (
	"context"
	"github.com/google/uuid"
)

type WalletService interface {
	Balance(ctx context.Context, address uuid.UUID) (float64, error)
	TransferMoney(ctx context.Context, from, to uuid.UUID, amount float64) error
}
