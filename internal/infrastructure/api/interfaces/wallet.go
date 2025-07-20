// Package interfaces определяет контракты для HTTP-обработчиков API.

package interfaces

import (
	"github.com/labstack/echo/v4"
)

// WalletHandler определяет контракт для обработчика операций с кошельками.
// Описывает методы API для работы с кошельками и переводами средств.
type WalletHandler interface {
	Send(c echo.Context) error
	Balance(c echo.Context) error
}
