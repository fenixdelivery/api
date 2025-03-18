package services

import (
	entities "api-fenix/internal/database/models"
	gorm "gorm.io/gorm"
)

func CreateClient(db *gorm.DB, client entities.Client) (entities.Client, error) {
	result := db.Create(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	return result, nil
}