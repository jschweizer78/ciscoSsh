package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
