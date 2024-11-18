package main

import (
	"net/http"
	"strconv"

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
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the person.html template
				"person.html",
				// Pass the data that the page uses
				gin.H{
					"name":    person.Name,
					"payload": person,
				},
			)

		} else {
			// If the person is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid person ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
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
