package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWT_SECRET string

func LoadEnv() {
	err := godotenv.Load(".env") // Load from .env file
	if err != nil {
		log.Println("Warning: No .env file found. Using system environment variables.")
	}

	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("JWT_SECRET is required but not set")
	}
}
