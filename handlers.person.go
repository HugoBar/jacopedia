package main

import (
	"net/http"
	"strconv"

	"jacopedia/database"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	people := getAllPeople()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": people}, "index.html")
}

func getPerson(c *gin.Context) {
	if personID, err := strconv.Atoi(c.Param("person_id")); err == nil {
		// Check if the person exists
		if person, err := getPersonByID(personID); err == nil {
			// Call the HTML method of the Context to render a template
			render(c, gin.H{
				"name":    person.Name,
				"payload": person}, "person.html")

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

	// Insert the new person into the database
	query := `
        INSERT INTO people (name, age, birthday, profile_picture_id, title)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id;
    `

	var newID int
	err := database.DB.QueryRow(query, newPerson.Name, newPerson.Age, newPerson.Birthday, newPerson.ProfilePictureID, newPerson.Title).Scan(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert person into database"})
		return
	}

	// Update newPerson with the returned ID
	newPerson.ID = newID

	c.IndentedJSON(http.StatusCreated, newPerson)
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
