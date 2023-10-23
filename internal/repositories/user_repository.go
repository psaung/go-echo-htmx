package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/models"
	"github.com/psaung/go-echo-htmx/internal/requests"
)

type UserRepositoryQ interface {
	GetUserByEmail(user *models.User, email string) error
	RegisterUser(request *requests.RegisterRequest) error
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

func (u *UserRepository) CheckPassword(user *models.User, password string) bool {
	return helpers.DecryptPassword(user.Password, password)
}

func (u *UserRepository) RegisterUser(req *requests.RegisterRequest) error {
	encryptedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	id := uuid.NewString()

	user := &models.User{
		ID:       &id,
		Name:     req.Name,
		Email:    req.Email,
		Password: encryptedPassword,
		Address:  req.Address,
	}

	return u.DB.Create(user).Error
}
