package main

import (
	"context"
	"log"
	"os"
	"time"

	"feed/internal/broker"
	"feed/internal/command"
	"feed/internal/config"
	"feed/internal/storage"
)

func main() {
	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Инициализация баз данных
	storage.InitPostgresDB(ctx, cfg)
	defer storage.ClosePostgresDB()

	// брокер
	err := broker.Connect()
	if err != nil {
		log.Println("connect to broker", err)
		os.Exit(1)
	}

	// CLI
	cmd := command.NewRootCMD()
	if err := cmd.Execute(); err != nil {
		log.Println("launch cli", err)
		os.Exit(1)
	}
}
