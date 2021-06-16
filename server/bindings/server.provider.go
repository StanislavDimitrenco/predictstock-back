package bindings

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func Boot(ctx context.Context) context.Context {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("ctx", ctx)
		return c.Next()
	})

	return context.WithValue(ctx, "fiber", app)
}
