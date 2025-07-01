package main

import (
	"context"

	"feed/internal/cli"
	"feed/internal/database"
)

func main() {
	ctx := context.Background()
	err := database.InitDB(ctx)
	if err != nil {
		return
	}

	cli.Execute()

	database.CLoseDB()
}
