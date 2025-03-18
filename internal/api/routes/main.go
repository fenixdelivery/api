package routes

import (
	fiber "github.com/gofiber/fiber/v2"
)

func GenerateRoutes(app *fiber.App) {
	ClientRoutes(app)
}