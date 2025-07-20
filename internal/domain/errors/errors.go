package errors

import "errors"

// Repo
var (
	ErrWalletNotFound         = errors.New("wallet not found")
	ErrWalletSenderNotFound   = errors.New("sender's wallet not found")
	ErrWalletReceiverNotFound = errors.New("receiver's wallet not found")
	ErrWalletExists           = errors.New("wallet already exists")
	ErrNotEnoughMoney         = errors.New("insufficient funds in the sender's wallet")
)

// Service
var (
	ErrTransactionNotFound = errors.New("no transactions")
	ErrSameWalletTransfer  = errors.New("impossible to send money to yourself")
	ErrInvalidAmount       = errors.New("the sum must be positive")
)

// Handlers
var (
	ErrInvalidCount = errors.New("invalid count query-params")
)
