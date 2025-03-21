package services

import (
	"fmt"
	"strings"
	"reflect"

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
		Type:     clientDto.Type,
		DisplayName: fmt.Sprintf("%s %s", clientDto.Name, clientDto.Lastname),
	}

	result := db.Create(&client)

	if result.Error != nil {
		return entities.Client{}, result.Error
	}

	return client, nil
}

func ReadClients(db *gorm.DB, filters dto.ClientDTO, page int) ([]entities.Client, error) {
	var builder strings.Builder
	clients := []entities.Client{}

	rv := reflect.ValueOf(filters)
    rt := reflect.TypeOf(filters)

	offset := (page - 1) * 40

	for i := 0; i < rt.NumField(); i++ {
        value := rv.Field(i).Interface().(string)

		if strings.Compare(value, "") == 0 {
			continue
		}

		field := rt.Field(i).Tag.Get("json")
		
		builder.WriteString(field)
		builder.WriteString(" = ")
		builder.WriteString(value)
		builder.WriteString(" AND ")
    }

	where := builder.String()
	
	if strings.Compare(where, "") != 0 {
		db.Limit(40).Offset(offset).Where(where[:len(where) - 5]).Find(&clients)
		return clients, nil
	}

	db.Limit(40).Offset(offset).Find(&clients)

	return clients, nil
}
