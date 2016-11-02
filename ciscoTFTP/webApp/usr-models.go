package main

import "gopkg.in/mgo.v2/bson"

// User to model a user
type User struct {
	ID    string        `bson:"ID" json:"id,omitempty"`
	HexID bson.ObjectId `bson:"-" json:"-"`
	Name  string        `bson:"Name" json:"myName,omitempty"`
}

// GetID to conform to stroage and rest interfaces
func (usr User) GetID() string {
	return usr.ID
}

// SetID to conform to stroage and rest interfaces
func (usr User) SetID(id string) error {
	usr.ID = id
	return nil
}

// GetName to conform to stroage and rest interfaces
func (usr User) GetName() string {
	return usr.Name
}

// SetName to conform to stroage and rest interfaces
func (usr User) SetName(name string) (err error) {
	usr.Name = name
	return
}
