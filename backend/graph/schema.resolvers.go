package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"ecomm-backend/graph/model"
	"ecomm-backend/internal/database"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CategoryInput) (*model.Category, error) {
	db := database.OrmDb
	category := &database.Category{Name: input.Name, Description: input.Description}

	result := db.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	categoryGQL := toCategoryGQL(category)
	return categoryGQL, nil
}

// RemoveCategory is the resolver for the removeCategory field.
func (r *mutationResolver) RemoveCategory(ctx context.Context, input model.DeleteCategoryInput) (*model.Category, error) {
	db := database.OrmDb
	var category database.Category
	result := db.Where("id = ?", input.ID).Take(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Where("id = ?", input.ID).Delete(database.Category{})
	if result.Error != nil {
		return nil, result.Error
	}

	categoryGQL := toCategoryGQL(&category)
	return categoryGQL, nil
}

// CreateSubCategory is the resolver for the createSubCategory field.
func (r *mutationResolver) CreateSubCategory(ctx context.Context, input model.SubCategoryInput) (*model.SubCategory, error) {
	db := database.OrmDb
	category_id, err := uuid.Parse(input.CategoryID)
	if err != nil {
		return nil, err
	}

	category := &database.SubCategory{Name: input.Name, Description: input.Description, CategoryID: category_id}

	result := db.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	categoryGQL := toSubCategoryGQL(category)
	return categoryGQL, nil
}

// RemoveSubCategory is the resolver for the removeSubCategory field.
func (r *mutationResolver) RemoveSubCategory(ctx context.Context, input model.DeleteSubCategoryInput) (*model.SubCategory, error) {
	db := database.OrmDb
	var category database.SubCategory
	result := db.Where("id = ?", input.ID).Take(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Where("id = ?", input.ID).Delete(database.SubCategory{})
	if result.Error != nil {
		return nil, result.Error
	}

	categoryGQL := toSubCategoryGQL(&category)
	return categoryGQL, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductInput) (*model.Product, error) {
	db := database.OrmDb

	sub_category_id, err := uuid.Parse(input.SubCategoryID)
	if err != nil {
		return nil, err
	}

	product := &database.Product{Name: input.Name, Description: input.Description, Summary: input.Summary, Cover: input.Cover, SubCategoryID: sub_category_id}

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

// AddCartItem is the resolver for the addCartItem field.
func (r *mutationResolver) AddCartItem(ctx context.Context, input model.CartItemInput) (*model.Cart, error) {
	var cart *model.Cart

	cart, err := ChangeCartItem(ctx, input)
	if err != nil {
		return nil, errors.New("couldn't create cart session after getting empty cart session")
	}
	return cart, nil
}

// DeleteCartItem is the resolver for the deleteCartItem field.
func (r *mutationResolver) DeleteCartItem(ctx context.Context, input model.DeleteCartItemInput) (*model.Cart, error) {
	var cart *model.Cart

	cart, err := ClearCart(ctx)
	if err != nil {
		return nil, errors.New("couldn't delete cart session after getting empty cart session")
	}
	return cart, nil
}

// ChangeCartItemQuantity is the resolver for the changeCartItemQuantity field.
func (r *mutationResolver) ChangeCartItemQuantity(ctx context.Context, input model.CartItemInput) (*model.Cart, error) {
	var cart *model.Cart

	cart, err := ClearCart(ctx)
	if err != nil {
		return nil, errors.New("couldn't delete cart session after getting empty cart session")
	}
	return cart, nil
}

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

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	db := database.OrmDb
	var categories []database.Category

	result := db.Select("id", "name", "created_at", "updated_at", "deleted_at", "description").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	var categoriesGQL []*model.Category
	for _, category := range categories {
		categoriesGQL = append(categoriesGQL, toCategoryGQL(&category))
	}
	return categoriesGQL, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	db := database.OrmDb
	var category database.Category
	result := db.Where("id = ?", id).Take(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	categoryGQL := toCategoryGQL(&category)
	return categoryGQL, nil
}

// Subcategories is the resolver for the subcategories field.
func (r *queryResolver) Subcategories(ctx context.Context, categoryID string) ([]*model.SubCategory, error) {
	db := database.OrmDb
	var categories []database.SubCategory

	result := db.Where("category_id = ?", categoryID).Select("id", "name", "created_at", "updated_at", "deleted_at", "description").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	var categoriesGQL []*model.SubCategory
	for _, category := range categories {
		categoriesGQL = append(categoriesGQL, toSubCategoryGQL(&category))
	}
	return categoriesGQL, nil
}

// Subcategory is the resolver for the subcategory field.
func (r *queryResolver) Subcategory(ctx context.Context, id string) (*model.SubCategory, error) {
	db := database.OrmDb
	var category database.SubCategory
	result := db.Where("id = ?", id).Take(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	categoryGQL := toSubCategoryGQL(&category)
	return categoryGQL, nil
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

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context) (*model.Cart, error) {
	var cart *model.Cart
	cart, err := RetrieveCart(ctx)
	if err != nil {
		return nil, err
	}
	return cart, nil
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
