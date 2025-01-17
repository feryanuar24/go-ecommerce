package models

import "time"

type Transaction struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"`
	CreatedAt time.Time
}
