package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	userName  string `gorm:"unique"`
	FirstName string
	LastName  string
	Password  string
	RoleID    uint
	Role      Role
}
