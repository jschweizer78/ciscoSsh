package main

import "github.com/jschweizer78/ciscoSsh/service"

func addRoutes(ds *service.Service) {
	deviceRoutes := ds.Engine.Group("/device")
	{
		deviceRoutes.POST("/:id", ds.Model.Update)
		deviceRoutes.PUT("", ds.Model.Add)
		deviceRoutes.GET("/:id", ds.Model.GetByID)
		deviceRoutes.DELETE("/:id", ds.Model.Delete)
	}
}
