package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserAuth struct {
	Email    string `json:"email" validate:"required" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required" example: "111111"`
}

func (auth UserAuth) Validate() error {
	return validation.ValidateStruct(
		&auth,
		validation.Field(&auth.Email, is.Email),
	)
}

type LoginRequest struct {
	UserAuth
}
