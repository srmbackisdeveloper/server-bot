package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"server-bot/internal/models"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service interface {
	Health() map[string]string

	GetAllProducts() ([]*models.Product, error)
	GetProduct(uint) (*models.Product, error)
	AddProduct(*models.Product) error
}

type service struct {
	db *gorm.DB
}

func New() Service {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/railway", dbUser, dbPassword, dbHost, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Printf("The database connection established successfully: %s:%s", dbUser, dbHost)

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	return &service{db: db}
}
