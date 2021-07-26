package models

import (
	"database/sql"
	"time"
)

// Base ...
type Base struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
