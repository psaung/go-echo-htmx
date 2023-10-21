package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *controllers) RenderHomeHandler(ec echo.Context) error {
	return ec.Render(http.StatusOK, "pages/home", map[string]interface{}{
		"Name":  "Guest!",
		"Title": "Home",
	})
}

func (c *controllers) RenderAboutHandler(ec echo.Context) error {
	return ec.Render(http.StatusOK, "pages/about", map[string]interface{}{
		"Title": "About",
	})
}

func (c *controllers) RenderNotFoundHandler(ec echo.Context) error {
	return ec.Render(http.StatusNotFound, "pages/404", map[string]interface{}{
		"Title": "Not Found",
	})
}
