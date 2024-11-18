package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", showIndexPage)

	router.GET("/person/view/:person_id", getPerson)

	router.Run()
}
