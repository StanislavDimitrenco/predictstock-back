package server

import (
	"context"
	"github.com/Paramosch/predictstock-backend-eng/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func Run(ctx context.Context) context.Context {
	app := fiber.New()
	// define routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,X-Requested-With, Content-Type, Accept",
	}))

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("ctx", ctx)
		return c.Next()
	})

	routes.Define(app)

	// start web-server
	go func(app *fiber.App) {
		port := os.Getenv("FIBER_PORT")
		err := app.Listen(":" + port)
		if err != nil {
			panic(err)
		}
	}(app)

	return context.WithValue(ctx, "webserver", app)

}
