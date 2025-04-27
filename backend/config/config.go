package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db               DBConfig
	Port             string
	Auth             AuthConfig
	PdfInvoiceFolder string
}

type DBConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
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
		Db: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
		Port:             os.Getenv("PORT"),
		PdfInvoiceFolder: os.Getenv("INVOICE_PDF_FOLDER"),
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}

}
