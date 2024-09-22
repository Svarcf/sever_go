package models

import "gorm.io/gorm"

type MechanicalPress struct {
	gorm.Model
	Name   string
	Code   string `gorm:"unique"`
	Number uint
	Tools  []*Tool `gorm:"many2many:tools_mechanicalpresses;"`
}
