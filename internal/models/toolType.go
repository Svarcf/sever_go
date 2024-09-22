package models

import "gorm.io/gorm"

type ToolType struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Code  string `gorm:"unique"`
	Tools []Tool
}
