package dto

import (
	"github.com/shopspring/decimal"
	"time"
)

type TransactionResponse struct {
	From      string          `json:"sender_address"`
	To        string          `json:"receiver_address"`
	Amount    decimal.Decimal `json:"amount"`
	CreatedAt time.Time       `json:"date"`
}
