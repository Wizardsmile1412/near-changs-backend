package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	cfg := &Config {
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        ServerPort: os.Getenv("SERVER_PORT"),
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = "8081"
	}

	return cfg, nil
}