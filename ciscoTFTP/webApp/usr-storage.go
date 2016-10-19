package main

import (
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/jschweizer78/ciscoSsh/interfaces"
)

type userStorage struct {
	name string
	col  *db.Col
}

func newUserStorage(db *db.DB) (us *userStorage) {
	name := "users"
	us = &userStorage{name: name, col: db.Use(name)}
	return
}

func (us *userStorage) AddOne(c interfaces.StorageItem) (id string) {
	return
}

func (us *userStorage) GetOne(id string) (item interfaces.StorageItem, err error) {

	return
}

func (us *userStorage) GetAll() (items []interfaces.StorageItem) {

	return
}

func (us *userStorage) DeleteOne(id string) (err error) {

	return
}

func (us *userStorage) UpdateOne(c interfaces.StorageItem) (err error) {

	return
}

func (us *userStorage) MyName() (name string) {

	return
}

/*

AddOne(c interfaces.StorageItem) string
GetOne(id string) (StorageItem, error)
GetAll() (items []interfaces.StorageItem)
DeleteOne(id string) (err error)
UpdateOne(c interfaces.StorageItem) (err error)
MyName() (id string)

*/