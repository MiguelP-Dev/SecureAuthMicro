package database

import (
	"SecureAuthMicro/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the variable returned by GetDB function
var DB *gorm.DB

// InitDB is the function in charge of starting the database
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db

	return db, nil
}

// GetDB is the function in charge of deliver the Database instance
func GetDB() *gorm.DB {
	return DB
}

// Migrate is the function in charge of migrate the models to database
func Migrate(models ...any) error {
	return DB.AutoMigrate(models...)
}
