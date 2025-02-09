package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type DefaultModel struct {
    ID            uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid()";json:"id"`
    CreatedAt     time.Time       
    UpdatedAt     time.Time
    DeletedAt     *time.Time
}

type User struct {
    DefaultModel

    FirstName      string    `gorm:"size:100" json:"first_name" binding:"required"`
    LastName       *string   `gorm:"size:100" json:"last_name" binding:"required"`
    PhoneNumber    string    `gorm:"unique;not null;size:12" json:"phone_number"`
    Role           uint      `gorm:"not null;default:0"`
}

type Address struct {
    DefaultModel

    AddressLine    string    `json:"address_line";validate:"required"`
    City           string    `json:"city";validate:"required"`
    UserID         uuid.UUID `gorm:"index;not null";json:"user_id";validate:"required"`
}

const (
    guest     string  = "guest"
    customer  string  = "customer"
    admin     string  = "admin"
)

const (
    Guest     uint = 0
    Customer       = 1
    Admin          = 5
)

func GetRoleString(role uint) (string, error) {
    switch role {
    case Guest:
        return guest, nil
    case Customer:
        return customer, nil
    case Admin:
        return admin, nil
    }

    return "", errors.New("invalid role")
}

func IsPermitted(roleLvl uint, permittedRoles []uint) bool {
    for _, permittedRole := range permittedRoles {
        if roleLvl >= permittedRole {
            return true
        }
    }
    return false
}
