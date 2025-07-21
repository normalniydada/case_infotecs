// Package repositories содержит реализации репозиториев для работы с хранилищами данных.
// Включает конкретные реализации интерфейсов доменного слоя.
package repositories

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/normalniydada/case_infotecs/internal/domain/repository"
	"gorm.io/gorm"
)

// transactionRepository реализует интерфейс TransactionRepository для работы с транзакциями в PostgreSQL.
// Использует GORM для взаимодействия с базой данных.
type transactionRepository struct {
	db *gorm.DB // Экземпляр GORM для работы с БД
}

// NewTransactionRepository создает новый экземпляр репозитория транзакций.
//
// Параметры:
//   - db: подключение к БД (*gorm.DB)
//
// Возвращает:
//   - repository.TransactionRepository: реализацию интерфейса репозитория
func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{db: db}
}

// LastNTransactions возвращает последние N транзакций из базы данных,
// отсортированные по дате создания (от новых к старым).
//
// Параметры:
//   - ctx: контекст выполнения (для отмены и таймаутов)
//   - n: количество возвращаемых транзакций
//
// Возвращает:
//   - []models.Transaction: список транзакций
//   - error: ошибка при выполнении запроса
//
// Возможные ошибки:
//   - gorm.ErrRecordNotFound: если транзакции не найдены
//   - context.DeadlineExceeded: при превышении таймаута
//   - другие ошибки базы данных
//
// Особенности:
//   - Использует контекст для отмены операций
//   - Сортировка по created_at DESC (новые сначала)
func (r *transactionRepository) LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(n).
		Find(&transactions).Error

	return transactions, err
}
