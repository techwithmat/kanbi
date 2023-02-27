package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/techwithmat/kanban-app/pkg/encryption"
	v "github.com/techwithmat/kanban-app/pkg/validation"
)

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	ctx := c.Context()
	var user RegisterRequest

	// validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	if errors := v.ValidateStruct(user); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Hashing the password and then inserting the user into the database.
	hashedPassword, err := encryption.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	user.Password = hashedPassword

	createdUserId, err := h.repository.Insert(ctx, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"user_id": createdUserId,
	})
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	ctx := c.Context()
	var input LoginRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	if errors := v.ValidateStruct(input); errors != nil {
		return c.Status(400).JSON(errors)
	}

	user, err := h.repository.GetByEmail(ctx, input.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	if err := encryption.PasswordMatch(user.Password, input.Password); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	// Generate Tokens
	access_token, err := encryption.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Path:     "/",
		MaxAge:   7 * 24,
		HTTPOnly: true,
		Secure:   true,
		Domain:   "localhost",
	})

	return c.JSON(fiber.Map{"status": "success"})
}

func (h *UserHandler) LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)

	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})

	return c.Status(200).JSON(fiber.Map{"status": "success"})
}
