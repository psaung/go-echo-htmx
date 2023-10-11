package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetInfoHandler(c echo.Context) error {
	res := map[string]interface{}{
		"name":  "someone",
		"Phone": "99999",
		"Email": "someone@gmail.com",
	}
	return c.Render(http.StatusOK, "htmx/name_card", res)
}
