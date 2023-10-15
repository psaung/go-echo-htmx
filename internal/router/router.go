package router

import (
	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/controllers"
	"github.com/psaung/go-echo-htmx/internal/middlewares"
)

func Init(e *echo.Echo, c controllers.Controllers) {
	e.GET("/", c.HomeHandler)
	e.GET("/about", c.AboutHandler)
	e.GET("/user-info", c.GetInfoHandler)
	e.GET("/404", c.NotFoundHandler)
	e.GET("/login", c.LoginHandler)

	app := e.Group("/app", middlewares.AuthMiddleware(c.GetSession()))
	{
		app.GET("/user-info", c.GetInfoHandler)
	}
}
