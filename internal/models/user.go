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

type CookieData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (User) TableName() string {
	return "users"
}
