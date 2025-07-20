package dto

import "github.com/shopspring/decimal"

type TransactionRequest struct {
	From   string          `json:"from"`
	To     string          `json:"to"`
	Amount decimal.Decimal `json:"amount"`
}
