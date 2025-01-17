package models

type Product struct {
	ID                uint            `gorm:"primaryKey"`
	Name              string          `json:"name" validate:"required"`
	Description       string          `json:"description"`
	Price             float64         `json:"price" validate:"required,gt=0"`
	ProductCategoryID uint            `json:"product_category_id"`
	ProductCategory   ProductCategory `json:"product_category" gorm:"foreignKey:ProductCategoryID"`
}
