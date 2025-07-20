package repositories

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/normalniydada/case_infotecs/internal/domain/repository"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(n).
		Find(&transactions).Error

	return transactions, err
}
