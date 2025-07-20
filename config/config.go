// Package config предоставляет функциональность для загрузки конфигурации приложения
// из YAML-файлов и переменных окружения. Поддерживается:
// - Чтение конфигурации из файла config.yaml
// - Загрузка чувствительных данных из .env файла
// - Иерархическая структура конфигурации
package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

// Config представляет основную структуру конфигурации приложения.
// Содержит все необходимые настройки для работы сервера и базы данных.
type Config struct {
	Server   ServerConfig   // Настройки HTTP сервера
	Database DatabaseConfig // Настройки подключения к базе данных
}

// DatabaseConfig содержит параметры для подключения к базе данных.
// Позволяет настроить пул соединений и параметры аутентификации.
type DatabaseConfig struct {
	Host            string // Хост базы данных
	Port            int    // Порт базы данных
	DBName          string // Имя базы данных
	User            string // Пользователь БД (загружается из .env)
	Password        string // Пароль пользователя (загружается из .env)
	SSLMode         string // Режим SSL (disable, require, verify-full и т.д.)
	MaxIdleConns    int    // Макс. количество неактивных соединений в пуле
	MaxOpenConns    int    // Макс. количество открытых соединений
	ConnMaxLifetime int    // Макс. время жизни соединения в секундах
}

// ServerConfig содержит параметры HTTP сервера.
type ServerConfig struct {
	Host string // Хост для запуска сервера
	Port int    // Порт для запуска сервера
}

// NewConfig создает и инициализирует новый объект Config.
// Загружает конфигурацию в следующем порядке:
//  1. Пытается загрузить переменные окружения из .env файла
//  2. Читает основной конфигурационный файл config.yaml
//  3. Комбинирует настройки из файла и переменных окружения
//
// Возвращает:
//
//	*Config - указатель на загруженную конфигурацию
func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("[WARN] .env the file was not found: %v", err)
	}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Error reading configuration file: %v", err)
	}

	cfg := &Config{
		Server: ServerConfig{
			Host: v.GetString("server.host"),
			Port: v.GetInt("server.port"),
		},
		Database: DatabaseConfig{
			Host:            v.GetString("database.host"),
			Port:            v.GetInt("database.port"),
			DBName:          v.GetString("database.dbname"),
			User:            os.Getenv("DATABASE_USER"),
			Password:        os.Getenv("DATABASE_PASSWORD"),
			SSLMode:         v.GetString("database.sslmode"),
			MaxIdleConns:    v.GetInt("database.max_idle_conns"),
			MaxOpenConns:    v.GetInt("database.max_open_conns"),
			ConnMaxLifetime: v.GetInt("database.conn_max_lifetime"),
		},
	}

	return cfg
}
