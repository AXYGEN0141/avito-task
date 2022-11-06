package app

import (
	"avito-task/internal/config"
	"avito-task/internal/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

const confFile = ".env"

// init parses .env file.
// If failed prints an error.
func init() {
	if err := godotenv.Load(confFile); err != nil {
		log.Printf(".env file does not exist ...\n App starting with default configuration")
	}
}

// Start starts the application.
func Start() {
	cfg := config.GetConfig()

	app := server.NewApp(cfg)

	app.Init()

	// Graceful Shutdown.
	ctx, shutdown := context.WithCancel(context.Background())
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		shutdown()
	}()
	app.Run(ctx)
}
