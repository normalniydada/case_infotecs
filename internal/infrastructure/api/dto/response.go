// Package dto (Data Transfer Objects) содержит модели для взаимодействия с API.
// Определяет структуры данных для входящих/исходящих запросов.
package dto

import (
	"github.com/shopspring/decimal"
	"time"
)

// TransactionResponse представляет структуру ответа с информацией о транзакции.
// Используется для сериализации данных о транзакции в API-ответах.
type TransactionResponse struct {
	From      string          `json:"sender_address"`
	To        string          `json:"receiver_address"`
	Amount    decimal.Decimal `json:"amount"`
	CreatedAt time.Time       `json:"date"`
}
