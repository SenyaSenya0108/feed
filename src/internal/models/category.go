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
