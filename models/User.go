package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"-" validate:"required,min=6"`
	Role     string `json:"role" validate:"required"`
}

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
