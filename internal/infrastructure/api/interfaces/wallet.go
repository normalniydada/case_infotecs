package interfaces

import (
	"github.com/labstack/echo/v4"
)

type WalletHandler interface {
	Send(c echo.Context) error
	Balance(c echo.Context) error
}
