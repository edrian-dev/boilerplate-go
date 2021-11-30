package restful

import (
	"net/http"

	"github.com/labstack/echo"

	_stp "github.com/nomada-sh/levita-stp/api/stp"
	_models "github.com/nomada-sh/levita-stp/models"
)

type deliveryLayer struct {
	STP    _stp.Usecase
	Router *echo.Echo
}

// Input ...
type Input struct {
	Router *echo.Echo
	STP    _stp.Usecase
}

func NewRouter(input Input) {
	delivery := deliveryLayer{
		STP:    input.STP,
		Router: input.Router,
	}

	delivery.Router.POST("/dispersion", delivery.dispersion)
	delivery.Router.POST("/signin", delivery.changeStatus)
}

func (delivery *deliveryLayer) dispersion(c echo.Context) error {
	var statusCode int
	var response _models.Response
	body := new(_models.DispersionInput)
	if err := c.Bind(body); err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/dispersion",
			Title:    "Error while decoding data",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/stp/delivery/restful/main.go",
		})

		return c.JSON(http.StatusBadRequest, response)
	}

	response = delivery.STP.Dispersion(*body)
	if response.Errors != nil {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = http.StatusOK
	}

	return c.JSON(statusCode, response)
}

func (delivery *deliveryLayer) changeStatus(c echo.Context) error {
	var statusCode int
	body := new(_models.DispersionStatus)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	response, err := delivery.STP.ChangeStatus(body)
	if err != nil {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = http.StatusCreated
	}

	return c.JSON(statusCode, response)
}
