package models

import (
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

type UserModel struct {
	gorm.Model
	Email string `gorm:"unique"`
	Password string 
	FirstName string
	LastName string
	DOB datatypes.Date
	IsAdmin bool
	Role string
}