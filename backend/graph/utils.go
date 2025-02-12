package graph

import (
	"ecomm-backend/graph/model"
	"ecomm-backend/internal/database"
	"time"
)

func dbToGQLModel(category *database.Category) *model.Category {
    createdAt := category.CreatedAt.String()
    updatedAt := category.UpdatedAt.String()

    var deletedAt string
    if category.DeletedAt == nil {
        // TODO: this SHOULDN'T BE Now()
        deletedAt = time.Now().String()
    } else {
        deletedAt = category.DeletedAt.String()
    }

    categoryGQL := &model.Category{ 
        ID: category.ID.String(),
        Name: category.Name, 
        Description: category.Description,
        CreatedAt: createdAt,
        UpdatedAt: updatedAt,
        DeletedAt: deletedAt,
    }
    return categoryGQL
}

