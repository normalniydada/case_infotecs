package dto

import "time"

type TransactionResponse struct {
	From      string    `json:"sender_address"`
	To        string    `json:"receiver_address"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"date"`
}
