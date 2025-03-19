package dto

type ClientDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address"`
	Phone    string `json:"phone" validate:"required"`
	Name     string `json:"name" validate:"required""`
	Lastname string `json:"lastname" validate:"required"`
	Dni      string `json:"dni" validate:"required"`
}
