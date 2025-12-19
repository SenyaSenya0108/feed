package repository

import (
	"feed/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BaseRepository struct {
	db *pgxpool.Pool
}

func NewBase() *BaseRepository {
	db := storage.PostgresDB()
	return &BaseRepository{db: db}
}
