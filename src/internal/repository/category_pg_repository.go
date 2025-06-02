package repository

import (
	"database/sql"
	"feed/internal/models"
)

type PostgresCategoryRepository struct {
	db *sql.DB
}

func NewPostgresCategoryRepository(db *sql.DB) *PostgresCategoryRepository {
	return &PostgresCategoryRepository{db: db}
}

func (repo *PostgresCategoryRepository) Create(category models.Category) error {
	_, err := repo.db.Exec("INSERT INTO category (id, site_id, slug, name, is_active) VALUES ($1, $2, $3, $4, $5)", category.Id, category.SiteId, category.Slug, category.Name, category.Active)

	return err
}
