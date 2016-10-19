package resource

import (
	"net/http"

	"github.com/jschweizer78/ciscoSsh/models"
	"github.com/jschweizer78/ciscoSsh/storage/tiedot"
	"gopkg.in/gin-gonic/gin.v1"
)

// UserResource for api2go routes
type UserResource struct {
	Stor *tiedot.UserStorage
}

// AddOne to satisfy APIResourcer data source interface
func (ur UserResource) AddOne(c *gin.Context) {
	var obj models.User
	err := c.BindJSON(obj)
	if err != nil {
		panic(err)
	}
	id := ur.Stor.AddOne(obj)
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetOne to satisfy APIResourcer data source interface
func (ur UserResource) GetOne(c *gin.Context) {
	id := c.Param("id")
	usr, err := ur.Stor.GetOne(id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": usr})
}

// GetAll to satisfy APIResourcer data source interface
func (ur UserResource) GetAll(c *gin.Context) {
	users := ur.Stor.GetAll()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DeleteOne to satisfy APIResourcer data source interface
func (ur UserResource) DeleteOne(c *gin.Context) {
	id := c.Param("id")
	err := ur.Stor.DeleteOne(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}
}

// UpdateOne to satisfy APIResourcer data source interface
func (ur UserResource) UpdateOne(c *gin.Context) {
	var obj models.User
	id := c.Param("id")
	err := c.BindJSON(obj)
	if err != nil {
		panic(err)
	}
	obj.SetID(id)
	err = ur.Stor.UpdateOne(obj)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}
}

// MyName to satisfy APIResourcer data source interface
func (ur UserResource) MyName() string {
	return ur.MyName()
}
