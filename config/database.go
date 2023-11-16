package config

import (
	"fmt"
	"os"

	"github.com/muhammadputraage/digibook-api/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	databaseHost := os.Getenv("DATABASE_HOST")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePort := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		databaseHost,
		databaseUser,
		databasePassword,
		databaseName,
		databasePort,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database!")
	}

	err = DB.AutoMigrate(&book.Entity{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
