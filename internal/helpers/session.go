package helpers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SessionStore interface {
	Set(c echo.Context, key string, value interface{}) error
	Get(c echo.Context, key string) (interface{}, error)
	Delete(c echo.Context, key string)
	SessionID() string
}

type CookieSessionStore struct {
	name   string
	secret string
	ttl    time.Duration
}

func NewCookieSessionStore(name, secret string, ttl time.Duration) *CookieSessionStore {
	return &CookieSessionStore{
		name:   name,
		secret: secret,
		ttl:    ttl,
	}
}

func (s *CookieSessionStore) Get(c echo.Context, key string) (interface{}, error) {
	cookie, err := c.Request().Cookie(s.name + "_" + key)
	if err != nil {
		return nil, err
	}

	var value interface{}
	err = json.Unmarshal([]byte(cookie.Value), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (s *CookieSessionStore) Set(c echo.Context, key string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     s.name + "_" + key,
		Value:    string(jsonValue),
		Path:     "/",
		Expires:  time.Now().Add(s.ttl),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)
	return nil
}

func (s *CookieSessionStore) Delete(c echo.Context, key string) {
	cookie := &http.Cookie{
		Name:     s.name + "_" + key,
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Second),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)
}

func (s *CookieSessionStore) SessionID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return uuid.String()
}
