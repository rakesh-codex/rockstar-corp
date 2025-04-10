package config

import (
	"os"
)

type Config struct {
	HTTP     HTTPConfig
	Database DBConfig
}

type HTTPConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() (*Config, error) {
	cfg := &Config{
		HTTP: HTTPConfig{
			Port: getEnv("HTTP_PORT", "8080"),
		},
		Database: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "rockstar"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
