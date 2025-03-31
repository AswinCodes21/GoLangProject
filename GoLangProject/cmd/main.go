package main

import (
	"log"
	"my_project/internal/router"
	"my_project/pkg/db"
	config "my_project/pkg/env"
	"net/http"
)

func main() {
	config.LoadEnv()
	db.ConnectDB()

	r := router.SetupRouter(db.DB)

	port := "8080"
	log.Println("Server started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
