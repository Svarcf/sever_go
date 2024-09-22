package models

type ToolType struct {
	Id    uint   `gorm:"primarykey"`
	Name  string `gorm:"unique"`
	Code  string `gorm:"unique"`
	Tools []Tool
}
