package dto

import "github.com/shopspring/decimal"

// TransactionRequest представляет структуру запроса на выполнение перевода между кошельками.
// Используется для десериализации входящих HTTP-запросов в API.
type TransactionRequest struct {
	From   string          `json:"from"`
	To     string          `json:"to"`
	Amount decimal.Decimal `json:"amount"`
}
