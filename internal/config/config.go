package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application-wide configuration.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig contains HTTP server settings
type ServerConfig struct {
	Port int
}

// DatabaseConfig contains PostgreSQL connection settings
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return Config{}, fmt.Errorf("parse SERVER_PORT: %w", err)
	}

	host := os.Getenv("DB_HOST")
	return Config{
		Server: ServerConfig{
			Port: port,
		},
		Database: DatabaseConfig{
			Host: host,
		},
	}, nil
}
