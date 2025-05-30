package app

import (
	"io"
	"log"
	"net/http"
	"os"
)

func getCategory() []byte {
	resp, err := http.Get(os.Getenv("CATEGORY_URI"))

	if err != nil {
		panic("Запрос неудачный")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
