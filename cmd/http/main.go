package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	c "github.com/techwithmat/kanban-app/config"
	"github.com/techwithmat/kanban-app/internal/user"
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
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://kanbi.vercel.app",
		AllowHeaders:     "Content-Type",
		AllowCredentials: true,
	}))

	userRepository := user.NewUserRepository(db)
	user.NewUserRoutes(app, userRepository)

	app.Listen(":3000")
}
