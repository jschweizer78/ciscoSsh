package service

import "github.com/gin-gonic/gin"

func newDefaultModel() *DefaultModel {
	return &DefaultModel{}
}

// DefaultModel used to stand in for interface
type DefaultModel struct {
}

// Update is an API to update Cisco devices
func (dm *DefaultModel) Update(c *gin.Context) {

}

// Add is an API add Cisco devices to the db
func (dm *DefaultModel) Add(c *gin.Context) {

}

// Delete is an API remove Cisco devices to the db
func (dm *DefaultModel) Delete(c *gin.Context) {

}

// GetByID is an API get Cisco devices to the db
func (dm *DefaultModel) GetByID(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "The ID = %s", id)
}
