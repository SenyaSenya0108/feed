package routes

import (
	"encoding/json"
	"feed/internal/models"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func GetCategorySyncData() (*models.CategoriesSyncDTO, error) {
	resp, err := client.Get(os.Getenv("CATEGORY_URI"))
	if err != nil {
		log.Printf("Category request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("status code: %d %s", resp.StatusCode, resp.Status)
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Read response body: %v", err)
		return nil, err
	}

	var categories models.CategoriesSyncDTO
	err = json.Unmarshal(data, &categories)
	if err != nil {
		log.Printf("Decode data %v", err)
		return nil, err
	}

	return &categories, nil
}
