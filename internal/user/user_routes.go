package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repository userRepository
}

func NewUserRoutes(app *fiber.App, repository userRepository) {
	handler := &UserHandler{
		repository: repository,
	}

	app.Post("/auth/register", handler.RegisterUser)
	app.Post("/auth/login", handler.LoginUser)
	app.Get("/auth/logout", handler.LogoutUser)
}
