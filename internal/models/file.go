package models

type File struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	Location string
}
