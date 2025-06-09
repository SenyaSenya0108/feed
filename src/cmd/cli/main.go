package main

import (
	"context"
	"feed/internal/cli"
	"feed/internal/database"
)

func main() {
	ctx := context.Background()
	database.InitDB(ctx)

	cli.Execute()

	database.CLoseDB()
}
