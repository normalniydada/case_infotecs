// Package postgres предоставляет реализацию хранилища данных на PostgreSQL.
// Содержит настройку подключения, миграции и базовые операции с БД.

package postgres

import (
	"fmt"
	"github.com/normalniydada/case_infotecs/internal/domain/models"
	"gorm.io/gorm"
)

// PostgresDB представляет соединение с PostgreSQL и реализует методы работы с БД.
// Инкапсулирует логику миграций и управления подключением.
type PostgresDB struct {
	db *gorm.DB
}

// GetDB возвращает экземпляр GORM для непосредственной работы с БД.
// Используется репозиториями для выполнения запросов.
//
// Возвращает:
//   - *gorm.DB: подключение к БД
func (p *PostgresDB) GetDB() *gorm.DB {
	return p.db
}

// Close корректно завершает соединение с базой данных.
// Должен вызываться при завершении работы приложения.
//
// Возвращает:
//   - error: ошибка закрытия соединения, если возникла
func (p *PostgresDB) Close() error {
	conn, err := p.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}
	return conn.Close()
}

// runMigrations выполняет автоматические миграции для моделей приложения.
// Создает необходимые таблицы и индексы в базе данных.
//
// Мигрируемые модели:
//   - models.Wallet: таблица кошельков
//   - models.Transaction: таблица транзакций
//
// Возвращает:
//   - error: ошибка выполнения миграций
func (p *PostgresDB) runMigrations() error {
	return p.db.AutoMigrate(
		&models.Wallet{},
		&models.Transaction{},
	)
}
