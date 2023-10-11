package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func NewPostgres() *gorm.DB {
	driver := os.Getenv("DATABASE_DRIVER")
	database := os.Getenv("DATABASE_URL")

	var err error

	DB, err = gorm.Open(driver, database)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Database Connected")
	return DB
}
