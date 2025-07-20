package interfaces

import (
	"github.com/labstack/echo/v4"
)

type TransactionHandler interface {
	Last(c echo.Context) error
}
