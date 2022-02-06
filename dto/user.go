package dto

type UserUpdateDTO struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"min=6"`
}

// type UserUpdateDTO struct {
// 	ID    uint64 `json:"id"`
// 	Name  string `json:"name" binding:"required"`
// 	Email string `json:"email" binding:"required,email"`
// 	Password string `json:"password,omitempty" binding:"min=6"`
// }

// type UserCreateDTO struct {
// 	ID    uint64 `json:"id" binding:"required"`
// 	Name  string `json:"name" binding:"required"`
// 	Email string `json:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" validate:"min:6" binding:"required"`
// }