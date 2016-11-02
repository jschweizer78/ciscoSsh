package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRes struct {
	name     string
	usrStore *userStorage
}

func newUserRes(store *userStorage) *userRes {
	var ur userRes
	ur.name = "users"
	ur.usrStore = store
	return &ur
}

func (ur *userRes) AddOne(c *gin.Context) {
	var usr User
	if c.BindJSON(&usr) == nil {
		id := ur.usrStore.AddOne(usr)
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func (ur *userRes) GetOne(c *gin.Context) {
	id := c.Param("id")
	// fmt.Printf("ID = %s\n", id)
	item, err := ur.usrStore.GetOne(id)
	checkErr("didn't return from storage correctly", err)

	usr, found := item.(User)
	if !found {
		fmt.Printf("This is not a user %v\n", usr)
	}

	fmt.Printf("found %s", usr.Name)
	c.JSON(http.StatusOK, usr)
}

func (ur *userRes) GetAll(c *gin.Context) {

}

func (ur *userRes) DeleteOne(c *gin.Context) {

}

func (ur *userRes) UpdateOne(c *gin.Context) {

}

func (ur *userRes) MyName() (name string) {
	name = ur.name
	return name
}

func (ur *userRes) setName(name string) {
	ur.name = name
}
