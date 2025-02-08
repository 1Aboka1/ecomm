package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
    gorm.Model
    Name        string
    Description string
}

type SubCategory struct {
    DefaultModel

    CategoryID              uuid.UUID    `gorm:index`
    Name                    string
    Description             string
}
