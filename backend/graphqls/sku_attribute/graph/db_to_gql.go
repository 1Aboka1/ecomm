package graph

import (
	"ecomm-backend/graphqls/sku_attribute/graph/model"
	"ecomm-backend/internal/database"
	"time"
)

func getDeletedAtString(deletedAt *time.Time) *string {
    if deletedAt == nil {
      return nil
    } else {
      deletedAtString := deletedAt.String()
      return &deletedAtString
    }

}

func toSkuGQL(productSku *database.ProductSku) *model.ProductSku {
    createdAt := productSku.CreatedAt.String()
    updatedAt := productSku.UpdatedAt.String()
    deletedAt := getDeletedAtString(productSku.DeletedAt)

    productGQL := &model.ProductSku{ 
        ID: productSku.ID.String(),
        ProductID: productSku.ProductID.String(),
        Sku: productSku.Sku,
        Price: int32(productSku.Price),
        Quantity: int32(productSku.Quantity),
        CreatedAt: createdAt,
        UpdatedAt: updatedAt,
        DeletedAt: deletedAt,
        ProductAttributes: nil, // idk why
    }
    return productGQL
}

func toAttributeGQL(attribute *database.ProductAttribute) *model.ProductAttribute {
    attributeGQL := &model.ProductAttribute{ 
        ID: attribute.ID.String(),
        Type: attribute.Type,
    }
    return attributeGQL
}

func toSkuAttributeGQL(skuAttribute *database.SkuAttributes) *model.SkuAttribute {
	db := database.OrmDb

	var attribute database.ProductAttribute
	result := db.Where("id = ?", skuAttribute.AttributeID).Take(&attribute)
	if result.Error != nil {
		return nil
	}

  skuAttributeGQL := &model.SkuAttribute{ 
    Type: attribute.Type,
    Value: skuAttribute.Value,
  }
  return skuAttributeGQL
}
