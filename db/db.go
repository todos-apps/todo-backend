package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"todo-backend/config"
	"todo-backend/models"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
    dsn := cfg.DSN()
    log.Println("🔗 Using DSN:", dsn)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

func AutoMigrate(db *gorm.DB) error {
    // migrate models
    return db.AutoMigrate(&models.Todo{})
}
