package dto

type ClientDTO struct {
	Email    string `json:"email" validation:"required,email"`
	Address  string `json:"address"`
	Phone    string `json:"phone" validation:"required"`
	Name     string `json:"name" validation:"required""`
	Lastname string `json:"lastname" validation:"required"`
	Dni      string `json:"dni" validation:"required"`
}
