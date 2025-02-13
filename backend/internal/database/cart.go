package database

import "github.com/google/uuid"

// TODO: should I just delete this?
type Cart struct {
    UserID                    uuid.UUID        `gorm:index json:"user_id"`
    Total                     uint             `json:"total"`
}

type CartItem struct {
    ProductID                 uuid.UUID        `gorm:"index" json:"product_id"`
    ProductSku                uuid.UUID        `gorm:"index" json:"product_sku"`
    Quantity                  uint             `gorm:"not null" json:"quantity"`
}
