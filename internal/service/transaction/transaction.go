package transaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/normalniydada/case_infotecs/internal/repository"
)

var (
	ErrTransactionNotFound = errors.New("транзакции отсутствуют")
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s *transactionService) LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error) {
	transactions, err := s.transactionRepo.LastNTransactions(ctx, n)

	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка транзакций: %w", err)
	}

	if len(transactions) == 0 {
		return nil, ErrTransactionNotFound
	}

	return transactions, nil
}
