package api

import (
	_stpUsecase "github.com/nomada-sh/levita-stp/api/stp/usecase"
	_userUsecase "github.com/nomada-sh/levita-stp/api/user/usecase"
)

// User ...
var User = _userUsecase.NewUsecase(_userUsecase.Input{
	User: user,
})

// STP ...
var STP = _stpUsecase.NewUsecase(_stpUsecase.Input{
	STP: stp,
})
