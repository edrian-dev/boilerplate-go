package repository

import (
	"gorm.io/gorm"

	_user "github.com/nomada-sh/levita-stp/api/user"
	_models "github.com/nomada-sh/levita-stp/models"
)

type postgresRepository struct {
	DB    *gorm.DB
	Model *_models.User
}

func NewRepository(input _models.InputRepository) _user.Repository {
	repository := &postgresRepository{
		DB:    input.DB,
		Model: &_models.User{},
	}

	repository.DB.AutoMigrate(&_models.User{})
	return repository
}

func (repository *postgresRepository) FindOne(filter _models.User) (*_models.User, error) {
	user := &_models.User{}

	result := repository.DB.Where(filter).First(user)
	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *postgresRepository) InsertOne(doc *_models.User) error {
	result := repository.DB.Create(doc)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *postgresRepository) Update(doc *_models.User) error {
	result := repository.DB.Save(doc)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
