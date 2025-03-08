package utils

import (
	"context"
	"ecomm-backend/internal/auth"
	"ecomm-backend/internal/database"
	"errors"
)

// TODO: may be try to unit test these funcs
func Authorizer(ctx context.Context, permittedRoles []uint) bool {
  // Check if guest is allowed
  if database.IsPermitted(database.Guest, permittedRoles) {
    return true
  }

  user, err := GetUserFromSession(ctx)
  if err != nil {
    return false
  }

  if !database.IsPermitted(user.Role, permittedRoles) {
    return false
  }
  return true
}

func GetUserFromSession(c context.Context) (*database.User, error) {
  var user database.User

  session, err := GetSession(c)
  if err != nil {
    return nil, errors.New("no session")
  }
  
  err = auth.RetrieveSession(session, &user)
  if err != nil {
    return nil, errors.New("no user")
  }

  return &user, nil
}
