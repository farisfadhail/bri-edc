package routes

import (
	"bri-edc/api/internal/injector"
	"bri-edc/api/internal/middleware"
	"bri-edc/api/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRouter(app *fiber.App, ct *injector.AppContainer) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 10 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			username := c.Locals("JWT_USERNAME")
			if usernameStr, ok := username.(string); ok {
				return usernameStr
			}
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too Many Requests",
			})
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to BRI EDC API",
		})
	})

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("https://app.swaggerhub.com/apis/farisfadhail-6d6/bri-edc-api/1.0.0")
	})

	api := app.Group("/api/v1")

	api.Post("/auth/_login", ct.AuthHandler.Login)

	tx := api.Group("/transactions", middleware.MustBeAuthenticated(ct))
	{
		tx.Post("/sale", ct.TransactionHandler.Sale)
		tx.Post("/settlement", ct.TransactionHandler.Settlement)
	}

	app.Use(func(c *fiber.Ctx) error {
		return utils.ResponseFailed(c, http.StatusNotFound, "Endpoint not found")
	})
}
