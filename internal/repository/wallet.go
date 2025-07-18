package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	Wallet(ctx context.Context, address uuid.UUID) (*models.Wallet, error)
	Transfer(ctx context.Context, from, to uuid.UUID, amount int64) error
	// WalletExists(ctx context.Context, address uuid.UUID) (bool, error)
}
