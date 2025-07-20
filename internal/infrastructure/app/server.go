// Package app предоставляет точку входа и основную логику запуска приложения.
// Управляет жизненным циклом приложения, инициализацией и graceful shutdown.

package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// runServer запускает HTTP-сервер приложения в отдельной goroutine.
// Сервер запускается на адресе, указанном в конфигурации приложения.
//
// Особенности:
//   - Логирует адрес запуска сервера
//   - В случае ошибки запуска завершает приложение с кодом 1
//   - Игнорирует ошибку http.ErrServerClosed (возникает при нормальном shutdown)
//
// Безопасность:
//   - Запуск в отдельной goroutine позволяет продолжить выполнение main потока
func (a *Application) runServer() {
	serverAddr := fmt.Sprintf("%s:%d", a.cfg.Server.Host, a.cfg.Server.Port)
	go func() {
		log.Printf("[INFO] Server started... %s", serverAddr)
		if err := a.echo.Start(serverAddr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] Server error: %v", err)
		}
	}()
}

// serverShutdown реализует механизм graceful shutdown сервера.
// Ожидает сигналов завершения (SIGINT, SIGTERM) и корректно останавливает сервер.
//
// Параметры:
//   - ctx: контекст приложения для propagation отмены
//
// Особенности:
//   - Регистрирует обработчики сигналов завершения
//   - Устанавливает таймаут 10 секунд на завершение операций
//   - Логирует процесс shutdown
//   - Гарантирует вызов cancel функции контекста
//
// Graceful shutdown включает:
//  1. Прекращение приема новых соединений
//  2. Ожидание завершения обработки текущих запросов
//  3. Корректное освобождение ресурсов
func (a *Application) serverShutdown(ctx context.Context) {
	// Канал для получения сигналов ОС
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[INFO] Start graceful shutdown...")

	// Контекст с таймаутом для shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel() // Гарантированное освобождение ресурсов

	// Остановка сервера
	if err := a.echo.Shutdown(shutdownCtx); err != nil {
		log.Printf("[ERROR] Server shutdown failed: %v", err)
	} else {
		log.Println("[INFO] Server abandoned")
	}
}
