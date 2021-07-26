package api

import (
	_userUsecase "github.com/siends/siends-api/api/user/usecase"
)

// User ...
var User = _userUsecase.NewUsecase(_userUsecase.Input{
	User: user,
})
