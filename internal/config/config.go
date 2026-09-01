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
	err := godotenv.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			return Config{}, fmt.Errorf("load .env: %w", err)
		}
	}

	serverPortStr, err := getEnv("SERVER_PORT")
	if err != nil {
		return Config{}, err
	}

	port, err := strconv.Atoi(serverPortStr)
	if err != nil {
		return Config{}, fmt.Errorf("parse SERVER_PORT: %w", err)
	}

	host, err := getEnv("DB_HOST")
	if err != nil {
		return Config{}, err
	}

	dbPortStr, err := getEnv("DB_PORT")
	if err != nil {
		return Config{}, err
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return Config{}, fmt.Errorf("parse DB_PORT: %w", err)
	}
	user, err := getEnv("DB_USER")
	if err != nil {
		return Config{}, err
	}

	password, err := getEnv("DB_PASSWORD")
	if err != nil {
		return Config{}, err
	}

	name, err := getEnv("DB_NAME")
	if err != nil {
		return Config{}, err
	}

	return Config{
		Server: ServerConfig{
			Port: port,
		},
		Database: DatabaseConfig{
			Host:     host,
			Port:     dbPort,
			User:     user,
			Password: password,
			Name:     name,
		},
	}, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("environment variable %s is required", key)
	}

	return value, nil
}
