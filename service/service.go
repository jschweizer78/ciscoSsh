package service

import (
	"github.com/gin-gonic/gin"
	"github.com/melvinmt/firebase"
)

// Service used to service all models
type Service struct {
	Engine   *gin.Engine
	Firebase *firebase.Reference
	Model    Servicer
}

// Servicer is used to service the api(s)
type Servicer interface {
	Update(*gin.Context)
	Add(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
}

// NewService creates a new service for working with API servers and db
func NewService(eng *gin.Engine, fDb *firebase.Reference) *Service {
	model := newDefaultModel
	return &Service{eng, fDb, model}
}
