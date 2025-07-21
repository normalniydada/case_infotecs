// Package transaction предоставляет сервисный слой для работы с транзакциями.
// Реализует бизнес-логику обработки транзакций между кошельками.
package transaction

import (
	"context"
	"fmt"
	er "github.com/normalniydada/case_infotecs/internal/domain/errors"
	"github.com/normalniydada/case_infotecs/internal/domain/repository"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/dto"
)

// transactionService реализует интерфейс TransactionService.
// Содержит репозиторий для работы с данными транзакций.
type transactionService struct {
	transactionRepo repository.TransactionRepository
}

// NewTransactionService создает новый экземпляр сервиса для работы с транзакциями.
//
// Параметры:
//   - transactionRepo: репозиторий для доступа к данным транзакций
//
// Возвращает:
//   - service.TransactionService: реализацию интерфейса сервиса транзакций
func NewTransactionService(transactionRepo repository.TransactionRepository) service.TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

// LastNTransactions возвращает последние N транзакций из системы.
// Если транзакции не найдены, возвращает ErrTransactionNotFound.
//
// Параметры:
//   - ctx: контекст выполнения
//   - n: количество запрашиваемых транзакций
//
// Возвращает:
//   - []dto.TransactionResponse: список транзакций в формате DTO
//   - error: ошибка, если не удалось получить транзакции
//
// Возможные ошибки:
//   - ErrTransactionNotFound: если транзакции не найдены
//   - Другие ошибки репозитория: при проблемах доступа к данным
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
