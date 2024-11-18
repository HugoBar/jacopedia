package main

import "errors"

type person struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Age              int    `json:"age"`
	Birthday         string `json:"birthday"`
	ProfilePictureID string `json:"profilePicture"`
	Title            string `json:"title"`
}

var peopleList = []person{
	{ID: 1, Name: "Alice", Age: 25, Birthday: "1999-05-15", ProfilePictureID: "placeholder", Title: "Founder"},
	{ID: 2, Name: "Bob", Age: 30, Birthday: "1994-08-22", ProfilePictureID: "placeholder", Title: "Employee"},
}

func getAllPeople() []person {
	return peopleList
}

func getPersonByID(id int) (*person, error) {
	for _, p := range peopleList {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("person not found")
}
