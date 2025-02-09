package auth

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyUser(c *gin.Context) (*IDClaims, error) {
    authHeader := c.GetHeader("Authorization")
    _, IDTokenString, ok := strings.Cut(authHeader, " ")
    if !ok {
        err := errors.New("provide correct Authorization header")
        return nil, err
    }
    IDToken, err := verifyIDToken(IDTokenString)

    if err != nil {
        return nil, err
    }
    IDClaims, ok := IDToken.Claims.(*IDClaims)
    if !ok {
        err := errors.New("couldn't derive claims")
        return nil, err
    }
    return IDClaims, nil
}

// func VerifySessionToken(c *gin.Context) (*SessionClaims, error) {
//     authHeader := c.GetHeader("Authorization")
//     _, sessionTokenString, ok := strings.Cut(authHeader, " ")
//     if !ok {
//         err := errors.New("provide correct Authorization header")
//         return nil, err
//     }
//     sessionToken, err := verifySessionToken(sessionTokenString)
// 
//     if err != nil {
//         return nil, err
//     }
//     sessionClaims, ok := sessionToken.Claims.(*SessionClaims)
//     if !ok {
//         err := errors.New("couldn't derive claims")
//         return nil, err
//     }
//     return sessionClaims, nil
// }
