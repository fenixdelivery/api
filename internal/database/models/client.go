package models

import (
	gorm "gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Email       string `gorm:"size:50;not null;unique"`
	Address     string `gorm:"size:80"`
	Phone       string `gorm:"size:15;unique;not null"`
	Name        string `gorm:"size:40;not null"`
	Lastname    string `gorm:"size:40;not null"`
	Dni         string `gorm:"size:15;not null;unique"`
	Active      bool   `gorm:"default:true"`
	DisplayName string `gorm:"size:90"`
}
