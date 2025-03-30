package models

type StandardPart struct {
	Id     uint   `gorm:"primarykey"`
	Name   string `gorm:"unique"`
	Code   string `gorm:"unique"`
	Number int
	Tools  []*Tool `gorm:"many2many:tools_standardparts;"`
}

func NewStandardPart(code string, name string, number int) *StandardPart {
	return &StandardPart{Name: name, Code: code, Number: number}
}
