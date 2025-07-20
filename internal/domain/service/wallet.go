package service

import (
	"context"
	"github.com/shopspring/decimal"
)

type WalletService interface {
	Balance(ctx context.Context, address string) (decimal.Decimal, error)
	TransferMoney(ctx context.Context, from, to string, amount decimal.Decimal) error
	CreateWallet(ctx context.Context, balance decimal.Decimal) error
	CountWallets(ctx context.Context) (int64, error)
}
