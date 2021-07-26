package restful

import (
	"net/http"

	"github.com/labstack/echo"

	_user "github.com/siends/siends-api/api/user"
	_models "github.com/siends/siends-api/models"
)

type deliveryLayer struct {
	User   _user.Usecase
	Router *echo.Echo
}

// Input ...
type Input struct {
	Router *echo.Echo
	User   _user.Usecase
}

func NewRouter(input Input) {
	delivery := deliveryLayer{
		User:   input.User,
		Router: input.Router,
	}

	delivery.Router.POST("/signup", delivery.signup)
	delivery.Router.POST("/signin", delivery.signin)
}

func (delivery *deliveryLayer) signin(c echo.Context) error {
	var statusCode int
	var response _models.Response
	body := new(_models.UserInput)
	if err := c.Bind(body); err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signin",
			Title:    "Error while decoding data",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/user/delivery/restful/main.go",
		})

		return c.JSON(http.StatusBadRequest, response)
	}

	response = delivery.User.Signin(body)
	if response.Errors != nil {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = http.StatusOK
	}

	return c.JSON(statusCode, response)
}

func (delivery *deliveryLayer) signup(c echo.Context) error {
	var statusCode int
	var response _models.Response
	body := new(_models.UserInput)
	if err := c.Bind(body); err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signup",
			Title:    "Error while decoding data",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/user/delivery/restful/main.go",
		})

		return c.JSON(http.StatusBadRequest, response)
	}

	response = delivery.User.Signup(body)
	if response.Errors != nil {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = http.StatusCreated
	}

	return c.JSON(statusCode, response)
}
