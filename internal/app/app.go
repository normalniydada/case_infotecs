package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/normalniydada/case_infotecs/internal/config"
	"github.com/normalniydada/case_infotecs/internal/infra/database/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfig()
	log.Println("[INFO] Успешное считывание конфигурации: ", cfg)

	db, err := postgres.ProvideDBClient(&cfg.Database)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize database: %v", err)
	}
	log.Println("[INFO] Успешная инициализация базы данных...")
	defer func() {
		log.Println("[INFO] Закрытие соединения БД...")
		if sqlDB, err := db.GetDB().DB(); err == nil {
			if err := sqlDB.Close(); err != nil {
				log.Printf("[WARN] Ошибка при закрытии соединения с БД: %v", err)
			} else {
				log.Println("[INFO] Успешно закрытие соединения!")
			}
		}
	}()

	e := echo.New()
	e.HideBanner = true

	/*taskRepo := repository.NewTaskRepository(db.GetDB())
	taskService := service.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)
	router.NewRouter(e, taskHandlers)*/

	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: e,
	}

	go func() {
		log.Printf("[INFO] Сервер запущен... %s", serverAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] Сервер-ошибка: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[INFO] Запуск graceful shutdown...")

	// Даём 10 секунд на завершение операций
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("[ERROR] Отключение сервера не удалось: %v", err)
	} else {
		log.Println("[INFO] Сервер оставлен успешно!")
	}
}
