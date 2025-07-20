// Package postgres предоставляет реализацию хранилища данных на PostgreSQL.
// Содержит настройку подключения, миграции и базовые операции с БД.

package postgres

import (
	"github.com/normalniydada/case_infotecs/config"
	"sync"
)

var (
	// dbInstance содержит singleton-экземпляр подключения к БД
	dbInstance Database
	// once гарантирует однократную инициализацию подключения
	once sync.Once
)

// ProvideDBClient создает и возвращает singleton-экземпляр подключения к БД.
// Реализует паттерн Singleton с thread-safe инициализацией.
//
// Параметры:
//   - cfg: конфигурация подключения к БД
//
// Возвращает:
//   - Database: интерфейс для работы с БД
//   - error: ошибка инициализации, если возникла
//
// Процесс инициализации:
//  1. Установка соединения с БД (createConnection)
//  2. Выполнение миграций (runMigrations)
//  3. Сохранение экземпляра в dbInstance
//
// Особенности:
//   - Гарантирует однократную инициализацию (sync.Once)
//   - Потокобезопасна
//   - Возвращает один и тот же экземпляр при повторных вызовах
//   - При ошибке инициализации последующие вызовы вернут ту же ошибку
func ProvideDBClient(cfg *config.DatabaseConfig) (Database, error) {
	var err error
	once.Do(func() {
		var db *PostgresDB
		db, err = createConnection(cfg)
		if err == nil {
			err = db.runMigrations()
		}
		dbInstance = db
	})
	return dbInstance, err
}
