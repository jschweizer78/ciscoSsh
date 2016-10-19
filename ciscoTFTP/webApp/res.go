package main

import "github.com/gin-gonic/gin"

type userRes struct {
	name string
}

func (ur *userRes) AddOne(c *gin.Context) {

}

func (ur *userRes) GetOne(c *gin.Context) {

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

/*

AddOne(c *gin.Context)
GetOne(c *gin.Context)
GetAll(c *gin.Context)
DeleteOne(c *gin.Context)
UpdateOne(c *gin.Context)
MyName() string

*/
