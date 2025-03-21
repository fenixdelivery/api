package routes

import (
	"strconv"

	dto "api-fenix/internal/api/dto"

	database "api-fenix/internal/database/services"

	fiber "github.com/gofiber/fiber/v2"
	validator "github.com/go-playground/validator/v10"
	gorm "gorm.io/gorm"
)

var validate *validator.Validate

func ClientRoutes(app *fiber.App) {
	app.Post("/client", func(ctx *fiber.Ctx) error {
		db, ok := ctx.Locals("db").(*gorm.DB)

		if !ok {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "error",
				"code":  500,
				"message": "ha ocurrido un error al crear el cliente",
				"data": fiber.Map{},
			})
		}

		var body dto.ClientDTO

		if err := ctx.BodyParser(&body); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "Error en el formato del cuerpo",
				"error":   err.Error(),
				"data": fiber.Map{},
				"code": 400,
			})
		}

		validate = validator.New()
		err := validate.Struct(body)

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "Error en el formato del cuerpo",
				"error":   err.Error(),
				"data": fiber.Map{},
				"code": 400,
			})
		}

		client, err := database.CreateClient(db, body)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": err,
				"code":  500,
				"data": fiber.Map{},
				"message": "Ha ocurrido un error al crear el cliente",
			})
		}

		return ctx.Status(201).JSON(fiber.Map{
			"data": fiber.Map{
				"client": client,
			},
			"message": "El cliente ha sido creado con exito",
			"code": 201,
			"error": nil,

		})
	})

	app.Get("/client", func(ctx *fiber.Ctx) error {
		db, ok := ctx.Locals("db").(*gorm.DB)

		if !ok {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "error",
				"code":  500,
				"message": "ha ocurrido un error al buscar los clientes",
				"data": fiber.Map{},
			})
		}

		var body dto.ClientDTO

		if err := ctx.BodyParser(&body); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "Error en el formato del cuerpo",
				"error":   err.Error(),
				"data": fiber.Map{},
				"code": 400,
			})
		}

		p := ctx.Query("page")
		page, _ := strconv.Atoi(p)

		clients, err := database.ReadClients(db, body, page)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": err,
				"code":  500,
				"data": fiber.Map{},
				"message": "Ha ocurrido un error al buscar los cliente",
			})
		}

		if len(clients) == 0 {
			return ctx.Status(404).JSON(fiber.Map{
				"data": fiber.Map{
					"clients": clients,
				},
				"message": "No se han encontrado clientes",
				"code": 404,
				"error": nil,
		})

		}

		return ctx.Status(200).JSON(fiber.Map{
			"data": fiber.Map{
				"clients": clients,
			},
			"message": "El cliente ha sido encontrado con exito",
			"code": 200,
			"error": nil,
		})
	})


}
