package config

import (
	"dummy_api/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// "gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN_PSQL")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")

	database.AutoMigrate(&models.User{})

	DB = database
}
