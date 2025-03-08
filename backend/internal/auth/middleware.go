package auth

import (
  "ecomm-backend/internal/database"
  "net/http"

  "github.com/gin-contrib/sessions"
  "github.com/gin-gonic/gin"
)


var (
  AllowOnlyAdmin = authorizer([]uint{database.Admin})
  AllowOnlyCustomerAndGreater = authorizer([]uint{database.Customer, database.Admin})
  AllowOnlyGuest = authorizer([]uint{database.Guest})
  AllowEveryone = authorizer([]uint{database.Guest, database.Admin, database.Customer})
)

func authorizer(permittedRoles []uint) gin.HandlerFunc {
  return func(c *gin.Context) {
    // Check if guest is allowed
    if database.IsPermitted(database.Guest, permittedRoles) {
      c.Next()
      return
    }

    user := GetUserFromSession(c)

    if !database.IsPermitted(user.Role, permittedRoles) {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "unauthorized. Not permitted",
      })
      c.Abort()
      return
    }

    c.Set("user", user)
    c.Next()
  }
}

func GetUserFromSession(c *gin.Context) *database.User {
  var user database.User

  session := sessions.Default(c)
  err := RetrieveSession(session, &user)

  if err != nil {
    return nil
  }

  return &user
}
