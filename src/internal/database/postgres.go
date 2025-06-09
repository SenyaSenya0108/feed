package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *PostgresPool

type PostgresPool struct {
	pool *pgxpool.Pool
}

func InitDB(ctx context.Context) {
	connect := os.Getenv("DB_URL")

	pool, err := pgxpool.New(ctx, connect)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Printf("Failed to ping database: %v\n", err)
		os.Exit(1)
	}

	db = &PostgresPool{pool: pool}
	log.Println("Successfully connected to the database")
}

func GetDB() *pgxpool.Pool {
	return db.pool
}

func CLoseDB() {
	db.pool.Close()
	log.Println("The connection to the database is closed ")
}
