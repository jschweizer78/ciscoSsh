package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jschweizer78/ciscoSsh/interfaces"
)

type apiResource struct {
	stores map[string]interfaces.Storager
	res    map[string]interfaces.APIResourcer
	engine *gin.Engine
}

func newAPIRes(engine *gin.Engine) *apiResource {
	return &apiResource{
		stores: make(map[string]interfaces.Storager),
		res:    make(map[string]interfaces.APIResourcer),
		engine: engine,
	}
}

func (ar *apiResource) indexPage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/public/index.html")
}

func (ar *apiResource) setResoreces() {
	rgAPI := ar.engine.Group("/api")

	for key, api := range ar.res {
		group := rgAPI.Group(fmt.Sprintf("/%s", key))
		group.GET("", api.GetAll)
		group.POST("", api.AddOne)
		group.GET("/:id", api.GetOne)
		group.PUT("/:id", api.UpdateOne)
		group.DELETE("/:id", api.DeleteOne)

	}
}

func (ar *apiResource) getResoreces() {
	// resStrArr := []string{"users"}

}
