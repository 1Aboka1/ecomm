package graph

import (
	"context"
	"ecomm-backend/graphqls/cart/graph/model"
	"ecomm-backend/graphqls/utils"

	"github.com/gin-contrib/sessions"
)

func countCartTotal(cartItems []*model.CartItem) (sum int32) {
  sum = 0
  for _, cartItem := range cartItems {
    sum += int32(cartItem.Quantity)
  }
  return
}

func createCart(session sessions.Session) (*model.Cart, error) {
  var cart *model.Cart = &model.Cart{}
  session.Set("total", 0)
  session.Set("cart_items", []*model.CartItem{})

  session.Save()

  cart.CartItems = []*model.CartItem{}
  cart.Total = 0
  return cart, nil
}

func RetrieveCart(ctx context.Context) (*model.Cart, error) {
  session, err := utils.GetSession(ctx)
  if err != nil {
    return nil, err
  }

  var newCart *model.Cart = &model.Cart{}

  total, ok := session.Get("total").(int32)
  if !ok {
    newCart, err := createCart(session)
    if err != nil {
      return nil, err
    }
    return newCart, nil
  }
  cartItems, ok := session.Get("cart_items").([]*model.CartItem)
  if !ok {
    newCart, err := createCart(session)
    if err != nil {
      return nil, err
    }
    return newCart, nil
  }

  newCart.Total = total
  newCart.CartItems = cartItems

  return newCart, nil
}

func ClearCart(ctx context.Context) (*model.Cart, error) {
  session, err := utils.GetSession(ctx)
  if err != nil {
    return nil, err
  }

  var cart *model.Cart

  session.Delete("total")
  session.Delete("cart_items")
  cart, err = createCart(session)
  if err != nil {
    return nil, err
  }
  return cart, nil
}

func ChangeCartItem(ctx context.Context, input model.CartItemInput) (*model.Cart, error) {
  var newCart *model.Cart
  newCart, err := RetrieveCart(ctx)
  if err != nil {
    return nil, err
  }

  cartItems := newCart.CartItems
  for _, cartItem := range cartItems {
    if cartItem.ProductID == input.ProductID && cartItem.ProductSku == cartItem.ProductSku {
      cartItem.Quantity = input.Quantity
    }
  }

  newCart.Total = countCartTotal(newCart.CartItems)

  return newCart, nil
}

func DeleteCartItem(ctx context.Context, input model.CartItemInput) (*model.Cart, error) {
  var newCart *model.Cart
  newCart, err := RetrieveCart(ctx)
  if err != nil {
    return nil, err
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
  return newCart, nil
}

func AddCartItem(ctx context.Context, input model.CartItemInput) (*model.Cart, error) {
  var newCart *model.Cart
  newCart, err := RetrieveCart(ctx)
  if err != nil {
    return nil, err
  }

  cartItem := model.CartItem{ Quantity: input.Quantity, ProductID: input.ProductID, ProductSku: input.ProductSku }

  newCart.CartItems = append(newCart.CartItems, &cartItem)
  newCart.Total = countCartTotal(newCart.CartItems)

  return newCart, nil
}
