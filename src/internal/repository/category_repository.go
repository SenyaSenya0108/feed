package repository

import "feed/internal/models"

type CategoryRepository interface {
	Create(category models.Category) error
}
