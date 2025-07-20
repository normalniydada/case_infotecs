package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/normalniydada/case_infotecs/config"
	"github.com/normalniydada/case_infotecs/internal/application/transaction"
	"github.com/normalniydada/case_infotecs/internal/application/wallet"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/handlers"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/router"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/db/postgres"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/db/postgres/repositories"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
)

type Application struct {
	cfg                *config.Config
	echo               *echo.Echo
	db                 *gorm.DB
	closers            []func()
	walletService      service.WalletService
	transactionService service.TransactionService
}

func setupApplication(ctx context.Context) (*Application, error) {
	cfg := config.NewConfig()
	log.Println("[INFO] Configuration loaded successfully")

	db, err := postgres.ProvideDBClient(&cfg.Database)
	if err != nil {
		return nil, err
	}
	log.Println("[INFO] The database is connected")

	walletRepo := repositories.NewWalletRepository(db.GetDB())
	transactionRepo := repositories.NewTransactionRepository(db.GetDB())

	app := &Application{
		cfg:                cfg,
		echo:               echo.New(),
		walletService:      wallet.NewWalletService(walletRepo),
		transactionService: transaction.NewTransactionService(transactionRepo),
	}

	app.closers = append(app.closers, func() {
		if sqlDB, err := db.GetDB().DB(); err == nil {
			if err := sqlDB.Close(); err != nil {
				log.Printf("[WARN] Error closing connection to DB: %v", err)
			} else {
				log.Println("[INFO] The connection to the database was closed")
			}
		}
	})

	app.setupEcho()

	if err := app.initWallets(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) setupEcho() {
	a.echo.HideBanner = true
	a.echo.Use(middleware.Recover(), middleware.Logger())

	walletHandler := handlers.NewWalletHandler(a.walletService)
	transactionHandler := handlers.NewTransactionHandler(a.transactionService)

	router.NewRouter(a.echo, walletHandler, transactionHandler)
}

func (a *Application) initWallets(ctx context.Context) error {
	initializer := NewWalletInitializer(a.walletService)
	return initializer.InitWallet(ctx, 10, decimal.NewFromFloat(100.0))
}

func (a *Application) Close() {
	for _, closer := range a.closers {
		closer()
	}
}
