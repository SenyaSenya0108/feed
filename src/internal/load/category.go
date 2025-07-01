package load

import (
	"log"
	"strconv"

	"feed/internal/models"
	"feed/internal/repository"
	"feed/internal/routes"
)

type CategoryService struct {
	repo *repository.PostgresCategoryRepository
}

func NewCategoryService() *CategoryService {
	repo := repository.NewPostgresCategoryRepository()
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CategorySync() {
	categories, _ := routes.GetCategorySyncData()

	for _, value := range *categories {
		category := convertModel(&value)
		if err := s.repo.Create(*category); err != nil {
			log.Fatalln(err)
		}
	}
}

func convertModel(category *models.CategorySyncDTO) *models.Category {
	categoryActive, err := strconv.ParseBool(category.Active)
	if err != nil {
		log.Fatalln(err)
	}

	model := &models.Category{
		SiteId: category.SiteId,
		Slug:   category.Slug,
		Name:   category.Name,
		Active: categoryActive,
	}

	return model
}
