package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Address uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
	Balance int64     `gorm:"type:bigint;not null;default:0"`
}
