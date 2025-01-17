package models

type ProductCategory struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name" validate:"required"`
}
