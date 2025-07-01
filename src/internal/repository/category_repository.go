package repository

import (
	"context"

	"feed/internal/database"
	"feed/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresCategoryRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCategoryRepository() *PostgresCategoryRepository {
	db := database.GetDB()
	return &PostgresCategoryRepository{db: db}
}

func (repo *PostgresCategoryRepository) Create(category models.Category) error {
	_, err := repo.db.Exec(
		context.Background(),
		"INSERT INTO category (id, site_id, slug, name, is_active) VALUES ($1, $2, $3, $4, $5)",
		category.Id, category.SiteId, category.Slug, category.Name, category.Active,
	)

	return err
}
