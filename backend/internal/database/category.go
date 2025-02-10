package database

import (
	"github.com/google/uuid"
)

type Category struct {
    DefaultModel                      
    Name        string                `json:"name";validate:"required"`       
    Description string                `json:"description";validate:"required"`
}                                     

type SubCategory struct {
    DefaultModel

    CategoryID              uuid.UUID    `gorm:index`
    Name                    string
    Description             string
}
