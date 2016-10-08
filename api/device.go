package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jschweizer78/ciscoSsh/models"
)

// DeviceAPI is used for the interface
type DeviceAPI struct {
	Device models.Device
}

// Update is an API to update Cisco devices
func (ds *DeviceAPI) Update(c *gin.Context) {

}

// Add is an API add Cisco devices to the db
func (ds *DeviceAPI) Add(c *gin.Context) {

}

// Delete is an API remove Cisco devices to the db
func (ds *DeviceAPI) Delete(c *gin.Context) {

}

// GetByID is an API get Cisco devices to the db
func (ds *DeviceAPI) GetByID(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "The ID = %s", id)
}
