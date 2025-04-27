package main

import (
	"github.com/Vitaljano/invly/backend/internal/customer"
	"github.com/Vitaljano/invly/backend/internal/user"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&user.User{}, &customer.Customer{})
}
