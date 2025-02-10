package database

import (
	"github.com/google/uuid"
)

type Product struct {
    DefaultModel

    Name                      string              `json:"name";validate:"required"`
    Description               string              `json:"description";validate:"required"`
    Summary                   string              `json:"summary";validate:"required"`
    Cover                     string              `json:"cover";validate:"required"`
    SubCategoryID             uuid.UUID            `gorm:index`
}

type ProductSku struct {
    DefaultModel

    ProductID                 uuid.UUID        `gorm:index`
    Sku                       string
    Price                     uint
    Quantity                  uint

    ProductAttributes         []*ProductAttribute    `gorm:"many2many:sku_attributes;"`

}

// TODO: Need to make value available only from ENUM (use some method from gorm).
// also need to make it available in admin dashboard
type ProductAttribute struct {
    DefaultModel

    Type                      string
    Value                     string

    ProductSkus               []*ProductSku    `gorm:"many2many:sku_attributes;"`
}
