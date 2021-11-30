package stp

import (
	_models "github.com/nomada-sh/levita-stp/models"
)

// Usecase ...
type Usecase interface {
	Dispersion(doc _models.DispersionInput) _models.Response
	ChangeStatus(doc *_models.DispersionStatus) (*_models.STPResponse, error)
}
