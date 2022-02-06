package dto

type BookUpdateDTO struct {
	ID uint16 `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID uint16 `json:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID uint16 `json:"user_id,omitempty"`
}