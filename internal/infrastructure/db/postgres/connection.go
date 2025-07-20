package postgres

import (
	"fmt"
	"github.com/normalniydada/case_infotecs/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func createConnection(cfg *config.DatabaseConfig) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	// Настройка подключения из пула
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

func validateConfig(cfg *config.DatabaseConfig) error {
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Password == "" || cfg.DBName == "" {
		return fmt.Errorf("invalid database configuration")
	}
	return nil
}
