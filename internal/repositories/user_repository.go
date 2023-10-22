package repositories

import (
	"gorm.io/gorm"

	"github.com/psaung/go-echo-htmx/internal/models"
)

type UserRepositoryQ interface {
	GetUserByEmail(user *models.User, email string) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) GetUserByEmail(user *models.User, email string) error {
	return u.DB.Where("email = ?", email).Find(user).Error
}
