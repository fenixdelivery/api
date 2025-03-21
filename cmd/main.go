package main

import (
	"os"
	"log"
	api "api-fenix/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	app := api.App

	err := godotenv.Load()
    
	if err != nil {
            log.Fatal("Error al cargar el archivo .env")
    }


	app.Listen(":" + os.Getenv("PORT"))
}