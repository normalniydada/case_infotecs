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

type walletService struct {
	walletRepo repository.WalletRepository
}

func NewWalletService(walletRepo repository.WalletRepository) service.WalletService {
	return &walletService{walletRepo: walletRepo}
}

func (s *walletService) Balance(ctx context.Context, address string) (decimal.Decimal, error) {
	wallet, err := s.walletRepo.Wallet(ctx, address)
	if err != nil {
		return decimal.NewFromInt(0), fmt.Errorf("error while getting balance: %w", err)
	}

	return wallet.Balance, nil
}

func (s *walletService) TransferMoney(ctx context.Context, from, to string, amount decimal.Decimal) error {
	if from == to {
		return er.ErrSameWalletTransfer
	}

	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return er.ErrInvalidAmount
	}

	return s.walletRepo.Transfer(ctx, from, to, amount)
}

func (s *walletService) CreateWallet(ctx context.Context, balance decimal.Decimal) error {
	wallet := models.Wallet{
		Address: generateWalletAddress(),
		Balance: balance,
	}

	return s.walletRepo.CreateWallet(ctx, &wallet)
}

func (s *walletService) CountWallets(ctx context.Context) (int64, error) {
	return s.walletRepo.Count(ctx)
}

func generateWalletAddress() string {
	u := uuid.New().String()
	hash := sha256.Sum256([]byte(u))
	return hex.EncodeToString(hash[:])
}
