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
	GetInfoHandler(c echo.Context) error

	// auth
	LoginHandler(c echo.Context) error
	RegisterHandler(c echo.Context) error

	// home
	HomeHandler(c echo.Context) error
	AboutHandler(c echo.Context) error
	NotFoundHandler(c echo.Context) error
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
