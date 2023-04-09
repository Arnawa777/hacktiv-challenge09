package database

import (
	"challenge-08/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "admin"
	DB_NAME     = "postgres"
	DB_PORT     = 5432
	DEBUG_MODE  = true // true/false
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	if DEBUG_MODE {
		return db.Debug()
	}

	return db
}

func GetUser(userID int) (*models.User, error) {
	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// MockDB returns a mock database connection for testing purposes
func MockDB() *gorm.DB {
	// Replace the connection details with your own PostgreSQL connection details
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Set up the database schema and any other necessary configuration
	// ...

	return db
}
