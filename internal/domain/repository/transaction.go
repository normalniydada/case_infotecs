package repository

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
)

type TransactionRepository interface {
	LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error)
}
