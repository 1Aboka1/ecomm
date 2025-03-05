package graph

import (
	"ecomm-backend/graphqls/product/graph/model"
	"ecomm-backend/internal/database"
)

func toProductGQL(product *database.Product) *model.Product {
    createdAt := product.CreatedAt.String()
    updatedAt := product.UpdatedAt.String()

    var deletedAt *string
    if product.DeletedAt == nil {
        deletedAt = nil
    } else {
        deletedAtString := product.DeletedAt.String()
        deletedAt = &deletedAtString
    }

    productGQL := &model.Product{ 
        ID: product.ID.String(),
        Name: product.Name, 
        Description: product.Description,
        Summary: product.Summary,
        Cover: product.Cover,
        SubCategoryID: product.SubCategoryID.String(),
        CreatedAt: createdAt,
        UpdatedAt: updatedAt,
        DeletedAt: deletedAt,
    }
    return productGQL
}
