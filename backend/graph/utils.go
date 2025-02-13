package graph

import (
	"ecomm-backend/graph/model"
	"ecomm-backend/internal/database"
	"errors"

	"github.com/gin-contrib/sessions"
)

func toCategoryGQL(category *database.Category) *model.Category {
    createdAt := category.CreatedAt.String()
    updatedAt := category.UpdatedAt.String()

    var deletedAt *string
    if category.DeletedAt == nil {
        deletedAt = nil
    } else {
        deletedAtString := category.DeletedAt.String()
        deletedAt = &deletedAtString
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

func toSubCategoryGQL(category *database.SubCategory) *model.SubCategory {
    createdAt := category.CreatedAt.String()
    updatedAt := category.UpdatedAt.String()

    var deletedAt *string
    if category.DeletedAt == nil {
        deletedAt = nil
    } else {
        deletedAtString := category.DeletedAt.String()
        deletedAt = &deletedAtString
    }

    categoryGQL := &model.SubCategory{ 
        ID: category.ID.String(),
        Name: category.Name, 
        Description: category.Description,
        CreatedAt: createdAt,
        UpdatedAt: updatedAt,
        DeletedAt: deletedAt,
        CategoryID: category.CategoryID.String(),
    }
    return categoryGQL
}


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

func CreateCartSession(session sessions.Session, cart *model.Cart) error {
    session.Set("total", cart.Total)
    session.Set("cart_items", cart.CartItems)

    session.Save()
    return nil
}

func RetrieveCartSession(session sessions.Session, cart *model.Cart) error {
    var newCart model.Cart

    total, ok := session.Get("total").(int32)
    if !ok {
        return errors.New("cart session doesn't have total")
    }
    cartItems, ok := session.Get("cart_items").([]*model.CartItem)
    if !ok {
        return errors.New("cart session doesn't have cart items")
    }

    newCart.Total = total
    newCart.CartItems = cartItems

    *cart = newCart

    return nil
}

func DeleteSessionCart(session sessions.Session) {
    session.Clear()
}
