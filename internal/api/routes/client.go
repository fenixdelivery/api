package routes

import (
	fiber "github.com/gofiber/fiber/v2"
)

func ClientRoutes(app *fiber.App) {
	app.Post("/client", func(c *fiber.Ctx) error {
		
		return c.SendString("Client")
	})
}