package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"ecomm-backend/graphqls/cart/graph/model"
	"errors"
)

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

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context) (*model.Cart, error) {
	var cart *model.Cart
	cart, err := RetrieveCart(ctx)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
