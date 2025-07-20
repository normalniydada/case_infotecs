// Package postgres предоставляет реализацию хранилища данных на PostgreSQL.
// Содержит настройку подключения, миграции и базовые операции с БД.

package postgres

import (
	"fmt"
	"github.com/normalniydada/case_infotecs/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// createConnection устанавливает подключение к PostgreSQL и возвращает инстанс PostgresDB.
//
// Параметры:
//   - cfg: конфигурация подключения к БД
//
// Возвращает:
//   - *PostgresDB: инстанс для работы с БД
//   - error: ошибка подключения
//
// Процесс подключения:
//  1. Валидация конфигурации
//  2. Установка соединения через GORM
//  3. Настройка пула соединений
//  4. Проверка подключения (ping)
func createConnection(cfg *config.DatabaseConfig) (*PostgresDB, error) {
	// Формирование DSN строки подключения
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	// Валидация конфигурации
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	// Установка соединения через GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	// Настройка подключения из пула
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql DB: %w", err)
	}

	// Конфигурация пула
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)                                    // Макс. количество бездействующих соединений
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)                                    // Макс. количество открытых соединений
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute) // Макс. время жизни соединения

	// Проверка подключения
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

// validateConfig проверяет обязательные параметры подключения к БД.
//
// Параметры:
//   - cfg: конфигурация подключения
//
// Возвращает:
//   - error: ошибка валидации, если обязательные поля не заполнены
//
// Проверяемые поля:
//   - Host
//   - Port
//   - User
//   - Password
//   - DBName
func validateConfig(cfg *config.DatabaseConfig) error {
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Password == "" || cfg.DBName == "" {
		return fmt.Errorf("invalid database configuration")
	}
	return nil
}
