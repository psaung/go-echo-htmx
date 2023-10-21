package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/psaung/go-echo-htmx/internal/helpers"
)

type Controllers interface {
	// Get Session storage
	GetSession() helpers.SessionStore

	// user
	RenderGetInfoHandler(c echo.Context) error

	// auth
	RenderLoginHandler(c echo.Context) error
	RenderRegisterHandler(c echo.Context) error
	RenderUnauthorizeHandler(c echo.Context) error

	// home
	RenderHomeHandler(c echo.Context) error
	RenderAboutHandler(c echo.Context) error
	RenderNotFoundHandler(c echo.Context) error
}

type controllers struct {
	Db      *gorm.DB
	Session helpers.SessionStore
}

type NewHandlersInput struct {
	Db *gorm.DB
}

func (c *controllers) GetSession() helpers.SessionStore {
	return c.Session
}

func NewControllers(input NewHandlersInput) Controllers {
	sessionStore := helpers.NewCookieSessionStore("session", "secret", time.Hour)
	return &controllers{
		Db:      input.Db,
		Session: sessionStore,
	}
}
