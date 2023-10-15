package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/psaung/go-echo-htmx/internal/config/database"
	"github.com/psaung/go-echo-htmx/internal/controllers"
	"github.com/psaung/go-echo-htmx/internal/helpers"
	"github.com/psaung/go-echo-htmx/internal/router"
)

// main function
func main() {
	e := echo.New()

	// init custom http error handler
	e.HTTPErrorHandler = helpers.CustomHTTPErrorHandler

	// init db
	db := database.NewPostgres()
	dbPool, err := db.DB()
	if err != nil {
		log.Printf("Can't init Db %s", err)
	}

	// init controllers
	c := controllers.NewControllers(controllers.NewHandlersInput{
		Db: db,
	})

	// init template renderer
	helpers.NewTemplateRenderer(e, "public/views/")

	// init routes and inject controllers
	router.Init(e, c)

	// setup the application port
	port := os.Getenv("APP_PORT")
	e.Server.Addr = ":" + port

	// setup statis folder for asssets
	e.Static("/", "static")

	// close the db
	if err = dbPool.Close(); err != nil {
		log.Printf("Error closing db connection %s", err)
	}

	log.Printf("Server started... at PORT:%s", port)
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
