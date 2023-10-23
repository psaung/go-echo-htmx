package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/models"
)

func AuthMiddleware(sessionStore helpers.SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionUser, err := sessionStore.Get(c, "auth")
			if err != nil {
				sessionStore.Delete(c, "auth")
				return c.Redirect(http.StatusTemporaryRedirect, "/htmx/unauthorize")
			}

			if err != nil {
				// TODO:redirect htmx/unauthorie route if the request is handled from the htmx specific route
				// or redirect login route if the request is handled from the other route
				return c.Redirect(http.StatusTemporaryRedirect, "/htmx/unauthorize")
			}

			user := models.CookieData{}
			decodedData, err := base64.StdEncoding.DecodeString(sessionUser)
			err = json.Unmarshal([]byte(decodedData), &user)

			c.Set("user", user)

			return next(c)
		}
	}
}
