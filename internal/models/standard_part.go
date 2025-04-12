package models

type StandardPart struct {
	Id     uint   `gorm:"primarykey"`
	Name   string `gorm:"unique"`
	Code   string `gorm:"unique"`
	Number int
	Tools  []*Tool `gorm:"many2many:tools_standardparts;"`
}

type StandardPartDTO struct {
	Id     uint       `json:"id"`
	Name   string     `json:"name"`
	Code   string     `json:"code"`
	Number int        `json:"number"`
	Tools  []*ToolDTO `json:"tools"`
}

func (sp *StandardPart) ToDTO() *StandardPartDTO {
	toolsDTO := Map(sp.Tools, func(tool *Tool) *ToolDTO {
		return tool.ToDTO()
	})
	return &StandardPartDTO{
		Id:     sp.Id,
		Name:   sp.Name,
		Code:   sp.Code,
		Number: sp.Number,
		Tools:  toolsDTO,
	}
}

func (sp *StandardPartDTO) ToModel() *StandardPart {
	tools := Map(sp.Tools, func(toolDTO *ToolDTO) *Tool {
		return toolDTO.ToModel()
	})
	return &StandardPart{
		Id:     sp.Id,
		Name:   sp.Name,
		Code:   sp.Code,
		Number: sp.Number,
		Tools:  tools,
	}
}

func NewStandardPart(code string, name string, number int) *StandardPart {
	return &StandardPart{Name: name, Code: code, Number: number}
}
