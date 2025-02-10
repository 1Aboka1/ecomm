package auth

import (
	"errors"

	"gorm.io/gorm"

	"ecomm-backend/internal/database"
)

func SignUpUser(postUser *database.User) error {
    ok := SignInUser(postUser)
    if ok {
        return nil
    }
    db := database.OrmDb

    var dbUser database.User

    dbUser.FirstName = postUser.FirstName
    dbUser.PhoneNumber = postUser.PhoneNumber
    dbUser.Role = database.Admin
    dbUser.LastName = postUser.LastName

    result := db.Create(&dbUser)
    if errors.Is(result.Error, gorm.ErrModelValueRequired) {
        return result.Error
    }

    *postUser = dbUser
    return nil
}

func SignInUser(postUser *database.User) bool {
    db := database.OrmDb
    var dbUser database.User

    result := db.Where("phone_number = ?", postUser.PhoneNumber).First(&dbUser)
    if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
        *postUser = dbUser
        return true
    }

    return false
}
