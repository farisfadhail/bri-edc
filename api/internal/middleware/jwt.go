package middleware

import (
	"bri-edc/api/internal/injector"
	"bri-edc/api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func extractAndParseToken(c *fiber.Ctx) (*string, error) {
	auth := c.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		return nil, utils.ResponseFailed(c, fiber.StatusUnauthorized, "Unauthorized: missing Bearer token")
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	username, err := utils.ParseJWT(token)
	if err != nil {
		return nil, utils.ResponseFailed(c, fiber.StatusUnauthorized, err.Error())
	}

	return username, nil
}

func MustBeAuthenticated(ct *injector.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username, err := extractAndParseToken(c)
		if err != nil {
			return err
		}

		c.Locals("JWT_USERNAME", &username)

		return c.Next()
	}
}
