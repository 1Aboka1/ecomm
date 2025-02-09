package auth

import (
	"ecomm-backend/internal/database"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authorizer(permittedRoles []uint) gin.HandlerFunc {
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

        log.Println(user)
        if database.IsPermitted(user.Role, permittedRoles) {
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
