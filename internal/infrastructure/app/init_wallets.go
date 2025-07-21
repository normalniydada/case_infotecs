// Package app предоставляет точку входа и основную логику запуска приложения.
// Управляет жизненным циклом приложения, инициализацией и graceful shutdown.
package app

import (
	"context"
	"fmt"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"github.com/shopspring/decimal"
	"log"
	"sync"
	"sync/atomic"
)

// WalletInitializer отвечает за инициализацию кошельков при старте приложения.
// Позволяет создать необходимое количество кошельков с указанным балансом.
type WalletInitializer struct {
	walletService service.WalletService
}

// NewWalletInitializer создает новый экземпляр WalletInitializer.
//
// Параметры:
//   - walletService: сервис для операций с кошельками
//
// Возвращает:
//   - *WalletInitializer: инициализатор кошельков
func NewWalletInitializer(walletService service.WalletService) *WalletInitializer {
	return &WalletInitializer{walletService: walletService}
}

// InitWallet инициализирует указанное количество кошельков с заданным балансом.
// Если кошельки уже существуют в системе, инициализация пропускается.
//
// Параметры:
//   - ctx: контекст выполнения
//   - count: количество создаваемых кошельков
//   - balance: начальный баланс для каждого кошелька
//
// Возвращает:
//   - error: ошибка, если не удалось создать кошельки
//
// Особенности:
//   - Проверяет существование кошельков перед созданием
//   - Использует конкурентное создание для повышения производительности
//   - Логирует процесс инициализации
//   - Возвращает агрегированную ошибку при частичном сбое
func (wi *WalletInitializer) InitWallet(ctx context.Context, count int, balance decimal.Decimal) error {
	// Проверяем существующие кошельки
	existingCount, err := wi.walletService.CountWallets(ctx)
	if err != nil {
		return err
	}

	if existingCount > 0 {
		log.Printf("[INFO] Wallets exist (%d), skip creation", existingCount)
		return nil
	}

	// Настраиваем механизм конкурентного создания
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

	// Обработка ошибок
	if len(errCh) > 0 {
		return fmt.Errorf("%d errors occurred while creating wallets", len(errCh))
	}

	return nil
}
