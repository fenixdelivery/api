package services

import (
	dto "api-fenix/internal/api/dto"
	entities "api-fenix/internal/database/models"

	gorm "gorm.io/gorm"
)

func CreateClient(db *gorm.DB, clientDto dto.ClientDTO) (entities.Client, error) {
	client := entities.Client{
		Name:     clientDto.Name,
		Address:  clientDto.Address,
		Email:    clientDto.Email,
		Phone:    clientDto.Phone,
		Lastname: clientDto.Lastname,
		Dni:      clientDto.Dni,
	}

	result := db.Create(&client)

	if result.Error != nil {
		return entities.Client{}, result.Error
	}

	return client, nil
}
