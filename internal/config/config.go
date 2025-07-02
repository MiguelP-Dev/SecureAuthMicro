package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Config is a struct that contains the configuration for the application
type Config struct {
	ServerPort         string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBSSLMode          string
	JWTSecret          string
	JWTExpiration      time.Duration
	CSRFSecret         string
	CORSAllowedOrigins []string
}

// LoadConfig loads the configuration for the application
func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse JWT expiration
	jwtExpiration, err := time.ParseDuration(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		jwtExpiration = 24 * time.Hour
	}

	// Parse CORS Origins
	var corsOrigins []string
	if origins := os.Getenv("CORS_ALLOWED_ORIGINS"); origins != "" {
		corsOrigins = strings.Split(origins, ",")
	}

	// Return the configuration
	return &Config{
		ServerPort:         os.Getenv("SERVER_PORT"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		DBSSLMode:          os.Getenv("DB_SSLMODE"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		JWTExpiration:      jwtExpiration,
		CSRFSecret:         os.Getenv("CSRF_SECRET"),
		CORSAllowedOrigins: corsOrigins,
	}
}
