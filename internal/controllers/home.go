package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "pages/home", map[string]interface{}{
		"Name":  "Guest!",
		"Title": "Home",
	})
}

func AboutHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "pages/about", map[string]interface{}{
		"Title": "About",
	})
}

func NotFoundHandler(c echo.Context) error {
	return c.Render(http.StatusNotFound, "pages/404", map[string]interface{}{
		"Title": "Not Found",
	})
}
