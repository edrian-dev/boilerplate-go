package user

import (
	_models "github.com/nomada-sh/levita-stp/models"
)

// Repository ...
type Repository interface {
	FindOne(filter _models.User) (*_models.User, error)
	InsertOne(doc *_models.User) error
	Update(doc *_models.User) error
}
