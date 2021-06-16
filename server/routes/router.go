package routes

import (
	"github.com/Paramosch/predictstock-backend-eng/server/controllers"
	"github.com/gofiber/fiber/v2"
)

// Function define all web-routes
func Define(app *fiber.App) {
	app.Post("/upload", controllers.UploadSharesCSV)
	app.Post("/robokassa/result", controllers.ResultInvoice)

	app.Get("/", controllers.IndexPage)

	api := app.Group("/api")
	api.Get("/users", controllers.GetAllUsers)
	api.Get("/users/:userId", func(c *fiber.Ctx) error {
		return controllers.GetTelegramId(c, c.Params("userId"))
	})
	api.Get("/users/:userId/history", func(c *fiber.Ctx) error {
		return controllers.GetUserHistory(c, c.Params("userId"))
	})

}
