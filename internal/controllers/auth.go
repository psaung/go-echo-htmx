package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/models"
	"github.com/psaung/go-echo-htmx/internal/repositories"
	"github.com/psaung/go-echo-htmx/internal/requests"
)

func (c *controllers) RenderLoginHandler(ec echo.Context) error {
	return ec.Render(http.StatusOK, "pages/login", nil)
}

func (c *controllers) RenderRegisterHandler(ec echo.Context) error {
	return ec.Render(http.StatusOK, "pages/register", nil)
}

func (c *controllers) RenderUnauthorizeHandler(ec echo.Context) error {
	return ec.Render(http.StatusOK, "htmx/unauthorize", nil)
}

func (c *controllers) LoginHandler(ec echo.Context) error {
	loginRequest := new(requests.LoginRequest)

	if err := ec.Bind(loginRequest); err != nil {
		return err
	}

	user := models.User{}
	userRepository := repositories.NewUserRepository(c.Db)
	err := userRepository.GetUserByEmail(&user, loginRequest.Email)
	if err != nil {
		return ec.Render(http.StatusUnauthorized, "htmx/auth_error", map[string]interface{}{
			"title":   "User is not found",
			"content": "Please try it again",
		})
	}

	return nil
}
