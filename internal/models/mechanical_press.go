package models

type MechanicalPress struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Code   string `gorm:"unique"`
	Number uint
	Tools  []*Tool `gorm:"many2many:tools_mechanicalpresses;"`
}

func NewMechanicalPress(code string, name string, number uint) *MechanicalPress {
	return &MechanicalPress{Name: name, Code: code, Number: number}
}
