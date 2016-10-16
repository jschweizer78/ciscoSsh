package tiedot

import (
	"encoding/json"
	"fmt"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/fatih/structs"
	"github.com/jschweizer78/ciscoSsh/models"
)

// NewUserStorage initializes the storage
func NewUserStorage(myDB *db.DB) *UserStorage {
	colName := "users"
	if !doesColExsit(colName, myDB) {
		myDB.Create(colName)
	}
	return &UserStorage{Col: myDB.Use(colName)}
}

// UserStorage stores all users
type UserStorage struct {
	Col *db.Col
}

// MyName returns the user map (because we need the ID as key too)
func (s UserStorage) MyName() string {
	return fmt.Sprint("users")
}

// GetAll returns the user map (because we need the ID as key too)
func (s UserStorage) GetAll() []models.User {
	users := make([]models.User, 10)

	s.Col.ForEachDoc(func(id int, doc []byte) bool {
		var user models.User
		err := json.Unmarshal(doc, &user)
		if err != nil {
			panic(err)
		}
		user.SetID(fmt.Sprint(id))
		users = append(users, user)
		return true
	})
	return users
}

// GetOne user
func (s UserStorage) GetOne(id string) (models.User, error) {
	var user models.User
	intID := convertStringtoID(id)
	raw, err := s.Col.Read(int(intID))
	if err != nil {
		panic(err)
	}

	err = fillStruct(&user, raw)
	if err != nil {
		panic(err)
	}

	user.SetID(fmt.Sprint(id))
	return user, nil
}

// AddOne a user
func (s *UserStorage) AddOne(c models.User) string {
	m := structs.Map(c)
	idInt, err := s.Col.Insert(m)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d", idInt)
}

// DeleteOne one :(
func (s *UserStorage) DeleteOne(id string) error {
	idInt := convertStringtoID(id)
	return s.Col.Delete(int(idInt))
}

// UpdateOne a user
func (s *UserStorage) UpdateOne(c models.User) error {
	m := structs.Map(c)
	idInt := convertStringtoID(c.GetID())
	return s.Col.Update(idInt, m)
}
