package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func (ws *webServer) serverHome(c *gin.Context) {
	data := gin.H{
		"title": "Home",
	}
	c.HTML(http.StatusOK, "home/index.tmpl", data)
}

func (ws *webServer) serverBat(c *gin.Context) {
	data := gin.H{
		"title": "Bat",
	}
	c.HTML(http.StatusOK, "home/index.tmpl", data)
}
