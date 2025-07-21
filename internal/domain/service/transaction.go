// Package service определяет бизнес-логику приложения.
// Содержит интерфейсы сервисного слоя, абстрагирующие бизнес-процессы.
package service

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/presentation/api/dto"
)

// TransactionService определяет контракт сервисного слоя для работы с транзакциями.
// Предоставляет бизнес-логику для операций с историей транзакций.
type TransactionService interface {
	LastNTransactions(ctx context.Context, n int) ([]dto.TransactionResponse, error)
}
