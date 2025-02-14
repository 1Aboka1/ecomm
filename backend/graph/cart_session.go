// TODO: I should refactor this file. Cart pointer shouldn't be in the parameters
// it should be in return. Otherwise, can't understand what the function does.
// plus, gin.Context type assertion should be in caller side (maybe)
package graph

import (
	"context"
	"ecomm-backend/graph/model"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func countCartTotal(cartItems []*model.CartItem) (sum int32) {
  sum = 0
  for _, cartItem := range cartItems {
    sum += int32(cartItem.Quantity)
  }
  return
}

func createCart(session sessions.Session, cart *model.Cart) error {
    session.Set("total", 0)
    session.Set("cart_items", []*model.CartItem{})

    session.Save()

    cart.CartItems = []*model.CartItem{}
    cart.Total = 0
    return nil
}

func RetrieveCart(ctx context.Context, cart *model.Cart) error {
  c, ok := ctx.(*gin.Context)
  if !ok {
    return errors.New("no context")
  }

  session := sessions.Default(c)

  var newCart model.Cart

  total, ok := session.Get("total").(int32)
  if !ok {
    createCart(session, &newCart)
    *cart = newCart
    return nil
  }
  cartItems, ok := session.Get("cart_items").([]*model.CartItem)
  if !ok {
    createCart(session, &newCart)
    *cart = newCart
    return nil
  }

  newCart.Total = total
  newCart.CartItems = cartItems

  *cart = newCart

  return nil
}

func ClearCart(ctx context.Context, cart *model.Cart) error {
  c, ok := ctx.(*gin.Context)
  if !ok {
    return errors.New("no context")
  }

  session := sessions.Default(c)

  session.Delete("total")
  session.Delete("cart_items")
  err := createCart(session, cart)
  if err != nil {
    return err
  }
  return nil
}

func ChangeCartItem(ctx context.Context, cart *model.Cart, input model.CartItemInput) error {
  c, ok := ctx.(*gin.Context)
  if !ok {
    return errors.New("no context")
  }

  var newCart model.Cart
  err := RetrieveCart(c, &newCart)
  if err != nil {
    return err
  }

  cartItems := newCart.CartItems
  for _, cartItem := range cartItems {
    if cartItem.ProductID == input.ProductID && cartItem.ProductSku == cartItem.ProductSku {
      cartItem.Quantity = input.Quantity
    }
  }

  cart.Total = countCartTotal(newCart.CartItems)

  return nil
}

func DeleteCartItem(ctx context.Context, cart *model.Cart, input model.CartItemInput) error {
  c, ok := ctx.(*gin.Context)
  if !ok {
    return errors.New("no context")
  }

  var newCart model.Cart
  err := RetrieveCart(c, &newCart)
  if err != nil {
    return err
  }

  cartItems := newCart.CartItems
  newCart.CartItems = []*model.CartItem{}
  for _, cartItem := range cartItems {
    if cartItem.ProductID != input.ProductID && cartItem.ProductSku != cartItem.ProductSku {
      newCart.CartItems = append(newCart.CartItems, cartItem)
      break
    }
  }

  newCart.Total = countCartTotal(newCart.CartItems)
  *cart = newCart
  return nil
}

func AddCartItem(ctx context.Context, cart *model.Cart, input model.CartItemInput) error {
  c, ok := ctx.(*gin.Context)
  if !ok {
    return errors.New("no context")
  }

  var newCart model.Cart
  err := RetrieveCart(c, &newCart)
  if err != nil {
    return err
  }

  cartItem := model.CartItem{ Quantity: input.Quantity, ProductID: input.ProductID, ProductSku: input.ProductSku }

  newCart.CartItems = append(newCart.CartItems, &cartItem)
  newCart.Total = countCartTotal(newCart.CartItems)

  *cart = newCart

  return nil
}
