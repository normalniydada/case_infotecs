package service

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/dto"
)

type TransactionService interface {
	LastNTransactions(ctx context.Context, n int) ([]dto.TransactionResponse, error)
}
