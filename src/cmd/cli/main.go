package main

import (
	"context"
	"log"
	"os"
	"time"

	"feed/internal/config"
	"feed/internal/storage"

	"feed/internal/command"
)

func main() {
	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Инициализация баз данных
	storage.InitPostgresDB(ctx, cfg)
	defer storage.ClosePostgresDB()

	// CLI
	cmd := command.NewRootCMD()
	if err := cmd.Execute(); err != nil {
		log.Println("launch cli", err)
		os.Exit(1)
	}
}
