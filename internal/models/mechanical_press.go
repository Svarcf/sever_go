package models

type MechanicalPress struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Code   string `gorm:"unique"`
	Number uint
	Tools  []*Tool `gorm:"many2many:tools_mechanicalpresses;"`
}
