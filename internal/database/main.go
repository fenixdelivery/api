package database

import (
	"fmt"
	"os"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

	entities "api-fenix/internal/database/models"
)

func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
    
	if err != nil {
            log.Fatal("Error al cargar el archivo .env")
    }

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("[db] Error init postgres")
		return nil, err
	}

	fmt.Println("[db] Start Postgres")

	return db, nil
}

func ChargeEntities(db *gorm.DB) {
	db.AutoMigrate(&entities.Client{})
}
