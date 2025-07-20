// Package models содержит бизнес-сущности и их представление в базе данных.
// Определяет структуры данных, используемые на всех уровнях приложения.

package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Wallet представляет модель кошелька в системе.
// Содержит уникальный адрес и текущий баланс.
// Наследует базовые поля gorm.Model (ID, CreatedAt, UpdatedAt, DeletedAt).
// Используется для хранения информации о пользовательских кошельках и их балансах.
type Wallet struct {
	gorm.Model
	Address string          `gorm:"type:string;uniqueIndex;not null"`
	Balance decimal.Decimal `gorm:"type:numeric(20,8);not null;default:0"`
}
