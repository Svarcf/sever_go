package models

import "gorm.io/gorm"

type StandardPart struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Code   string `gorm:"unique"`
	Number uint
	Tools  []*Tool `gorm:"many2many:tools_standardparts;"`
}
