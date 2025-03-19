package api

import (
	middlewares "api-fenix/internal/api/middlewares"
	routes "api-fenix/internal/api/routes"
	database "api-fenix/internal/database"

	fiber "github.com/gofiber/fiber/v2"
	gorm "gorm.io/gorm"
)

var App = fiber.New()

func init() {
	db, _ := database.Connect()
	database.ChargeEntities(db)

	App.Get("/live", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	App.Use(middlewares.ExtendsContext[*gorm.DB](db, "db"))

	routes.GenerateRoutes(App)
}
