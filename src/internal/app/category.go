package app

import (
	"encoding/json"
	"feed/internal/models"
	"feed/internal/repository"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type Category struct {
	SiteId   int    `json:"site_id"`
	Active   string `json:"active"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentId int    `json:"parent_id"`
}

func CategoryLoad() {
	categories := getCategoriesFromResponse()
	repo := repository.NewPostgresCategoryRepository()

	for _, value := range categories {
		category := convertModel(&value)
		if err := repo.Create(*category); err != nil {
			log.Fatalln(err)
		}
	}
}

func getCategoriesFromResponse() []Category {
	var categories []Category
	response := getCategory()

	json.Unmarshal(response, &categories)

	return categories
}

func convertModel(category *Category) *models.Category {
	categoryId, err := uuid.NewUUID()

	if err != nil {
		log.Fatal(err)
	}

	categoryActive, err := strconv.ParseBool(category.Active)

	if err != nil {
		log.Fatalln(err)
	}

	model := &models.Category{
		Id:     categoryId,
		SiteId: category.SiteId,
		Slug:   category.Slug,
		Name:   category.Name,
		Active: categoryActive,
	}

	return model
}
