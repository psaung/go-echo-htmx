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
				// TODO:redirect htmx/unauthorie route if the request is handled from the htmx specific route
				// or redirect login route if the request is handled from the other route
				return c.Redirect(http.StatusTemporaryRedirect, "/htmx/unauthorize")
			}

			c.Set("user", user)

			return next(c)
		}
	}
}
