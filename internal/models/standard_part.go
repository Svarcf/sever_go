package models

type StandardPart struct {
	Id     uint   `gorm:"primarykey"`
	Name   string `gorm:"unique"`
	Code   string `gorm:"unique"`
	Number int
	Tools  []*Tool `gorm:"many2many:tools_standardparts;"`
}
