package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "pages/login", nil)
}
