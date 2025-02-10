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

    CategoryID              uuid.UUID    `gorm:"index" json:"category_id"`
    Name                    string       `json:"name"`       
    Description             string       `json:"description"`
}
