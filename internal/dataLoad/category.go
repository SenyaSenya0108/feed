package dataLoad

import (
	"encoding/json"
	"feed/internal/client"
	"fmt"
)

type Category struct {
	SiteId   int    `json:"site_id"`
	Active   string `json:"active"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentId int    `json:"parent_id"`
}

func loadCategory() {
	var categories []Category
	response := client.Category()

	json.Unmarshal(response, &categories)

	for _, category := range categories {
		fmt.Printf("%+v\n", category)
	}
}

func Load() {
	loadCategory()
}
