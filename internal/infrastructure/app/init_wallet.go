package app

import (
	"context"
	"fmt"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"log"
	"sync"
	"sync/atomic"
)

type WalletInitializer struct {
	walletService service.WalletService
}

func NewWalletInitializer(walletService service.WalletService) *WalletInitializer {
	return &WalletInitializer{walletService: walletService}
}

func (wi *WalletInitializer) InitWallet(ctx context.Context, count int, balance float64) error {
	existingCount, err := wi.walletService.CountWallets(ctx)
	if err != nil {
		return err
	}

	if existingCount > 0 {
		log.Printf("[INFO] Wallets exist (%d), skip creation", existingCount)
		return nil
	}

	wg := &sync.WaitGroup{}
	errCh := make(chan error, count)
	var successCount int32

	for i := 0; i < count; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := wi.walletService.CreateWallet(ctx, balance)
			if err != nil {
				errCh <- fmt.Errorf("error initializing wallet %d: %w", i+1, err)
				return
			}
			atomic.AddInt32(&successCount, 1)
		}()
	}

	wg.Wait()
	close(errCh)

	log.Printf("[INFO] %d/%d wallets successfully created", successCount, count)

	if len(errCh) > 0 {
		return fmt.Errorf("%d errors occurred while creating wallets", len(errCh))
	}

	return nil
}
