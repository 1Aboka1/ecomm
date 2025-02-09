package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type IDClaims struct {
    PhoneNumber           string    `json:phone_number`
    jwt.RegisteredClaims
}

func CreateIDToken(phoneNumber string) (string, error) {
    claims := IDClaims{
        phoneNumber,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func VerifyIDToken(tokenString string) (*jwt.Token, error) {
    var claimsOTP IDClaims
    _ = tokenString

    token, err := jwt.ParseWithClaims(tokenString, &claimsOTP, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return token, nil
}
