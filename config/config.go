package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string
    Port       string
}

func LoadConfigFromEnv() (*Config, error) {
    c := &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "admin"),
        DBName:     getEnv("DB_NAME", "todo"),
        DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
        Port:       getEnv("PORT", "8080"),
    }

    if c.DBUser == "" || c.DBPassword == "" || c.DBName == "" {
        return nil, errors.New("DB_USER, DB_PASSWORD and DB_NAME must be set")
    }
    return c, nil
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}

func (c *Config) DSN() string {
    // return a DSN suitable for gorm postgres driver
    // postgres://user:pass@host:port/dbname?sslmode=disable
    return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        c.DBHost, c.DBPort, c.DBUser, c.DBName, c.DBPassword, c.DBSSLMode)
}
