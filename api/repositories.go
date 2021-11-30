package api

import (
	_stpRepository "github.com/nomada-sh/levita-stp/api/stp/repository"
	_userRepository "github.com/nomada-sh/levita-stp/api/user/repository"
	_db "github.com/nomada-sh/levita-stp/db"
	_models "github.com/nomada-sh/levita-stp/models"
)

var input = _models.InputRepository{
	DB: _db.NewConnection(),
}

var user = _userRepository.NewRepository(input)
var stp = _stpRepository.NewRepository(input)
