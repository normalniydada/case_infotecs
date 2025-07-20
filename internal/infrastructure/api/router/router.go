package router

import (
	"github.com/labstack/echo/v4"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/interfaces"
)

func NewRouter(e *echo.Echo, walletHandler interfaces.WalletHandler, transactionHandler interfaces.TransactionHandler) {
	api := e.Group("/api")
	{
		api.GET("/wallet/:address/balance", walletHandler.Balance)
		api.GET("/transactions", transactionHandler.Last)
		api.POST("/send", walletHandler.Send)
	}
}
