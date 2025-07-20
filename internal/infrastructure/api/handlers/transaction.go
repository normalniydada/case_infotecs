package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	er "github.com/normalniydada/case_infotecs/internal/domain/errors"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/dto"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/interfaces"
	"net/http"
	"strconv"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) interfaces.TransactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

func (h *transactionHandler) Last(c echo.Context) error {
	ctx := c.Request().Context()

	count, err := strconv.Atoi(c.QueryParam("count"))
	if count < 0 || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"invalid_count": er.ErrInvalidCount.Error()})
	}

	transactions, err := h.transactionService.LastNTransactions(ctx, count)
	if err != nil {
		if errors.Is(err, er.ErrTransactionNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"no transactions": er.ErrTransactionNotFound.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"transaction error:": "could not fetch transactions"})
	}

	return c.JSON(http.StatusOK, map[string][]dto.TransactionResponse{"transactions": transactions})
}
