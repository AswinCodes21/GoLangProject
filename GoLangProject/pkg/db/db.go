package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	for i := 0; i < 5; i++ {
		DB, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			break
		}
		log.Println("Failed to connect, retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Connected to the database successfully")
}
