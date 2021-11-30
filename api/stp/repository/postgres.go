package repository

import (
	"gorm.io/gorm"

	_stp "github.com/nomada-sh/levita-stp/api/stp"
	_models "github.com/nomada-sh/levita-stp/models"
)

type postgresRepository struct {
	DB    *gorm.DB
	Model *_models.Dispersion
}

func NewRepository(input _models.InputRepository) _stp.Repository {
	repository := &postgresRepository{
		DB:    input.DB,
		Model: &_models.Dispersion{},
	}

	repository.DB.AutoMigrate(&_models.Dispersion{})
	return repository
}

func (repository *postgresRepository) FindOne(filter _models.Dispersion) (*_models.Dispersion, error) {
	dispersion := &_models.Dispersion{}

	result := repository.DB.Where(filter).First(dispersion)
	if err := result.Error; err != nil {
		return nil, err
	}

	return dispersion, nil
}

func (repository *postgresRepository) InsertOne(doc *_models.Dispersion) error {
	result := repository.DB.Create(doc)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *postgresRepository) Update(doc *_models.Dispersion) error {
	result := repository.DB.Save(doc)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
