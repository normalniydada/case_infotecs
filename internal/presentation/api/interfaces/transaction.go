// Package interfaces определяет контракты для HTTP-обработчиков API.
// Содержит интерфейсы, которые должны реализовывать обработчики маршрутов.
package interfaces

import (
	"github.com/labstack/echo/v4"
)

// TransactionHandler определяет контракт для обработчика операций с транзакциями.
// Реализации этого интерфейса должны обрабатывать запросы, связанные с историей транзакций.
type TransactionHandler interface {
	Last(c echo.Context) error
}
