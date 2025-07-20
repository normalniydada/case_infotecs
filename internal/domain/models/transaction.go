package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	From   string          `gorm:"type:string;not null"`
	To     string          `gorm:"type:string;not null"`
	Amount decimal.Decimal `gorm:"type:numeric(20,8);not null"`
}
