package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *controllers) RenderGetInfoHandler(ec echo.Context) error {
	res := map[string]interface{}{
		"name":  "someone",
		"Phone": "99999",
		"Email": "someone@gmail.com",
	}
	return ec.Render(http.StatusOK, "htmx/name_card", res)
}
