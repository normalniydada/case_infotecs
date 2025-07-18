package wallet

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/normalniydada/case_infotecs/internal/repository"
	"math"
)

var (
	ErrSameWalletTransfer = errors.New("невозможно отправить деньги самому себе")
	ErrInvalidAmount      = errors.New("сумма должна быть положительной")
)

const precision = 8 // количество знаков после запятой

type walletService struct {
	walletRepo repository.WalletRepository
}

func NewWalletService(walletRepo repository.WalletRepository) *walletService {
	return &walletService{walletRepo: walletRepo}
}

func (s *walletService) Balance(ctx context.Context, address uuid.UUID) (float64, error) {
	wallet, err := s.walletRepo.Wallet(ctx, address)
	if err != nil {
		return 0, fmt.Errorf("ошибка при получении баланса: %w", err)
	}

	return convertIntToFloat(wallet.Balance), nil
}

func (s *walletService) TransferMoney(ctx context.Context, from, to uuid.UUID, amount float64) error {
	if from == to {
		return ErrSameWalletTransfer
	}

	if amount < 0 {
		return ErrInvalidAmount
	}

	return s.walletRepo.Transfer(ctx, from, to, convertFloatToInt(amount))
}

func convertIntToFloat(x int64) float64 {
	return float64(x) / (math.Pow(10, precision))
}

func convertFloatToInt(x float64) int64 {
	return int64(x * math.Pow(10, precision))
}
