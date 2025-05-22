package client

import (
	"io"
	"log"
	"net/http"
)

func Category() []byte {
	resp, err := http.Get("")

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
