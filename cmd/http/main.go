package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	c "github.com/techwithmat/kanban-app/config"
	"github.com/techwithmat/kanban-app/pkg/database"
)

func main() {
	ctx := context.Background()
	config := c.NewConfig()

	db, err := database.NewDBConnection(ctx, config.Database)

	if err != nil {
		log.Println("Unable to connect to database")
	}

	defer db.Close(ctx)

	app := fiber.New()
	app.Listen(":3000")
}
