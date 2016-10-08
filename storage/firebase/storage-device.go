package firebaseStorage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/examples/model"
)

// NewUserStorage initializes the storage
func NewUserStorage() *UserStorage {

	return &UserStorage{}
}

// UserStorage stores all users
type UserStorage struct {
	users   map[string]*model.User
	idCount int
}

// GetAll returns the user map (because we need the ID as key too)
func (s UserStorage) GetAll() map[string]*model.User {
	return s.users
}

// GetOne user
func (s UserStorage) GetOne(id string) (model.User, error) {

	errMessage := fmt.Sprintf("User for id %s not found", id)
	return model.User{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a user
func (s *UserStorage) Insert(c model.User) string {
	id := `INSERT-ID`
	return id
}

// Delete one :(
func (s *UserStorage) Delete(id string) error {

	return nil
}

// Update a user
func (s *UserStorage) Update(c model.User) error {

	return nil
}
