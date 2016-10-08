package modelStorage

import (
	"github.com/jschweizer78/ciscoSsh/models"
)

// STInterface standard storage
type STInterface interface {
	CreateOne(sw models.Switch) (string, error)
	CreateMany(sw []models.Switch) ([]string, error)

	ReadOne(id string) (models.Switch, error)
	ReadMany(ids []string) ([]models.Switch, error)

	UpdateOne(sw models.Switch) error
	UpdateMany(sws []models.Switch) error

	DeleteOne(sw string) error
	DeleteMany(sws []string) error

	// CreateOne(user models.User) (string, error)
}
