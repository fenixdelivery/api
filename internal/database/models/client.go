package models

import (
	gorm "gorm.io/gorm"
	"database/sql/driver"
)

type TypeClient string

const (
  Restaurant     TypeClient = "restaurant"
  Rider          TypeClient = "rider"
)

func (self *TypeClient) Scan(value interface{}) error {
    *self = TypeClient(value.([]byte))
    return nil
}

func (self TypeClient) Value() (driver.Value, error) {
    return string(self), nil
}

type Client struct {
	gorm.Model
	Email       string     `gorm:"size:50;not null;unique"`
	Address     string     `gorm:"size:80"`
	Phone       string     `gorm:"size:15;unique;not null"`
	Name        string     `gorm:"size:40;not null"`
	Lastname    string     `gorm:"size:40;not null"`
	Active      bool       `gorm:"default:true"`
	DisplayName string     `gorm:"size:90"`
	Type 	    string     `gorm:"size:12;not null"`
}
