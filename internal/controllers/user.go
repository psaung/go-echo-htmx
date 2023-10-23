package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/models"
)

func (c *controllers) RenderGetInfoHandler(ec echo.Context) error {
	user := ec.Get("user").(models.CookieData)
	if user.ID == "" {
		return ec.Render(http.StatusUnprocessableEntity, "htmx/auth_error", map[string]interface{}{
			"title":   "Somehting went wrong",
			"content": "Please try it again",
		})
	}

	return ec.Render(http.StatusOK, "htmx/name_card", map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
	})
}
