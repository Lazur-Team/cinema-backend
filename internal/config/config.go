package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppUrl  string
	AppPort string

	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     int

	DBSSLMode string

	LogLevel string
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env not found")
	}

	config := &Config{}

	// App
	config.AppEnv = getEnv("APP_ENV", "local")
	config.AppUrl = getEnv("APP_URL", "localhost")
	config.AppPort = getEnv("APP_PORT", "8080")
	// Database
	config.DBUser = getEnv("DATABASE_USER", "")
	config.DBPassword = getEnv("DATABASE_PASSWORD", "")
	config.DBName = getEnv("DATABASE_NAME", "")
	config.DBSSLMode = getEnv("POSTGRES_SSL_MODE", "disable")
	config.DBHost = getEnv("DATABASE_HOST", "localhost")
	config.DBPort, err = strconv.Atoi(getEnv("DATABASE_PORT", "5432"))

	if err != nil {
		panic(err)
	}

	if config.DBHost == "" || config.DBUser == "" || config.DBPassword == "" || config.DBName == "" {
		panic("DATABASE_HOST, DATABASE_USER, DATABASE_PASSWORD, DATABASE_NAME must be set")
	}

	// Logs
	config.LogLevel = getEnv("LOG_LEVEL", "info")

	return config
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
