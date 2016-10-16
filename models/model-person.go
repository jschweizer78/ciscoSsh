package models

import (
	"errors"
	"fmt"

	"github.com/manyminds/api2go/jsonapi"
)

// PersonName is an object repersenting a person's name
type PersonName struct {
	Given string `json:"first"`
	Sir   string `json:"last"`
}

// Person is a struct repersenting a person
type Person struct {
	ID     string     `json:"id"`
	Name   PersonName `json:"name"`
	Phone  string     `json:"phone"`
	Email  string     `json:"email"`
	exists bool
}

// NewPerson creates a new person struct
func NewPerson(given, sir string) *Person {
	personName := PersonName{given, sir}
	var person Person
	person.Name = personName
	return &person
}

// DisplayName returns a string with first and last names
func (p *Person) DisplayName() string {

	return fmt.Sprintf("%s %s", p.Name.Given, p.Name.Sir)
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (p *Person) GetID() string {
	return p.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (p *Person) SetID(id string) error {
	p.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (p *Person) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "chocolates",
			Name: "sweets",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (p *Person) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (p *Person) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	return result
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (p *Person) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "sweets" {

		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new sweets that a users loves so much
func (p *Person) AddToManyIDs(name string, IDs []string) error {
	if name == "sweets" {

		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (p *Person) DeleteToManyIDs(name string, IDs []string) error {

	return errors.New("There is no to-many relationship with the name " + name)
}
