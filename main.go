package main

import (
	"log"

	"github.com/Zain0205/gdgoc-subbmission-be-go/config"
	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/routes"
)

func main() {
	config.LoadConfig()
	database.ConnectDatabase()
	database.MigrateDatabase()

	r := routes.SetupRouter()

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

