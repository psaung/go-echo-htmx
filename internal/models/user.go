package models

type User struct {
	ID       *string `json:"id"       gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email    string  `json:"email"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Address  string  `json:"address"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`

	// Orders []Order `gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "users"
}
