package migrations

import (
	"log"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/psaung/go-echo-htmx/internal/config/database"
)

type users struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;uniqueIndex"`
	Name      string
	Password  string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (users) TableName() string {
	return "users"
}

func GetMigrationEngine() *gormigrate.Gormigrate {
	var err error
	postgresURL := database.GetPostgresURL()
	var db *gorm.DB
	dialector := postgres.New(postgres.Config{
		DSN: postgresURL,
	})
	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrateOpt := gormigrate.DefaultOptions
	migrateOpt.TableName = "migrations"
	migrateOpt.ValidateUnknownMigrations = true
	return gormigrate.New(db, migrateOpt, []*gormigrate.Migration{{
		ID: "201608301400",
		Migrate: func(tx *gorm.DB) error {
			err = tx.Migrator().CreateTable(&users{})
			if err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Migrator().DropTable("users")
			if err != nil {
				return err
			}
			return nil
		},
	}})
}
