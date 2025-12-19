package storage

import (
	"context"
	"fmt"
	"time"

	"feed/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pgDB *PostgresPool

type PostgresPool struct {
	pool *pgxpool.Pool
}

func InitPostgresDB(ctx context.Context, cfg *config.Config) {
	connect := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name,
	)

	dbContext, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(dbContext, connect)
	if err != nil {
		panic(fmt.Sprintf("unable to connect to database: %v\n", err))
	}

	if err := pool.Ping(context.Background()); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v\n", err))
	}

	pgDB = &PostgresPool{pool: pool}
}

func PostgresDB() *pgxpool.Pool {
	return pgDB.pool
}

func ClosePostgresDB() {
	pgDB.pool.Close()
}
