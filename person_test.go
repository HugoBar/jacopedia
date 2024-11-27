package main

import (
	"testing"
)

// Test the function that fetches all articles
func TestGetAllPeople(t *testing.T) {
	alist := getAllPeople()

	// Check that the length of the list of people returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(peopleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Name != peopleList[i].Name ||
			v.ID != peopleList[i].ID ||
			v.Title != peopleList[i].Title {

			t.Fail()
			break
		}
	}
}

func TestGetPersonByID(t *testing.T) {
	id := 2
	expectedName := "Bob"
	person, _ := getPersonByID(id)

	if person.Name != expectedName {
		t.Fail()
	}
}
