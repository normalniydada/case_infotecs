// Package repository определяет интерфейсы для работы с хранилищами данных.
// Содержит контракты, которые должны реализовывать репозитории приложения.

package repository

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/shopspring/decimal"
)

// WalletRepository определяет контракт для работы с хранилищем кошельков.
// Описывает методы для управления кошельками и операциями перевода средств.
type WalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	Wallet(ctx context.Context, address string) (*models.Wallet, error)
	Transfer(ctx context.Context, from, to string, amount decimal.Decimal) error
	Count(ctx context.Context) (int64, error)
}
