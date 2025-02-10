package database

import (
	"github.com/google/uuid"
)

type Category struct {
    DefaultModel                      
    Name        string                `json:"name"`       
    Description string                `json:"description"`
}                                     

type SubCategory struct {
    DefaultModel

    CategoryID              uuid.UUID    `gorm:index`
    Name                    string
    Description             string
}
