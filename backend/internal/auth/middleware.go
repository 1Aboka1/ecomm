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

        var user database.User

        session := sessions.Default(c)
        err := RetrieveSession(session, &user)

        if err != nil {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if user.ID.String() == "00000000-0000-0000-0000-000000000000" {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no user found"})
            return
        }
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

