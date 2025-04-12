package models

type ToolType struct {
	Id    uint   `gorm:"primarykey"`
	Name  string `gorm:"unique"`
	Code  string `gorm:"unique"`
	Tools []Tool
}

type ToolTypeDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (tt *ToolType) ToDTO() *ToolTypeDTO {
	return &ToolTypeDTO{
		Id:   tt.Id,
		Name: tt.Name,
		Code: tt.Code,
	}
}

func (tt *ToolTypeDTO) ToModel() *ToolType {
	return &ToolType{
		Id:   tt.Id,
		Name: tt.Name,
		Code: tt.Code,
	}
}

func NewToolType(code string, name string) *ToolType {
	return &ToolType{Name: name, Code: code}
}
