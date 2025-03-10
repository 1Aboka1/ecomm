package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"ecomm-backend/graphqls/product/graph/model"
	"ecomm-backend/internal/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductInput) (*model.Product, error) {
	db := database.OrmDb

	subCategoryID, err := uuid.Parse(input.SubCategoryID)
	if err != nil {
		return nil, err
	}

	product := &database.Product{Name: input.Name, Description: input.Description, Summary: input.Summary, Cover: input.Cover, SubCategoryID: subCategoryID}

	result := db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	productGQL := toProductGQL(product)
	return productGQL, nil
}

// RemoveProduct is the resolver for the removeProduct field.
func (r *mutationResolver) RemoveProduct(ctx context.Context, input model.DeleteProductInput) (*model.Product, error) {
	db := database.OrmDb
	var product database.Product
	result := db.Where("id = ?", input.ID).Take(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Where("id = ?", input.ID).Delete(database.Product{})
	if result.Error != nil {
		return nil, result.Error
	}

	productGQL := toProductGQL(&product)
	return productGQL, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, subCategoryID *string) ([]*model.Product, error) {
	db := database.OrmDb
	var products []database.Product
	var result *gorm.DB

	if subCategoryID != nil {
		result = db.Where("sub_category_id = ?", *subCategoryID).Select("id", "name", "created_at", "updated_at", "deleted_at", "description", "summary", "cover").Find(&products)
	} else {
		result = db.Select("id", "name", "created_at", "updated_at", "deleted_at", "description", "summary", "cover").Find(&products)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	var productsGQL []*model.Product
	for _, product := range products {
		productsGQL = append(productsGQL, toProductGQL(&product))
	}
	return productsGQL, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	db := database.OrmDb
	var product database.Product
	result := db.Where("id = ?", id).Take(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	productGQL := toProductGQL(&product)
	return productGQL, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
