package resource

import (
	"net/http"

	fbs "github.com/jschweizer78/ciscoSsh/storage/firebase"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/examples/model"
)

// UserResource for api2go routes
type UserResource struct {
	UserStorage *fbs.UserStorage
}

// FindAll to satisfy api2go data source interface
func (s UserResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []model.User
	users := s.UserStorage.GetAll()

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load users in chunks
func (s UserResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []model.User
		number, size, offset, limit string
		keys                        []int
	)
	users := s.UserStorage.GetAll()

	err := s.UserStorage.Update(user)
	return &Response{Res: user, Code: http.StatusNoContent}, err
}
