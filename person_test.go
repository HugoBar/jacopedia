package main

import (
	"testing"
)

func TestGetPersonByID(t *testing.T) {
	id := 2
	expectedName := "Bob"
	person, _ := getPersonByID(id)

	if person.Name != expectedName {
		t.Fail()
	}
}
