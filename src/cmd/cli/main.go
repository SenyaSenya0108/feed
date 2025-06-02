package main

import (
	"feed/internal/app"
	"feed/internal/cli"
	"feed/internal/database"
)

func main() {
	db := database.InitDB()
	ctx := &app.AppContext{
		DB: db,
	}
	cli.Execute(ctx)

	database.CLoseDB(db)
}
