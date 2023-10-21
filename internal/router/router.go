package router

import (
	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/controllers"
	"github.com/psaung/go-echo-htmx/internal/middlewares"
)

func Init(e *echo.Echo, c controllers.Controllers) {
	// pages
	e.GET("/", c.RenderHomeHandler)
	e.GET("/about", c.RenderAboutHandler)
	e.GET("/404", c.RenderNotFoundHandler)
	e.GET("/login", c.RenderLoginHandler)
	e.GET("/register", c.RenderRegisterHandler)

	// htmx specific route
	e.GET("/htmx/unauthorize", c.RenderUnauthorizeHandler)

	// sensitive routes
	app := e.Group("/app", middlewares.AuthMiddleware(c.GetSession()))
	{
		app.GET("/user-info", c.RenderGetInfoHandler)
	}
}
