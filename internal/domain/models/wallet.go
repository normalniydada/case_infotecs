package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Address string          `gorm:"type:string;uniqueIndex;not null"`
	Balance decimal.Decimal `gorm:"type:numeric(20,8);not null;default:0"`
}
