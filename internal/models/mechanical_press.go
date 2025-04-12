package models

type MechanicalPress struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Code   string `gorm:"unique"`
	Number uint
	Tools  []*Tool `gorm:"many2many:tools_mechanicalpresses;"`
}

type MechanicalPressDTO struct {
	Id     uint       `json:"id"`
	Name   string     `json:"name"`
	Code   string     `json:"code"`
	Number uint       `json:"number"`
	Tools  []*ToolDTO `json:"tools"`
}

func (mp *MechanicalPress) ToDTO() *MechanicalPressDTO {
	toolsDTO := Map(mp.Tools, func(tool *Tool) *ToolDTO {
		return tool.ToDTO()
	})

	return &MechanicalPressDTO{
		Id:     mp.Id,
		Name:   mp.Name,
		Code:   mp.Code,
		Number: mp.Number,
		Tools:  toolsDTO,
	}
}

func (mp *MechanicalPressDTO) ToModel() *MechanicalPress {
	tools := Map(mp.Tools, func(toolDTO *ToolDTO) *Tool {
		return toolDTO.ToModel()
	})

	return &MechanicalPress{
		Id:     mp.Id,
		Name:   mp.Name,
		Code:   mp.Code,
		Number: mp.Number,
		Tools:  tools,
	}
}

func NewMechanicalPress(code string, name string, number uint) *MechanicalPress {
	return &MechanicalPress{Name: name, Code: code, Number: number}
}
