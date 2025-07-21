// Package wallet предоставляет сервисный слой для работы с кошельками.
// Реализует бизнес-логику управления кошельками и их балансами
package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	er "github.com/normalniydada/case_infotecs/internal/domain/errors"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"github.com/normalniydada/case_infotecs/internal/domain/repository"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"github.com/shopspring/decimal"
)

// walletService реализует интерфейс WalletService.
// Содержит репозиторий для работы с данными кошельков.
type walletService struct {
	walletRepo repository.WalletRepository
}

// NewWalletService создает новый экземпляр сервиса для работы с кошельками.
//
// Параметры:
//   - walletRepo: репозиторий для доступа к данным кошельков
//
// Возвращает:
//   - service.WalletService: реализацию интерфейса сервиса кошельков
func NewWalletService(walletRepo repository.WalletRepository) service.WalletService {
	return &walletService{walletRepo: walletRepo}
}

// Balance возвращает текущий баланс указанного кошелька.
//
// Параметры:
//   - ctx: контекст выполнения
//   - address: адрес кошелька
//
// Возвращает:
//   - decimal.Decimal: текущий баланс кошелька
//   - error: ошибка, если кошелек не найден или произошла другая ошибка
//
// Возможные ошибки:
//   - ErrWalletNotFound: если кошелек не найден
//   - Другие ошибки репозитория: при проблемах доступа к данным
func (s *walletService) Balance(ctx context.Context, address string) (decimal.Decimal, error) {
	wallet, err := s.walletRepo.Wallet(ctx, address)
	if err != nil {
		return decimal.NewFromInt(0), fmt.Errorf("error while getting balance: %w", err)
	}

	return wallet.Balance, nil
}

// TransferMoney выполняет перевод средств между кошельками.
// Проверяет валидность параметров перед выполнением перевода.
//
// Параметры:
//   - ctx: контекст выполнения
//   - from: адрес кошелька отправителя
//   - to: адрес кошелька получателя
//   - amount: сумма перевода (должна быть положительной)
//
// Возвращает:
//   - error: ошибка, если перевод не удался
//
// Возможные ошибки:
//   - ErrSameWalletTransfer: при попытке перевода на тот же кошелек
//   - ErrInvalidAmount: при невалидной сумме перевода (<= 0)
//   - ErrInsufficientFunds: если недостаточно средств на кошельке отправителя
//   - ErrWalletNotFound: если один из кошельков не найден
func (s *walletService) TransferMoney(ctx context.Context, from, to string, amount decimal.Decimal) error {
	if from == to {
		return er.ErrSameWalletTransfer
	}

	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return er.ErrInvalidAmount
	}

	return s.walletRepo.Transfer(ctx, from, to, amount)
}

// CreateWallet создает новый кошелек с указанным начальным балансом.
// Генерирует уникальный адрес кошелька автоматически.
//
// Параметры:
//   - ctx: контекст выполнения
//   - balance: начальный баланс кошелька
//
// Возвращает:
//   - error: ошибка, если создание не удалось
func (s *walletService) CreateWallet(ctx context.Context, balance decimal.Decimal) error {
	wallet := models.Wallet{
		Address: generateWalletAddress(),
		Balance: balance,
	}

	return s.walletRepo.CreateWallet(ctx, &wallet)
}

// CountWallets возвращает общее количество кошельков в системе.
//
// Параметры:
//   - ctx: контекст выполнения
//
// Возвращает:
//   - int64: количество кошельков
//   - error: ошибка, если подсчет не удался
func (s *walletService) CountWallets(ctx context.Context) (int64, error) {
	return s.walletRepo.Count(ctx)
}

// generateWalletAddress генерирует уникальный адрес кошелька.
// Использует UUID и SHA-256 хеш для создания адреса.
//
// Возвращает:
//   - string: сгенерированный адрес кошелька (хеш в hex-формате)
func generateWalletAddress() string {
	u := uuid.New().String()
	hash := sha256.Sum256([]byte(u))
	return hex.EncodeToString(hash[:])
}
