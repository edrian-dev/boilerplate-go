package user

import (
	_models "github.com/siends/siends-api/models"
)

// Repository ...
type Repository interface {
	FindOne(filter _models.User) (*_models.User, error)
	InsertOne(doc *_models.User) error
	Update(doc *_models.User) error
}
