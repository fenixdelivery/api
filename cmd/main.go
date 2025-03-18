package main

import (
	api "api-fenix/internal/api"
)

func main() {
	app := api.App

	app.Listen(":3000")
}