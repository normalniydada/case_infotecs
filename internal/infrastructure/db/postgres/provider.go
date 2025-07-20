package postgres

import (
	"github.com/normalniydada/case_infotecs/config"
	"sync"
)

var (
	dbInstance Database
	once       sync.Once
)

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
