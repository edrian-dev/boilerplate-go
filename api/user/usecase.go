package user

import (
	_models "github.com/nomada-sh/levita-stp/models"
)

// Usecase ...
type Usecase interface {
	Signin(doc *_models.UserInput) _models.Response
	Signup(doc *_models.UserInput) _models.Response
}
