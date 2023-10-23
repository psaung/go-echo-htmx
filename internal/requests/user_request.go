package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserAuth struct {
	Email    string `form:"email" validate:"required" example:"john.doe@example.com"`
	Password string `form:"password" validate:"required" example: "111111"`
}

type RegisterRequest struct {
	UserAuth
	Name            string `form:"name" validate:"required" example: "name"`
	ConfirmPassword string `form:"confirm" validate:"required" example: "111111"`
	Address         string `form:"address" example: "somewhere"`
}

type LoginRequest struct {
	UserAuth
}

func (auth UserAuth) Validate() error {
	return validation.ValidateStruct(
		&auth,
		validation.Field(&auth.Email, is.Email),
	)
}

func (req RegisterRequest) Validate() error {
	err := req.UserAuth.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(
		&req,
	)
}
