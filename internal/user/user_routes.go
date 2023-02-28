package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techwithmat/kanbi/internal/middleware"
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
	app.Get("/auth/session", middleware.AuthMiddleware, handler.GetUserSession)
}
