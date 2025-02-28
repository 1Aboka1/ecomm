package database

import (
	"github.com/google/uuid"
)

type Product struct {
    DefaultModel

    Name                      string              `json:"name"`
    Description               string              `json:"description"`
    Summary                   string              `json:"summary"`
    Cover                     string              `json:"cover"`
    SubCategoryID             uuid.UUID            `gorm:"index" json:"sub_category_id"`
}

type ProductSku struct {
    DefaultModel

    ProductID                 uuid.UUID         `gorm:index`
    Sku                       string            `json:"sku"`       
    Price                     uint              `json:"price"`
    Quantity                  uint              `json:"quantity"`    // quantity left

    ProductAttributes         []*ProductAttribute    `gorm:"many2many:sku_attributes;"`

}

type ProductAttribute struct {
    DefaultModel

    Type                      string             `json:"type"`   
    Value                     string             `json:"value"`

    ProductSkus               []*ProductSku    `gorm:"many2many:sku_attributes;"`
}
