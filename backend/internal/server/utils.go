package server

import (
	"ecomm-backend/internal/auth"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyUser(c *gin.Context) (*auth.IDClaims, error) {
    authHeader := c.GetHeader("Authorization")
    _, IDTokenString, ok := strings.Cut(authHeader, " ")
    if !ok {
        err := errors.New("provide correct Authorization header")
        return nil, err
    }
    IDToken, err := auth.VerifyIDToken(IDTokenString)

    if err != nil {
        return nil, err
    }
    IDClaims, ok := IDToken.Claims.(*auth.IDClaims)
    if !ok {
        err := errors.New("couldn't derive claims")
        return nil, err
    }
    return IDClaims, nil
}
