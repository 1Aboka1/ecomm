package database

import "github.com/google/uuid"

type Order struct {
  DefaultModel

  UserID                    uuid.UUID        `gorm:index json:"user_id"`
  Total                     uint             `json:"total"`
}

type OrderItem struct {
  DefaultModel

  OrderID                    uuid.UUID        `gorm:"index"" json:"order_id"`
  ProductID                 uuid.UUID        `gorm:"index" json:"product_id"`
  ProductSku                uuid.UUID        `gorm:"index" json:"product_sku"`
  Quantity                  uint             `gorm:"not null" json:"quantity"`
}
