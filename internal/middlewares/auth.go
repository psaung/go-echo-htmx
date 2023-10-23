package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/models"
)

func AuthMiddleware(sessionStore helpers.SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			contextUser := c.Get("user")

			if contextUser == nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/htmx/unauthorize")
			}

			user := contextUser.(models.CookieData)

			if user.ID == "" {
				// TODO:redirect htmx/unauthorie route if the request is handled from the htmx specific route
				// 	// or redirect login route if the request is handled from the other route
				return c.Redirect(http.StatusTemporaryRedirect, "/htmx/unauthorize")
			}

			return next(c)
		}
	}
}
