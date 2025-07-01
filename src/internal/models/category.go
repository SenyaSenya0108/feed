package models

import (
	"github.com/google/uuid"
)

type Category struct {
	Id     uuid.UUID
	SiteId int
	Slug   string
	Name   string
	Active bool
}

type CategorySyncDTO struct {
	SiteId   int    `json:"site_id"`
	Active   string `json:"active"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentId int    `json:"parent_id"`
}

type CategoriesSyncDTO []CategorySyncDTO

type CategoryRepository interface {
	Create(category Category) error
}
