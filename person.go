package main

import (
	"errors"
	"jacopedia/database"
)

type person struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Age              int    `json:"age"`
	Birthday         string `json:"birthday"`
	ProfilePictureID string `json:"profile_picture_id"`
	Title            string `json:"title"`
}

var peopleList = []person{
	{ID: 1, Name: "Alice", Age: 25, Birthday: "1999-05-15", ProfilePictureID: "placeholder", Title: "Founder"},
	{ID: 2, Name: "Bob", Age: 30, Birthday: "1994-08-22", ProfilePictureID: "placeholder", Title: "Employee"},
}

func getAllPeople() ([]person, error) {
	// Query to retrieve all people
	query := `SELECT id, name, age, birthday, profile_picture_id, title FROM people`

	// Execute the query
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold the retrieved people
	var people []person

	// Iterate over rows and scan the data into the slice
	for rows.Next() {
		var p person
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Birthday, &p.ProfilePictureID, &p.Title); err != nil {
			return nil, err
		}
		people = append(people, p)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return people, nil
}

func getPersonByID(id int) (*person, error) {
	for _, p := range peopleList {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("person not found")
}
