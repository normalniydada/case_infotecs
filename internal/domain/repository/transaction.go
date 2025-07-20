// Package repository определяет интерфейсы для работы с хранилищами данных.
// Содержит контракты, которые должны реализовывать репозитории приложения.

package repository

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
)

// TransactionRepository определяет контракт для работы с хранилищем транзакций.
// Описывает методы доступа к данным транзакций.
type TransactionRepository interface {
	LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error)
}
