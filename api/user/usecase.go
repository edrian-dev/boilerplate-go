package user

import (
	_models "github.com/siends/siends-api/models"
)

// Usecase ...
type Usecase interface {
	Signin(doc *_models.UserInput) _models.Response
	Signup(doc *_models.UserInput) _models.Response
}
