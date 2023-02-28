package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techwithmat/kanbi/pkg/encryption"
)

func AuthMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")

	if len(accessToken) == 0 {
		return c.Status(401).JSON(fiber.Map{
			"status":  "fail",
			"message": "Unauthorized",
		})
	}

	claims, err := encryption.VerifyAccessToken(accessToken)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status":  "fail",
			"message": "Unauthorized",
		})
	}

	userId := claims["sub"].(string)
	c.Locals("user_id", userId)
	c.Locals("expires", claims["exp"])

	return c.Next()
}
