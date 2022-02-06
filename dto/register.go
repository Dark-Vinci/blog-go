package dto

type RegisterDTO struct {
	Name string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}