package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Address uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
	Balance int64     `gorm:"type:bigint;not null;default:0"`
}

type Transaction struct {
	gorm.Model
	From   uuid.UUID `gorm:"type:uuid;not null"`
	To     uuid.UUID `gorm:"type:uuid;not null"`
	Amount int64     `gorm:"type:bigint;not null"`

	FromWallet Wallet `gorm:"foreignKey:From;references:Address;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ToWallet   Wallet `gorm:"foreignKey:To;references:Address;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
