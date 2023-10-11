package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/helpers"
)

func AuthMiddleware(sessionStore helpers.SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := sessionStore.Get(c, "user")
			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}

			c.Set("user", user)

			return next(c)
		}
	}
}
