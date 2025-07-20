package repository

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/shopspring/decimal"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	Wallet(ctx context.Context, address string) (*models.Wallet, error)
	Transfer(ctx context.Context, from, to string, amount decimal.Decimal) error
	Count(ctx context.Context) (int64, error)
}
