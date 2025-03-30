package models

type ToolType struct {
	Id    uint   `gorm:"primarykey"`
	Name  string `gorm:"unique"`
	Code  string `gorm:"unique"`
	Tools []Tool
}

func NewToolType(code string, name string) *ToolType {
	return &ToolType{Name: name, Code: code}
}
