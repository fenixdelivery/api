package routes

import (
	"fmt"

	dto "api-fenix/internal/api/dto"

	database "api-fenix/internal/database/services"

	fiber "github.com/gofiber/fiber/v2"
	gorm "gorm.io/gorm"
)

func ClientRoutes(app *fiber.App) {
	app.Post("/client", func(ctx *fiber.Ctx) error {
		db, ok := ctx.Locals("db").(*gorm.DB)

		var body dto.ClientDTO

		fmt.Println(db)

		if err := ctx.BodyParser(&body); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "Error en el formato del cuerpo",
				"error":   err.Error(),
			})
		}

		fmt.Println(body)

		if !ok {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "error",
				"code":  500,
			})
		}

		client, err := database.CreateClient(db, body)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "error",
				"code":  500,
			})
		}

		return ctx.JSON(fiber.Map{
			"data": fiber.Map{
				"client": client,
			},
			"message": "test",
		})
	})
}
