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
	req := new(requests.LoginRequest)

	if err := ec.Bind(req); err != nil {
		return err
	}

	user := models.User{}
	userRepository := repositories.NewUserRepository(c.Db)
	err := userRepository.GetUserByEmail(&user, req.Email)
	if err != nil {
		return ec.Render(http.StatusUnauthorized, "htmx/auth_error", map[string]interface{}{
			"title":   "User is not found",
			"content": "Please try it again",
		})
	}

	return nil
}

func (c *controllers) RegisterHandler(ec echo.Context) error {
	req := new(requests.RegisterRequest)

	if err := ec.Bind(req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return ec.Render(http.StatusUnprocessableEntity, "htmx/auth_error", map[string]interface{}{
			"title":   "Unprocessible Entity",
			"content": "Required fields are emptry or invalid",
		})
	}

	user := models.User{}
	userRepository := repositories.NewUserRepository(c.Db)
	err := userRepository.GetUserByEmail(&user, req.Email)
	if err != nil {
		return ec.Render(http.StatusUnprocessableEntity, "htmx/auth_error", map[string]interface{}{
			"title":   "Unprocessible Entity",
			"content": "Please try it again",
		})
	}

	if user.ID != nil {
		return ec.Render(http.StatusUnauthorized, "htmx/auth_error", map[string]interface{}{
			"title":   "Email is already existed",
			"content": "Please try it again",
		})
	}

	err = userRepository.RegisterUser(req)

	if err != nil {
		return ec.Render(http.StatusUnprocessableEntity, "htmx/auth_error", map[string]interface{}{
			"title":   "Unprocessible Entity",
			"content": "Please try it again",
		})
	}

	return ec.Render(http.StatusOK, "htmx/ok", map[string]interface{}{
		"title":    "Reigstration complete",
		"content":  "Thanks for your registration",
		"redirect": "/login",
		"link":     "Login",
	})
}
