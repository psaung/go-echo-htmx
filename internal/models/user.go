package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        *string        `json:"id"        gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Status    *string        `json:"status"`

	Orders []Order `gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "users"
}
