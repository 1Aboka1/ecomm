package auth

import (
	"ecomm-backend/internal/database"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/google/uuid"
)

func CreateSession(session sessions.Session, user *database.User) error {
    userID, err := user.ID.MarshalText()

    if err != nil {
        return err
    }
    session.Set("id", userID)
    session.Set("first_name", user.FirstName)
    session.Set("last_name", user.LastName)
    session.Set("phone_number", user.PhoneNumber)
    session.Set("role", user.Role)
    session.Set("address_line", user.AddressLine)
    session.Set("city", user.City)

    session.Save()
    return nil
}

func RetrieveSession(session sessions.Session, user *database.User) error {
    var newUser database.User

    id, ok := session.Get("id").(uuid.UUID)
    if !ok {
        return errors.New("session doesn't have id")
    }
    firstName, ok := session.Get("first_name").(string)
    if !ok {
        return errors.New("session doesn't have first_name")
    }
    lastName, ok := session.Get("last_name").(string)
    if !ok {
        return errors.New("session doesn't have last_name")
    }
    phoneNumber, ok := session.Get("phone_number").(string)
    if !ok {
        return errors.New("session doesn't have phone_number")
    }
    role, ok := session.Get("role").(uint)
    if !ok {
        return errors.New("session doesn't have role")
    }
    addressLine, ok := session.Get("address_line").(string)
    if !ok {
        return errors.New("session doesn't have address_line")
    }
    city, ok := session.Get("city").(string)
    if !ok {
        return errors.New("session doesn't have city")
    }

    newUser.ID = id
    newUser.FirstName = firstName
    newUser.LastName = &lastName // is it ok?
    newUser.PhoneNumber = phoneNumber
    newUser.Role = role
    newUser.AddressLine = addressLine
    newUser.City = city

    return nil
}
