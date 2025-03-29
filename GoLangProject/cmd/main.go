package main

import (
	"log"
	"my_project/pkg/db"
	"my_project/pkg/env"
)

func main() {
	env.LoadEnv()
	db.ConnectDB()

	log.Println("Server is running...")
}
