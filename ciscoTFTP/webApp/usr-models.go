package main

type user struct {
	ID   string `json:"id"`
	Name string `json:"myName,omitempty"`
}

func (usr *user) GetID() string {
	return usr.ID
}

func (usr *user) SetID(id string) (err error) {
	usr.ID = id
	return
}
