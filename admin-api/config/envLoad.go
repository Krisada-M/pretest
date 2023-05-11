package config

import (
	"log"

	"github.com/joho/godotenv"
)

// EnvLoad is load .env file
func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
