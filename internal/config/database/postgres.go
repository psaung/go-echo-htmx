package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetPostgresURL() string {
	pgHost := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgHost,
		port,
		dbUser,
		dbPassword,
		dbName,
	)
}

func NewPostgres() *gorm.DB {
	var err error
	postgresURL := GetPostgresURL()
	dialector := postgres.New(postgres.Config{
		DSN: postgresURL,
	})
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected")
	return DB
}
