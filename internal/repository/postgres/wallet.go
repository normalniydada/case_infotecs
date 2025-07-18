package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrWalletNotFound = errors.New("кошелек не найден")
	ErrWalletExists   = errors.New("кошелек уже существует")
	ErrNotEnoughFunds = errors.New("недостаточно средств на кошельке отправителя")
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingWallet models.Wallet

		err := tx.Clauses(clause.Locking{
			Strength: clause.LockingStrengthUpdate,
			Options:  clause.LockingOptionsNoWait,
		}).Where("address = ?", wallet.Address).
			First(&existingWallet).Error

		if err == nil {
			return ErrWalletExists
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ошибка при проверки существования кошелька: %w", err)
		}

		if err = tx.Create(wallet).Error; err != nil {
			return fmt.Errorf("ошибка при создании кошелька: %w", err)
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, ErrWalletExists) {
			return ErrWalletExists
		}
		return fmt.Errorf("ошибка при создании кошелька: %w", err)
	}

	return nil
}

func (r *walletRepository) Wallet(ctx context.Context, address uuid.UUID) (*models.Wallet, error) {
	var wallet models.Wallet

	err := r.db.WithContext(ctx).First(&wallet, "address = ?", address).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWalletNotFound
		}
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepository) Transfer(ctx context.Context, from, to uuid.UUID, amount int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sender, receiver, err := r.lockAndValidateWallets(tx, from, to, amount)
		if err != nil {
			return err
		}

		if err = r.updateBalance(tx, sender, receiver, amount); err != nil {
			return err
		}

		return r.createTransaction(tx, from, to, amount)
	})
}

func (r *walletRepository) lockAndValidateWallets(tx *gorm.DB, from, to uuid.UUID, amount int64) (*models.Wallet,
	*models.Wallet, error) {
	var wallets []models.Wallet

	if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		Where("address IN (?, ?)", from, to).
		Find(&wallets).Error; err != nil {
		return nil, nil, fmt.Errorf("ошибка при блокировки кошельков: %w", err)
	}

	if len(wallets) != 2 {
		return nil, nil, ErrWalletNotFound
	}

	var sender, receiver *models.Wallet
	for _, wallet := range wallets {
		if wallet.Address == from {
			sender = &wallet
		} else {
			receiver = &wallet
		}
	}

	if sender.Balance < amount {
		return nil, nil, ErrNotEnoughFunds
	}

	return sender, receiver, nil
}

func (r *walletRepository) updateBalance(tx *gorm.DB, sender, receiver *models.Wallet, amount int64) error {
	if err := tx.Model(sender).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		return fmt.Errorf("ошибка при списании средств: %w", err)
	}

	if err := tx.Model(receiver).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return fmt.Errorf("ошибка при зачислении средств: %w", err)
	}

	return nil
}

func (r *walletRepository) createTransaction(tx *gorm.DB, from, to uuid.UUID, amount int64) error {
	transaction := models.Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		return fmt.Errorf("ошибка при создании транзакции: %w", err)
	}

	return nil
}

/*func (r *WalletRepository) WalletExists(ctx context.Context, address uuid.UUID) (bool, error) {
	exists := false

	err := r.db.WithContext(ctx).
		Model(&models.Wallet{}).
		Select("count(*) > 0").
		Where("address = ?", address).
		Find(&exists).Error

	if err != nil {
		return false, fmt.Errorf("ошибка проверки существования кошелька: %w", err)
	}

	return exists, nil
}*/
