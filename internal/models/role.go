package models

type Role struct {
	Id   uint   `gorm:"primarykey"`
	Name string `gorm:"unique"`
}

func NewRole(name string) *Role {
	return &Role{Name: name}
}
