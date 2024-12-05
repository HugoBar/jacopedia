package main

import (
	"database/sql"
	"jacopedia/database"
	"log"
)

type person struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Age              int    `json:"age"`
	Birthday         string `json:"birthday"`
	ProfilePictureID string `json:"profile_picture_id"`
	Title            string `json:"title"`
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

func getPersonByID(id int) (person, error) {
	var p person

	// Query to retrieve a person by id
	query := `SELECT id, name, age, birthday, profile_picture_id, title FROM people WHERE id = $1`
	err := database.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Age, &p.Birthday, &p.ProfilePictureID, &p.Title)

	if err != nil {
		if err == sql.ErrNoRows {
			return p, nil // Return an empty struct if no rows are found
		}
		log.Println("Error scanning row:", err)
		return p, err
	}

	return p, nil
}

func addPerson(newPerson person) (person, error) {
	// Insert the new person into the database
	query := `
        INSERT INTO people (name, age, birthday, profile_picture_id, title)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id;
    `

	var newID int
	err := database.DB.QueryRow(query, newPerson.Name, newPerson.Age, newPerson.Birthday, newPerson.ProfilePictureID, newPerson.Title).Scan(&newID)
	if err != nil {
		log.Println("Error adding record:", err)
		return person{}, err
	}

	// Update newPerson with the returned ID
	newPerson.ID = newID

	return newPerson, nil
}
