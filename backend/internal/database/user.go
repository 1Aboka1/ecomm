package database

import (
  "errors"
  "time"

  "github.com/google/uuid"
)

// Base fields used by almost all models
type DefaultModel struct {
  // ID is created randomly using gen_random_uuid()
  ID            uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
  CreatedAt     time.Time       `json:"created_at"`
  UpdatedAt     time.Time       `json:"updated_at"`
  DeletedAt     *time.Time      `json:"deleted_at"`
}

// User model
type User struct {
  DefaultModel

  FirstName      string    `gorm:"size:100" json:"first_name" binding:"required"`
  LastName       *string   `gorm:"size:100" json:"last_name" binding:"required"`
  PhoneNumber    string    `gorm:"unique;not null;size:12" json:"phone_number"`
  Role           uint      `gorm:"not null;default:0"`
}

// Address model: one2one connection
type Address struct {
  DefaultModel

  AddressLine    string    `json:"address_line"`
  City           string    `json:"city"`
  UserID         uuid.UUID `gorm:"index;not null" json:"user_id"`
}

// Role enum
const (
  guest     string  = "guest"
  customer  string  = "customer"
  admin     string  = "admin"
)

// Role is stored as uint in DB.
// Role enum in uint to make role comparable.
// e.g if level is 1, 5 >= 1. Therefore, Admin is also allowed
const (
  Guest     uint = 0
  Customer       = 1
  Admin          = 5
)

// Converts uint to string
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

// Makes comparison to find if user allowed
func IsPermitted(roleLvl uint, permittedRoles []uint) bool {
  for _, permittedRole := range permittedRoles {
    if roleLvl >= permittedRole {
      return true
    }
  }
  return false
}
