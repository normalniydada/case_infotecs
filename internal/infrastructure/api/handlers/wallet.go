package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	er "github.com/normalniydada/case_infotecs/internal/domain/errors"
	"github.com/normalniydada/case_infotecs/internal/domain/service"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/dto"
	"github.com/normalniydada/case_infotecs/internal/infrastructure/api/interfaces"
	"net/http"
)

type walletHandler struct {
	walletService service.WalletService
}

func NewWalletHandler(walletService service.WalletService) interfaces.WalletHandler {
	return &walletHandler{walletService: walletService}
}

func (h *walletHandler) Send(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"invalid json": "json is formatted incorrectly"})
	}

	err := h.walletService.TransferMoney(ctx, req.From, req.To, req.Amount)
	if err != nil {
		if errors.Is(err, er.ErrNotEnoughMoney) ||
			errors.Is(err, er.ErrInvalidAmount) ||
			errors.Is(err, er.ErrSameWalletTransfer) ||
			errors.Is(err, er.ErrWalletSenderNotFound) ||
			errors.Is(err, er.ErrWalletReceiverNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{"invalid value": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"transaction": "transaction failed and canceled"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "transaction succeeded"})
}

func (h *walletHandler) Balance(c echo.Context) error {
	ctx := c.Request().Context()

	address := c.Param("address")
	balance, err := h.walletService.Balance(ctx, address)
	if err != nil {
		if errors.Is(err, er.ErrWalletNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"wallet error": err.Error()})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get balance")
	}

	return c.JSON(http.StatusOK, map[string]float64{"balance": balance})
}
