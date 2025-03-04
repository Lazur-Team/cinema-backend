package config

import (
	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	ServerPort string
	DBUser     string
	DBPass     string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() {
	viper.AutomaticEnv()

	Cfg = &Config{
		ServerPort: viper.GetString("SERVER_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPass:     viper.GetString("DB_PASS"),
		DBName:     viper.GetString("DB_NAME"),
		DBHost:     viper.GetString("DB_HOST_CONTAINER"),
		DBPort:     viper.GetString("DB_PORT_CONTAINER"),
	}
}
