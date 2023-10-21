package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
