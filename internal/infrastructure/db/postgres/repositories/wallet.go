// Package repositories содержит реализации репозиториев для работы с хранилищами данных.
// Включает конкретные реализации интерфейсов доменного слоя.
package repositories

import (
	"context"
	"errors"
	"fmt"
	er "github.com/normalniydada/case_infotecs/internal/domain/errors"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/normalniydada/case_infotecs/internal/domain/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// walletRepository реализует интерфейс WalletRepository для работы с кошельками в PostgreSQL.
// Обеспечивает безопасное выполнение операций с блокировками и транзакциями.
type walletRepository struct {
	db *gorm.DB // Экземпляр GORM для работы с БД
}

// NewWalletRepository создает новый экземпляр репозитория кошельков.
//
// Параметры:
//   - db: подключение к БД (*gorm.DB)
//
// Возвращает:
//   - repository.WalletRepository: реализацию интерфейса репозитория
func NewWalletRepository(db *gorm.DB) repository.WalletRepository {
	return &walletRepository{db: db}
}

// CreateWallet создает новый кошелек в базе данных.
// Выполняется в транзакции с проверкой уникальности адреса.
//
// Параметры:
//   - ctx: контекст выполнения
//   - wallet: указатель на создаваемый кошелек
//
// Возвращает:
//   - error: ошибка при создании:
//   - er.ErrWalletExists: если кошелек с таким адресом уже существует
//   - другие ошибки базы данных
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

// Wallet возвращает кошелек по его адресу.
//
// Параметры:
//   - ctx: контекст выполнения
//   - address: адрес кошелька
//
// Возвращает:
//   - *models.Wallet: найденный кошелек
//   - error: ошибка при поиске:
//   - er.ErrWalletNotFound: если кошелек не существует
//   - другие ошибки базы данных
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

// Transfer выполняет перевод средств между кошельками.
// Операция выполняется атомарно в транзакции.
//
// Параметры:
//   - ctx: контекст выполнения
//   - from: адрес кошелька-отправителя
//   - to: адрес кошелька-получателя
//   - amount: сумма перевода
//
// Возвращает:
//   - error: ошибка при переводе:
//   - er.ErrWalletSenderNotFound: отправитель не найден
//   - er.ErrWalletReceiverNotFound: получатель не найден
//   - er.ErrNotEnoughMoney: недостаточно средств
//   - другие ошибки базы данных
func (r *walletRepository) Transfer(ctx context.Context, from, to string, amount decimal.Decimal) error {
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

// lockAndValidateWallets блокирует и проверяет кошельки для перевода.
// Внутренний метод, используется в Transfer.
func (r *walletRepository) lockAndValidateWallets(tx *gorm.DB, from, to string, amount decimal.Decimal) (*models.Wallet,
	*models.Wallet, error) {
	var sender, receiver models.Wallet

	if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		First(&sender, "address = ?", from).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, er.ErrWalletSenderNotFound
		}
		return nil, nil, fmt.Errorf("error blocking sender's wallet: %w", err)
	}

	if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		First(&receiver, "address = ?", to).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, er.ErrWalletReceiverNotFound
		}
		return nil, nil, fmt.Errorf("error blocking receiver's wallet: %w", err)
	}

	if sender.Balance.LessThan(amount) {
		return nil, nil, er.ErrNotEnoughMoney
	}

	return &sender, &receiver, nil
}

// updateBalance обновляет балансы кошельков после перевода.
// Внутренний метод, используется в Transfer.
func (r *walletRepository) updateBalance(tx *gorm.DB, sender, receiver *models.Wallet, amount decimal.Decimal) error {
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

// createTransaction создает запись о транзакции.
// Внутренний метод, используется в Transfer.
func (r *walletRepository) createTransaction(tx *gorm.DB, from, to string, amount decimal.Decimal) error {
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

// Count возвращает общее количество кошельков в системе.
//
// Параметры:
//   - ctx: контекст выполнения
//
// Возвращает:
//   - int64: количество кошельков
//   - error: ошибка при подсчете
func (r *walletRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Wallet{}).Count(&count).Error
	return count, err
}
