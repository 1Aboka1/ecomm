package database

import "github.com/google/uuid"


type Cart struct {
    DefaultModel

    UserID                    uuid.UUID        `gorm:index json:"user_id"`
    Total                     uint             `json:"total"`
}

type CartItem struct {
    DefaultModel

    CartID                    uuid.UUID        `gorm:"index" json:"cart_id"`              
    ProductID                 uuid.UUID        `gorm:"index" json:"product_id"`
    ProductSku                uuid.UUID        `gorm:"index" json:"product_sku"`
    Quantity                  uint             `gorm:"not null" json:"quantity"`
}
