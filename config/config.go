package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ENV      string
	DBConfig dbConfig
}

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

func initConfig() AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Default().Println("Error loading .env file")
	}

	return AppConfig{
		ENV: getENV("ENV_MODE", "production"),
		DBConfig: dbConfig{
			Host:     getENV("DB_HOST", ""),
			User:     getENV("DB_USER", ""),
			Password: getENV("DB_PASSWORD", ""),
			DBName:   getENV("DB_NAME", ""),
			Port:     getENV("DB_PORT", "5432"),
		},
	}
}

var Config = initConfig()
