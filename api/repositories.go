package api

import (
	_userRepository "github.com/siends/siends-api/api/user/repository"
	_db "github.com/siends/siends-api/db"
	_models "github.com/siends/siends-api/models"
)

var input = _models.InputRepository{
	DB: _db.NewConnection(),
}

var user = _userRepository.NewRepository(input)
