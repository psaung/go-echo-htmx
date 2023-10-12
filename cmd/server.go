package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/psaung/go-echo-htmx/internal/config/database"
	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/router"
)

// main function
func main() {
	e := echo.New()

	e.HTTPErrorHandler = helpers.CustomHTTPErrorHandler

	db := database.NewPostgres()
	dbPool, err := db.DB()

	helpers.NewTemplateRenderer(e, "public/views/")

	router.Init(e)

	port := os.Getenv("APP_PORT")
	e.Server.Addr = ":" + port

	e.Static("/", "static")

	if err = dbPool.Close(); err != nil {
		log.Printf("Error closing db connection %s", err)
	}

	log.Printf("Server started... at PORT:%s", port)
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
