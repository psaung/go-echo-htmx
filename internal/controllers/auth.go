package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *controllers) LoginHandler(ec echo.Context) error {
	return ec.Render(http.StatusOK, "htmx/login", nil)
}

func (c *controllers) RegisterHandler(ec echo.Context) error {
	return ec.Render(http.StatusOK, "htmx/register", nil)
}
