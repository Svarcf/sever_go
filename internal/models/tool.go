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

type ToolDTO struct {
	Id                   uint                   `json:"id"`
	Name                 string                 `json:"name"`
	Code                 string                 `json:"code"`
	Dimension            string                 `json:"dimension"`
	Note                 string                 `json:"note"`
	ToolStroke           string                 `json:"tool_stroke"`
	WorkpieceDescription string                 `json:"workpiece_description"`
	ToolType             *ToolTypeDTO           `json:"tool_type"`
	ToolTypeID           uint                   `json:"tool_type_id"`
	ToolAssociation      *ToolDTO               `json:"tool_association"`
	ToolRepairRecords    []*ToolRepairRecordDTO `json:"tool_repair_records"`
	RawMaterials         []*RawMaterialDTO      `json:"raw_materials"`
	StandardParts        []*StandardPartDTO     `json:"standard_parts"`
	MechanicalPresses    []*MechanicalPressDTO  `json:"mechanical_presses"`
}

func (t *Tool) ToDTO() *ToolDTO {
	toolTypeDTO := t.ToolType.ToDTO()
	toolAssociationDTO := (*ToolDTO)(nil)
	if t.ToolAssociation != nil {
		toolAssociationDTO = t.ToolAssociation.ToDTO()
	}
	toolRepairRecordsDTO := Map(t.ToolRepairRecords, func(trr *ToolRepairRecord) *ToolRepairRecordDTO {
		return trr.ToDTO()
	})
	rawMaterialsDTO := Map(t.RawMaterials, func(rm *RawMaterial) *RawMaterialDTO {
		return rm.ToDTO()
	})
	standardPartsDTO := Map(t.StandardParts, func(sp *StandardPart) *StandardPartDTO {
		return sp.ToDTO()
	})
	mechanicalPressesDTO := Map(t.MechanicalPresses, func(mp *MechanicalPress) *MechanicalPressDTO {
		return mp.ToDTO()
	})

	return &ToolDTO{
		Id:                   t.Id,
		Name:                 t.Name,
		Code:                 t.Code,
		Dimension:            t.Dimension,
		Note:                 t.Note,
		ToolStroke:           t.ToolStroke,
		WorkpieceDescription: t.WorkpieceDescription,
		ToolType:             toolTypeDTO,
		ToolTypeID:           t.ToolTypeID,
		ToolAssociation:      toolAssociationDTO,
		ToolRepairRecords:    toolRepairRecordsDTO,
		RawMaterials:         rawMaterialsDTO,
		StandardParts:        standardPartsDTO,
		MechanicalPresses:    mechanicalPressesDTO,
	}
}

func (t *ToolDTO) ToModel() *Tool {
	toolType := t.ToolType.ToModel()
	toolAssociation := (*Tool)(nil)
	if t.ToolAssociation != nil {
		toolAssociation = t.ToolAssociation.ToModel()
	}
	toolRepairRecords := Map(t.ToolRepairRecords, func(trrDTO *ToolRepairRecordDTO) *ToolRepairRecord {
		return trrDTO.ToModel()
	})
	rawMaterials := Map(t.RawMaterials, func(rmDTO *RawMaterialDTO) *RawMaterial {
		return rmDTO.ToModel()
	})
	standardParts := Map(t.StandardParts, func(spDTO *StandardPartDTO) *StandardPart {
		return spDTO.ToModel()
	})
	mechanicalPresses := Map(t.MechanicalPresses, func(mpDTO *MechanicalPressDTO) *MechanicalPress {
		return mpDTO.ToModel()
	})

	return &Tool{
		Id:                   t.Id,
		Name:                 t.Name,
		Code:                 t.Code,
		Dimension:            t.Dimension,
		Note:                 t.Note,
		ToolStroke:           t.ToolStroke,
		WorkpieceDescription: t.WorkpieceDescription,
		ToolTypeID:           toolType.Id,
		ToolType:             *toolType,
		ToolAssociation:      toolAssociation,
		ToolRepairRecords:    toolRepairRecords,
		RawMaterials:         rawMaterials,
		StandardParts:        standardParts,
		MechanicalPresses:    mechanicalPresses,
	}
}

func NewTool(code string, name string, dimension string, note string, toolStroke string,
	workpieceDescription string, toolTypeID uint, toolAssociationID *uint) *Tool {
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
