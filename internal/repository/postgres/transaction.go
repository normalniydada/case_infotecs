package postgres

import (
	"context"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) LastNTransactions(ctx context.Context, n int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.WithContext(ctx).
		Order("created_at desc").
		Limit(n).
		Find(&transactions).Error

	return transactions, err
}
