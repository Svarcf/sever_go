package models

type RawMaterial struct {
	Id     uint   `gorm:"primarykey"`
	Name   string `gorm:"unique"`
	Code   string `gorm:"unique"`
	Number int
	Tools  []*Tool `gorm:"many2many:tools_rawmaterials;"`
}

type RawMaterialDTO struct {
	Id     uint       `json:"id"`
	Name   string     `json:"name"`
	Code   string     `json:"code"`
	Number int        `json:"number"`
	Tools  []*ToolDTO `json:"tools"`
}

func (rm *RawMaterial) ToDTO() *RawMaterialDTO {
	toolsDTO := Map(rm.Tools, func(tool *Tool) *ToolDTO {
		return tool.ToDTO()
	})

	return &RawMaterialDTO{
		Id:     rm.Id,
		Name:   rm.Name,
		Code:   rm.Code,
		Number: rm.Number,
		Tools:  toolsDTO,
	}
}

func (rm *RawMaterialDTO) ToModel() *RawMaterial {
	tools := Map(rm.Tools, func(tDTO *ToolDTO) *Tool {
		return tDTO.ToModel()
	})
	return &RawMaterial{
		Id:     rm.Id,
		Name:   rm.Name,
		Code:   rm.Code,
		Number: rm.Number,
		Tools:  tools,
	}
}

func NewRawMaterial(code string, name string, number int) *RawMaterial {
	return &RawMaterial{Name: name, Code: code, Number: number}
}
