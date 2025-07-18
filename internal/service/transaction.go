package service

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
)

type TransactionService interface {
	LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error)
}
