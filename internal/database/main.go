package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	entities "api-fenix/internal/database/models"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=admin dbname=fenix port=5432 sslmode=disable TimeZone=Europe/Madrid"
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
