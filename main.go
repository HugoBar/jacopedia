package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"jacopedia/database"
)

var router *gin.Engine

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.InitDB()

	// Ensure the database connection is closed on exit
	defer database.CloseDB()

	// Application logic here
	log.Println("Application started")

	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.GET("/person/all", getPeople)
	router.GET("/person/:person_id", getPerson)
	router.POST("/person", createPerson)

	router.Run()
}
