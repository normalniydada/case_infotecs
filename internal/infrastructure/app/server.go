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

func (a *Application) runServer() {
	serverAddr := fmt.Sprintf("%s:%d", a.cfg.Server.Host, a.cfg.Server.Port)
	go func() {
		log.Printf("[INFO] Server started... %s", serverAddr)
		if err := a.echo.Start(serverAddr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] Server error: %v", err)
		}
	}()
}

func (a *Application) serverShutdown(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[INFO] Start graceful shutdown...")

	// 10 секунд на завершение операций
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	if err := a.echo.Shutdown(shutdownCtx); err != nil {
		log.Printf("[ERROR] Server shutdown failed: %v", err)
	} else {
		log.Println("[INFO] Server abandoned")
	}
}
