package db

import (
	_helpers "github.com/nomada-sh/levita-stp/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const postgresDB = "POSTGRES_DB"

// NewConnection ...
func NewConnection() *gorm.DB {
	dsn := _helpers.GetEnvironmentVariable(postgresDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
