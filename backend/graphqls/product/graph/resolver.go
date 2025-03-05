package graph

import "ecomm-backend/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    ProductStore map[string]database.Product
}
