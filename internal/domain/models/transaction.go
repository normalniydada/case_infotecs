// Package models содержит бизнес-сущности и их представление в базе данных.
// Определяет структуры данных, используемые на всех уровнях приложения.

package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Transaction представляет модель транзакции между кошельками в системе.
// Содержит информацию об отправителе, получателе и сумме перевода.
// Реализует gorm.Model для базовых полей (ID, CreatedAt, UpdatedAt, DeletedAt).
type Transaction struct {
	gorm.Model
	From   string          `gorm:"type:string;not null"`
	To     string          `gorm:"type:string;not null"`
	Amount decimal.Decimal `gorm:"type:numeric(20,8);not null"`
}
