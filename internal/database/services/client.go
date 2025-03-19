package services

import (
	"fmt"
	// "strings"
	"reflect"

	dto "api-fenix/internal/api/dto"
	entities "api-fenix/internal/database/models"

	gorm "gorm.io/gorm"
)

type Filters struct {
	name string
	dni string
}

func CreateClient(db *gorm.DB, clientDto dto.ClientDTO) (entities.Client, error) {
	client := entities.Client{
		Name:     clientDto.Name,
		Address:  clientDto.Address,
		Email:    clientDto.Email,
		Phone:    clientDto.Phone,
		Lastname: clientDto.Lastname,
		Dni:      clientDto.Dni,
		DisplayName: fmt.Sprintf("%s %s", clientDto.Name, clientDto.Lastname),
	}

	result := db.Create(&client)

	if result.Error != nil {
		return entities.Client{}, result.Error
	}

	return client, nil
}

func ReadClients(db *gorm.DB, filters dto.ClientDTO, page int) ([]entities.Client, error) {
	// var builder strings.Builder
	clients := []entities.Client{}

	rv := reflect.ValueOf(filters)
    rt := reflect.TypeOf(filters)

	offset := (page - 1) * 40

	for i := 0; i < rt.NumField(); i++ {
        field := rt.Field(i)
        value := rv.Field(i)
		
		fmt.Println(field)
		fmt.Println(value)
    }

	db.Limit(40).Offset(offset).Find(&clients)

	fmt.Println(clients)

	return clients, nil
}
