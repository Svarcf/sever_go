package models

type Role struct {
	Id uint `gorm:"primarykey"`

	Name string `gorm:"unique"`
}
