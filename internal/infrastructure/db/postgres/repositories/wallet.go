package repositories

import (
	"context"
	"errors"
	"fmt"
	er "github.com/normalniydada/case_infotecs/internal/domain/errors"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/normalniydada/case_infotecs/internal/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) repository.WalletRepository {
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
			return er.ErrWalletExists
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("error checking wallet existence: %w", err)
		}

		if err = tx.Create(wallet).Error; err != nil {
			return fmt.Errorf("error creating wallet: %w", err)
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, er.ErrWalletExists) {
			return er.ErrWalletExists
		}
		return fmt.Errorf("error creating wallet: %w", err)
	}

	return nil
}

func (r *walletRepository) Wallet(ctx context.Context, address string) (*models.Wallet, error) {
	var wallet models.Wallet

	err := r.db.WithContext(ctx).First(&wallet, "address = ?", address).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, er.ErrWalletNotFound
		}
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepository) Transfer(ctx context.Context, from, to string, amount int64) error {
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

func (r *walletRepository) lockAndValidateWallets(tx *gorm.DB, from, to string, amount int64) (*models.Wallet,
	*models.Wallet, error) {
	var wallets []models.Wallet

	if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		Where("address IN (?, ?)", from, to).
		Find(&wallets).Error; err != nil {
		return nil, nil, fmt.Errorf("error blocking wallets: %w", err)
	}

	if len(wallets) != 2 {
		return nil, nil, er.ErrWalletNotFound
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
		return nil, nil, er.ErrNotEnoughMoney
	}

	return sender, receiver, nil
}

func (r *walletRepository) updateBalance(tx *gorm.DB, sender, receiver *models.Wallet, amount int64) error {
	if err := tx.Model(sender).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		return fmt.Errorf("error while writing off funds: %w", err)
	}

	if err := tx.Model(receiver).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return fmt.Errorf("error while crediting funds: %w", err)
	}

	return nil
}

func (r *walletRepository) createTransaction(tx *gorm.DB, from, to string, amount int64) error {
	transaction := models.Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		return fmt.Errorf("error creating transaction: %w", err)
	}

	return nil
}

func (r *walletRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Wallet{}).Count(&count).Error
	return count, err
}
