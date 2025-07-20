// Package router предоставляет функциональность для настройки маршрутизации API.
// Определяет структуру эндпоинтов и связывает их с соответствующими обработчиками.
// Package router предоставляет функциональность для настройки маршрутизации API.
// Определяет структуру эндпоинтов и связывает их с соответствующими обработчиками.

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/interfaces"
)

// NewRouter инициализирует маршруты API и связывает их с обработчиками.
//
// Параметры:
//   - e: экземпляр Echo для настройки маршрутов
//   - walletHandler: обработчик операций с кошельками
//   - transactionHandler: обработчик операций с транзакциями
//
// Определяемые маршруты:
//
//	GET    /api/wallet/:address/balance - Получение баланса кошелька
//	GET    /api/transactions           - Получение последних транзакций
//	POST   /api/send                   - Перевод средств между кошельками
//
// Группировка:
//
//	Все маршруты префиксируются /api для версионирования и разделения API.
func NewRouter(e *echo.Echo, walletHandler interfaces.WalletHandler, transactionHandler interfaces.TransactionHandler) {
	api := e.Group("/api")
	{
		api.GET("/wallet/:address/balance", walletHandler.Balance)
		api.GET("/transactions", transactionHandler.Last)
		api.POST("/send", walletHandler.Send)
	}
}
