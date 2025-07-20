// Package handlers предоставляет HTTP-обработчики для API сервиса кошельков.
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

// transactionHandler реализует интерфейс TransactionHandler.
// Обрабатывает HTTP-запросы, связанные с транзакциями.
type transactionHandler struct {
	transactionService service.TransactionService
}

// NewTransactionHandler создает новый экземпляр обработчика транзакций.
//
// Параметры:
//   - transactionService: сервис для работы с транзакциями
//
// Возвращает:
//   - interfaces.TransactionHandler: реализацию интерфейса обработчика
func NewTransactionHandler(transactionService service.TransactionService) interfaces.TransactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

// Last обрабатывает запрос на получение последних транзакций.
// GET /transactions/last?count={n}
//
// Параметры запроса:
//   - count: количество транзакций (положительное число)
//
// Возможные ответы:
//   - 200 OK: {"transactions": [...]} - успешный запрос
//   - 400 Bad Request: {"invalid_count": "..."} - невалидный параметр count
//   - 404 Not Found: {"no transactions": "..."} - транзакции не найдены
//   - 500 Internal Server Error: {"transaction error": "..."} - ошибка сервера
func (h *transactionHandler) Last(c echo.Context) error {
	ctx := c.Request().Context()

	count, err := strconv.Atoi(c.QueryParam("count"))
	if count < 0 || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"invalid_count": er.ErrInvalidCount.Error()})
	}

	// Получение транзакций из сервиса
	transactions, err := h.transactionService.LastNTransactions(ctx, count)
	if err != nil {
		if errors.Is(err, er.ErrTransactionNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"no transactions": er.ErrTransactionNotFound.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"transaction error:": "could not fetch transactions"})
	}

	// Успешный ответ
	return c.JSON(http.StatusOK, map[string][]dto.TransactionResponse{"transactions": transactions})
}
