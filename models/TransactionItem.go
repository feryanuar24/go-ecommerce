package models

type TransactionItem struct {
	ID            uint        `gorm:"primaryKey"`
	TransactionID uint        `json:"transaction_id"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
	ProductID     uint        `json:"product_id"`
	Product       Product     `json:"product" gorm:"foreignKey:ProductID"`
	Quantity      uint        `json:"quantity"`
	Subtotal      float64     `json:"subtotal"`
}
