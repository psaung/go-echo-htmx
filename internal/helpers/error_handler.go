package helpers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	host := c.Request().Host
	URI := c.Request().RequestURI
	qs := c.QueryString()

	c.Logger().Error(err, fmt.Sprintf(" on: %s%s%s error code: %d", host, URI, qs, code))
	if code == 404 {
		c.Redirect(http.StatusTemporaryRedirect, "/404")
	}
	c.String(code, fmt.Sprintf("error code: %d", code))
}
