package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port        string
	FrontendURL string
	Environment string
}

type JWTConfig struct {
	Secret string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file (only works locally, not on Render)
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found")
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("PGHOST", "localhost"),
			Port:     getEnv("PGPORT", "5432"),
			User:     getEnv("PGUSER", "postgres"),
			Password: getEnv("PGPASSWORD", "password"),
			DBName:   getEnv("PGDATABASE", "medicare"),
			SSLMode:  getEnv("PGSSLMODE", "disable"),
		},
		Server: ServerConfig{
			Port:        getEnv("PORT", "8080"),
			FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),
			Environment: getEnv("ENV", "development"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", ""),
		},
	}

	// Validate required fields
	if config.JWT.Secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return config, nil
}

// GetDatabaseURL returns PostgreSQL connection string
func (c *Config) GetDatabaseURL() string {
	// First check for DATABASE_URL (for production/Render/Supabase)
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return dbURL
	}

	// Fallback to individual environment variables (for local development)
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
