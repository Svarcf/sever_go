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

func NewTool(code string, name string, dimension string, note string, toolStroke string, workpieceDescription string, toolTypeID uint, toolAssociationID *uint) *Tool {
	return &Tool{
		Name:                 name,
		Code:                 code,
		Dimension:            dimension,
		Note:                 note,
		ToolStroke:           toolStroke,
		WorkpieceDescription: workpieceDescription,
		ToolTypeID:           toolTypeID,
		ToolAssociationID:    toolAssociationID,
	}
}
