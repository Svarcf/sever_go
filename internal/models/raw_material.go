package models

type RawMaterial struct {
	Id     uint   `gorm:"primarykey"`
	Name   string `gorm:"unique"`
	Code   string `gorm:"unique"`
	Number int
	Tools  []*Tool `gorm:"many2many:tools_rawmaterials;"`
}

func NewRawMaterial(code string, name string, number int) *RawMaterial {
	return &RawMaterial{Name: name, Code: code, Number: number}
}
