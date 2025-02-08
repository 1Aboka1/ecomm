package database

import "github.com/google/uuid"


type Cart struct {
    DefaultModel

    UserID                    uuid.UUID        `gorm:index`
    Total                     uint             
}

type CartItem struct {
    DefaultModel

    CartID                    uuid.UUID        `gorm:index`
    ProductID                 uuid.UUID        `gorm:index`
    ProductSku                uuid.UUID        `gorm:index`
    Quantity                  uint             `gorm:"not null"`
}
