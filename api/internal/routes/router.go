package routes

import (
	"bri-edc/api/internal/injector"
	"bri-edc/api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRouter(app *fiber.App, ct *injector.AppContainer) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return utils.ResponseSuccess(c, "BRI-EDC API up and running", nil)
	})

	api := app.Group("/api")

	//tx := api.Group("/transactions", middleware.MustBeAuthenticated(ct))
	tx := api.Group("/transactions")
	{
		tx.Post("/sale", ct.TransactionHandler.Sale)
		tx.Post("/settlement", ct.TransactionHandler.Settlement)
	}

	app.Use(func(c *fiber.Ctx) error {
		return utils.ResponseFailed(c, http.StatusNotFound, "Endpoint not found")
	})
}
