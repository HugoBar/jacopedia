package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPeople(c *gin.Context) {
	people, err := getAllPeople()
	if err != nil {
		// If there's an error, return a 500 internal server error with a message
		log.Println("Error fetching people:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve people",
		})
		return
	}

	// Return the list of people as a JSON response
	c.JSON(http.StatusOK, people)
}

func getPerson(c *gin.Context) {
	if personID, err := strconv.Atoi(c.Param("person_id")); err == nil {
		// Check if the person exists
		if person, err := getPersonByID(personID); err == nil {
			c.JSON(http.StatusOK, person)
		} else {
			// If the person is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid person ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createPerson(c *gin.Context) {
	var newPerson person

	// Call BindJSON to bind the received JSON to
	// newPerson.
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	createdPerson, err := addPerson(newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert person into database"})
		return
	}

	// Respond with the created person
	c.IndentedJSON(http.StatusCreated, createdPerson)
}
