package app

import (
	"context"
	"log"
)

func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := setupApplication(ctx)
	if err != nil {
		log.Fatalf("[ERROR] Error during initialization: %v", err)
	}
	defer app.Close()

	app.runServer()
	app.serverShutdown(ctx)
}
