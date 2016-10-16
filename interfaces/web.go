package interfaces

// APIResourcer used to abstract the storage
import "gopkg.in/gin-gonic/gin.v1"

// APIResourcer standard web REST
type APIResourcer interface {
	AddOne(c *gin.Context)
	GetOne(c *gin.Context)
	GetAll(c *gin.Context)
	DeleteOne(c *gin.Context)
	UpdateOne(c *gin.Context)
	MyName() string
}

// APIManyResourcer for batch storage abstraction
type APIManyResourcer interface {
	DeleteMany(c *gin.Context)
	AddMany(c *gin.Context)
	UpdateMany(c *gin.Context)
	Query(c *gin.Context)
}
