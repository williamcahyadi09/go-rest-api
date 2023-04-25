package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST        string
	DB_PORT        string
	DB_NAME        string
	DB_USER        string
	DB_PASS        string
	MIGRATION_PATH string
	HTTP_ADDR      string
}

var config *Config
var once sync.Once

func loadEnv(env_path string) {
	err := godotenv.Load(env_path)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func GetConfig() Config {

	once.Do(func() {
		loadEnv(".env")

		config = &Config{
			DB_HOST:        os.Getenv("DB_HOST"),
			DB_PORT:        os.Getenv("DB_PORT"),
			DB_NAME:        os.Getenv("DB_NAME"),
			DB_USER:        os.Getenv("DB_USER"),
			DB_PASS:        os.Getenv("DB_PASS"),
			MIGRATION_PATH: os.Getenv("MIGRATION_PATH"),
			HTTP_ADDR:      os.Getenv("HTTP_ADDR"),
		}
	})

	return *config
}
