package models

import "time"

type Order struct {
	ID        *string   `json:"id"        gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    string    `json:"userId"`
}

func (Order) TableName() string {
	return "orders"
}
