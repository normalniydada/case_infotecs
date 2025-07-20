package transaction

import (
	"context"
	"fmt"
	er "github.com/normalniydada/case_infotecs/internal/domain/errors"
	"github.com/normalniydada/case_infotecs/internal/domain/repository"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/dto"
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) service.TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s *transactionService) LastNTransactions(ctx context.Context, n int) ([]dto.TransactionResponse, error) {
	transactions, err := s.transactionRepo.LastNTransactions(ctx, n)

	if err != nil {
		return nil, fmt.Errorf("error getting transaction list: %w", err)
	}

	if len(transactions) == 0 {
		return nil, er.ErrTransactionNotFound
	}

	respTransactions := make([]dto.TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		respTransactions = append(respTransactions, dto.TransactionResponse{
			From:      transaction.From,
			To:        transaction.To,
			Amount:    transaction.Amount,
			CreatedAt: transaction.CreatedAt,
		})
	}

	return respTransactions, nil
}
