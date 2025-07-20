package postgres

import (
	"fmt"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func (p *PostgresDB) GetDB() *gorm.DB {
	return p.db
}

func (p *PostgresDB) Close() error {
	conn, err := p.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}
	return conn.Close()
}

func (p *PostgresDB) runMigrations() error {
	return p.db.AutoMigrate(
		&models.Wallet{},
		&models.Transaction{},
	)
}
