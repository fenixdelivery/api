package models

import (
	gorm "gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Company  string `json:"company"`
	Email    string `json:"email"`
	address  string `json:"address"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Description string `json:"description"`
	Nif string `json:"nif"`
	Web string `json:"web"`
	Active bool `json:"active"`
	DisplayName string `json:"displayName"`
}
