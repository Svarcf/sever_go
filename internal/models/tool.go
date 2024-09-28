package models

type Tool struct {
	Id                   uint   `gorm:"primarykey"`
	Name                 string `gorm:"unique"`
	Code                 string `gorm:"unique"`
	Dimension            string
	Note                 string
	ToolStroke           string
	WorkpieceDescription string
	ToolTypeID           uint
	ToolType             ToolType
	ToolAssociationID    *uint
	ToolAssociation      *Tool
	ToolRepairRecords    []*ToolRepairRecord
	RawMaterials         []*RawMaterial     `gorm:"many2many:tools_rawmaterials;"`
	StandardParts        []*StandardPart    `gorm:"many2many:tools_standardparts;"`
	MechanicalPresses    []*MechanicalPress `gorm:"many2many:tools_mechanicalpresses;"`
}
