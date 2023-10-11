package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/router"
)

var port = ":9000"

// main function
func main() {
	e := echo.New()

	e.HTTPErrorHandler = helpers.CustomHTTPErrorHandler

	helpers.NewTemplateRenderer(e, "public/views/")

	router.Init(e)

	e.Server.Addr = port
	log.Printf("Server started... at PORT%s", port)
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
