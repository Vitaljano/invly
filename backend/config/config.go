package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{}
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error load .env file, using default config")
	}

	return &Config{
		Port: os.Getenv("PORT"),
	}

}
