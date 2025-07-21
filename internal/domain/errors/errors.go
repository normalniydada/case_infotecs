// Package errors определяет доменные ошибки приложения, разделенные по уровням архитектуры.
// Ошибки сгруппированы в соответствии с их происхождением в слоистой архитектуре:
package errors

import "errors"

// Ошибки уровня репозитория (data access layer).
// Возникают при взаимодействии с хранилищем данных.
var (
	// ErrWalletNotFound возвращается при попытке доступа к несуществующему кошельку.
	// HTTP-аналог: 404 Not Found
	ErrWalletNotFound = errors.New("wallet not found")

	// ErrWalletSenderNotFound возвращается когда кошелек отправителя не найден.
	// HTTP-аналог: 400 Bad Request
	ErrWalletSenderNotFound = errors.New("sender's wallet not found")

	// ErrWalletReceiverNotFound возвращается когда кошелек получателя не найден.
	// HTTP-аналог: 400 Bad Request
	ErrWalletReceiverNotFound = errors.New("receiver's wallet not found")

	// ErrWalletExists возвращается при попытке создать уже существующий кошелек.
	ErrWalletExists = errors.New("wallet already exists")

	// ErrNotEnoughMoney возвращается при недостаточном балансе для перевода.
	// HTTP-аналог: 400 Bad Request
	ErrNotEnoughMoney = errors.New("insufficient funds in the sender's wallet")
)

// Ошибки уровня сервиса (business logic layer).
// Возникают при нарушении бизнес-правил приложения.
var (
	// ErrTransactionNotFound возвращается когда транзакции не найдены.
	// HTTP-аналог: 404 Not Found
	ErrTransactionNotFound = errors.New("no transactions")

	// ErrSameWalletTransfer возвращается при попытке перевода самому себе.
	// HTTP-аналог: 400 Bad Request
	ErrSameWalletTransfer = errors.New("impossible to send money to yourself")

	// ErrInvalidAmount возвращается при невалидной сумме перевода (<= 0).
	// HTTP-аналог: 400 Bad Request
	ErrInvalidAmount = errors.New("the sum must be positive")
)

// Ошибки уровня обработчиков (API layer).
// Возникают при обработке входящих HTTP-запросов.
var (
	// ErrInvalidCount возвращается при невалидном параметре count в запросе.
	// Используется для пагинации и лимитирования выборок.
	// HTTP-аналог: 400 Bad Request
	ErrInvalidCount = errors.New("invalid count query-params")
)
