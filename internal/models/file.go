package models

type File struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	Location string
}

func NewFile(name string, location string) *File {
	return &File{Name: name, Location: location}
}
