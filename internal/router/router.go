package router

import (
	"github.com/labstack/echo/v4"

	"github.com/psaung/go-echo-htmx/internal/controllers"
	"github.com/psaung/go-echo-htmx/internal/middlewares"
)

func Init(e *echo.Echo, c controllers.Controllers) {
	e.Use(middlewares.SessionMiddleware(c.GetSession()))
	// pages
	e.GET("/", c.RenderHomeHandler)
	e.GET("/about", c.RenderAboutHandler)
	e.GET("/404", c.RenderNotFoundHandler)
	e.GET("/login", c.RenderLoginHandler)
	e.GET("/register", c.RenderRegisterHandler)

	// post request
	e.POST("/login", c.LoginHandler)
	e.POST("/register", c.RegisterHandler)
	e.POST("/logout", c.LogoutHandler)

	// htmx specific route
	e.GET("/htmx/unauthorize", c.RenderUnauthorizeHandler)

	// sensitive routes
	app := e.Group("/app", middlewares.AuthMiddleware(c.GetSession()))
	{
		app.GET("/user-info", c.RenderGetInfoHandler)
	}
}
