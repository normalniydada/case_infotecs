package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host            string
	Port            int
	DBName          string
	User            string
	Password        string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

type ServerConfig struct {
	Host string
	Port int
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("[WARN] .env файл не был найден: %v", err)
	}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Ошибка при чтении файла конфигурации: %v", err)
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
