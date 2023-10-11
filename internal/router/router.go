package router

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/controllers"
	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/middlewares"
)

func Init(e *echo.Echo) {
	e.GET("/", controllers.HomeHandler)
	e.GET("/about", controllers.AboutHandler)
	e.GET("/user-info", controllers.GetInfoHandler)
	e.GET("/404", controllers.NotFoundHandler)
	e.GET("/login", controllers.LoginHandler)

	sessionStore := helpers.NewCookieSessionStore("session", "secret", time.Hour)

	app := e.Group("/app", middlewares.AuthMiddleware(sessionStore))
	{
		app.GET("/user-info", controllers.GetInfoHandler)
	}
}
