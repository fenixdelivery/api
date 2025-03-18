package api

import (
	fiber "github.com/gofiber/fiber/v2"
	routes "api-fenix/internal/api/routes"
	db "api-fenix/internal/database"
	middlewares "api-fenix/internal/api/middlewares"
	gorm "gorm.io/gorm"
)

var App = fiber.New()

func init() {
	db, _ := db.Connect()

	App.Get("/live", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	App.Use(middlewares.ExtendsContext[*gorm.DB](db, "db"))

	routes.GenerateRoutes(App)
}

