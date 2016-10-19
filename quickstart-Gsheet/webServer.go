package main

import (
	"fmt"

	"github.com/jschweizer78/ciscoSsh/interfaces"

	"github.com/gin-gonic/gin"
)

type webServer struct {
	api map[string]interfaces.APIResourcer
}

func newWebServer() webServer {
	return webServer{
		api: make(map[string]interfaces.APIResourcer),
	}
}

func (ws *webServer) run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	rgHome := router.Group("/home")
	rgHome.GET("/index", ws.serverHome)

	rgBat := router.Group("/bat")
	rgBat.GET("/index", ws.serverBat)

	rgAPI := router.Group("/api")

	for key, api := range ws.api {
		group := rgAPI.Group(fmt.Sprintf("/%s", key))
		group.GET("", api.GetAll)
		group.POST("", api.AddOne)
		group.GET("/:id", api.GetOne)
		group.PUT("/:id", api.UpdateOne)
		group.DELETE("/:id", api.DeleteOne)

	}

	router.Run(":8080")
}

func getResources() []interfaces.APIResourcer {
	// resList := make([]interfaces.APIResourcer, 4)
	return nil
}
