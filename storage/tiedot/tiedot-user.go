package tiedot

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/HouzuoGuo/tiedot/dberr"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/examples/model"
)

// NewUserStorage initializes the storage
func NewUserStorage() *UserStorage {
	colName := "Users"
	if !doesColExsit(colName) {
		myDB.Create(colName)
	}
	return &UserStorage{myDB.Use(colName)}
}

// UserStorage stores all users
type UserStorage struct {
	users *db.Col
}

// GetAll returns the user map (because we need the ID as key too)
func (s UserStorage) GetAll() *db.Col {
	return s.users
}

// GetOne user
func (s UserStorage) GetOne(id string) (model.User, error) {
	user, err := s.GetOne(id)
	if err != nil {
		errMessage := fmt.Sprintf("User for id %s not found", id)
		return model.User{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
	}
	return user, nil
}

// Insert a user
func (s *UserStorage) Insert(c model.User) string {
	id := s.Insert(c)
	return id
}

// Delete one :(
func (s *UserStorage) Delete(id string) error {
	err := s.Delete(id)
	if err != nil {
		switch dberr.Type(err) {
		case dberr.ErrorNoDoc:
			fmt.Println("The document was already deleted")
		default:
			panic(err)
		}
	}
	return nil
}

// Update a user
func (s *UserStorage) Update(c model.User) error {
	if err := s.Update(c); err != nil {
		panic(err)
	}
	return nil
}
