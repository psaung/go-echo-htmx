package middlewares

import (
	"encoding/base64"
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/models"
)

func SessionMiddleware(sessionStore helpers.SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionUser, err := sessionStore.Get(c, "auth")
			if err != nil {
				// TODO: error handling
				return next(c)
			}

			user := models.CookieData{}
			decodedData, err := base64.StdEncoding.DecodeString(sessionUser)
			err = json.Unmarshal([]byte(decodedData), &user)

			c.Set("user", user)

			return next(c)
		}
	}
}
