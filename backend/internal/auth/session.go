package auth

import (
	"ecomm-backend/internal/database"
	"errors"
	_ "log"

	"github.com/gin-contrib/sessions"
	"github.com/google/uuid"
)

func CreateSession(session sessions.Session, user *database.User) error {
    session.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 7 * 2})

    userID, err := user.ID.MarshalBinary()
    if err != nil {
        return err
    }

    session.Set("id", userID)
    session.Set("first_name", user.FirstName)
    session.Set("last_name", user.LastName)
    session.Set("phone_number", user.PhoneNumber)
    session.Set("role", user.Role)

    session.Save()
    return nil
}

func RetrieveSession(session sessions.Session, user *database.User) error {
    var newUser database.User

    id, ok := session.Get("id").([]byte)
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

    uuID, err := uuid.FromBytes(id)
    if err != nil {
        return err
    }
    newUser.ID = uuID
    newUser.FirstName = firstName
    newUser.LastName = &lastName // is it ok?
    newUser.PhoneNumber = phoneNumber
    newUser.Role = role

    *user = newUser

    return nil
}

func DeleteSession(session sessions.Session) {
    session.Clear()
}
