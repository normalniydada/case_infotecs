package postgres

import (
	"fmt"
	"github.com/normalniydada/case_infotecs/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Database interface {
	GetDB() *gorm.DB
	Close() error
}

type PostgresDB struct {
	db *gorm.DB
}

func createConnection(cfg *config.DatabaseConfig) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к БД: %w", err)
	}

	// Настройка подключения из пула
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить sql DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Не удалось ping БД: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) GetDB() *gorm.DB {
	return p.db
}

func (p *PostgresDB) Close() error {
	conn, err := p.db.DB()
	if err != nil {
		return fmt.Errorf("Не удалось получить соединение с базой данных: %w", err)
	}
	return conn.Close()
}

var (
	dbInstance Database
	once       sync.Once
)

func ProvideDBClient(cfg *config.DatabaseConfig) (Database, error) {
	var err error
	once.Do(func() {
		dbInstance, err = createConnection(cfg)
	})
	if err != nil {
		return nil, err
	}
	return dbInstance, err

}

func validateConfig(cfg *config.DatabaseConfig) error {
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Password == "" || cfg.DBName == "" {
		return fmt.Errorf("Невалидная конфигурация БД")
	}
	return nil
}
