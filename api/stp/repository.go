package stp

import (
	_models "github.com/nomada-sh/levita-stp/models"
)

// Repository ...
type Repository interface {
	FindOne(filter _models.Dispersion) (*_models.Dispersion, error)
	InsertOne(doc *_models.Dispersion) error
	Update(doc *_models.Dispersion) error
}
