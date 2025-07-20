package repository

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
)

// TransactionRepository определяет контракт для работы с хранилищем транзакций.
// Описывает методы доступа к данным транзакций между кошельками.
type TransactionRepository interface {
	LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error)
}
