package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"ecomm-backend/graphqls/sku_attribute/graph/model"
	"fmt"
)

// CreateProductSku is the resolver for the createProductSku field.
func (r *mutationResolver) CreateProductSku(ctx context.Context, input model.ProductSkuInput) (*model.ProductSku, error) {
	panic(fmt.Errorf("not implemented: CreateProductSku - createProductSku"))
}

// CreateProductAttribute is the resolver for the createProductAttribute field.
func (r *mutationResolver) CreateProductAttribute(ctx context.Context, input model.ProductAttributeInput) (*model.ProductAttribute, error) {
	panic(fmt.Errorf("not implemented: CreateProductAttribute - createProductAttribute"))
}

// ConnectSkuAttribute is the resolver for the connectSkuAttribute field.
func (r *mutationResolver) ConnectSkuAttribute(ctx context.Context, input model.SkuAttributeInput) (*model.SkuAttribute, error) {
	panic(fmt.Errorf("not implemented: ConnectSkuAttribute - connectSkuAttribute"))
}

// ProductSku is the resolver for the product_sku field.
func (r *queryResolver) ProductSku(ctx context.Context, id string) (*model.ProductSku, error) {
	panic(fmt.Errorf("not implemented: ProductSku - product_sku"))
}

// ProductSkus is the resolver for the product_skus field.
func (r *queryResolver) ProductSkus(ctx context.Context, productID string) ([]*model.ProductSku, error) {
	panic(fmt.Errorf("not implemented: ProductSkus - product_skus"))
}

// SkusAttributes is the resolver for the skus_attributes field.
func (r *queryResolver) SkusAttributes(ctx context.Context, productSkuID string) ([]*model.SkuAttribute, error) {
	panic(fmt.Errorf("not implemented: SkusAttributes - skus_attributes"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
