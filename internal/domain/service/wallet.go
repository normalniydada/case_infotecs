// Package service определяет бизнес-логику приложения.
// Содержит интерфейсы сервисного слоя, абстрагирующие бизнес-процессы.

package service

import (
	"context"
	"github.com/shopspring/decimal"
)

// WalletService определяет контракт сервисного слоя для работы с кошельками.
// Предоставляет бизнес-логику для управления кошельками и переводами средств.
// Все методы должны быть безопасны для конкурентного вызова.
type WalletService interface {
	Balance(ctx context.Context, address string) (decimal.Decimal, error)
	TransferMoney(ctx context.Context, from, to string, amount decimal.Decimal) error
	CreateWallet(ctx context.Context, balance decimal.Decimal) error
	CountWallets(ctx context.Context) (int64, error)
}
