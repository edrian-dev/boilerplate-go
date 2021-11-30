package usecase

import (
	_stp "github.com/nomada-sh/levita-stp/api/stp"
	_models "github.com/nomada-sh/levita-stp/models"
)

type usecase struct {
	STP       _stp.Repository
	STPAPIURL string
}

// Input ...
type Input struct {
	STP       _stp.Repository
	STPAPIURL string
}

func NewUsecase(input Input) _stp.Usecase {
	return &usecase{
		STP: input.STP,
	}
}

func (ucase *usecase) Dispersion(doc _models.DispersionInput) (response _models.Response) {
	dispersion, err := ucase.dispersionRequest(doc)
	if err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/dispersion",
			Title:    "Dispersion request rrror",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/stp/usecase/ucase.go",
		})

		return response
	}

	if err := ucase.STP.InsertOne(dispersion); err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/dispersion",
			Title:    "Insert dispersion",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/stp/usecase/ucase.go",
		})

		return response
	}

	response.Data = dispersion
	return response
}

func (ucase *usecase) ChangeStatus(doc *_models.DispersionStatus) (*_models.STPResponse, error) {
	response := &_models.STPResponse{
		Message: "recibido",
	}

	dispersion, err := ucase.STP.FindOne(_models.Dispersion{
		Base: _models.Base{
			ID: uint(doc.ID),
		},
	})
	if err != nil {
		return nil, err
	}

	dispersion.Estado = doc.Estado
	dispersion.CausaDevolucion = doc.CausaDevolucion
	dispersion.TSLiquidacion = doc.TSLiquidacion

	if err := ucase.STP.Update(dispersion); err != nil {
		return nil, err
	}

	return response, nil
}
