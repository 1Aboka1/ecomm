package database

import "github.com/google/uuid"

type Order struct {
    DefaultModel

    UserID                    uuid.UUID        `gorm:index`
    Total                     uint             
}

type OrderItem struct {
    DefaultModel

    OrderID                    uuid.UUID        `gorm:index`
    ProductID                 uuid.UUID        `gorm:index`
    ProductSku                uuid.UUID        `gorm:index`
    Quantity                  uint             `gorm:"not null"`
}
