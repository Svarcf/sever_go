package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type ToolService struct {
	DB *gorm.DB
}

func NewToolService(db *gorm.DB) *ToolService {
	return &ToolService{DB: db}
}

func (s *ToolService) CreateTool(tool *models.Tool) (*models.ToolDTO, error) {
	err := s.DB.Create(tool).Error
	if err != nil {
		return nil, err
	}
	return tool.ToDTO(), nil
}

func (s *ToolService) GetToolById(id uint) (*models.ToolDTO, error) {
	var tool *models.Tool
	err := s.DB.First(&tool, id).Error
	if err != nil {
		return nil, err
	}
	return tool.ToDTO(), nil
}

func (s *ToolService) GetTools() ([]*models.ToolDTO, error) {
	var tools []*models.Tool
	err := s.DB.Find(&tools).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(tools, func(tool *models.Tool) *models.ToolDTO {
		return tool.ToDTO()
	})

	return dtos, nil
}

func (s *ToolService) GetToolByCode(code string) (*models.ToolDTO, error) {
	var tool *models.Tool
	err := s.DB.Where("code = ?", code).First(&tool).Error
	if err != nil {
		return nil, err
	}
	return tool.ToDTO(), nil
}

func (s *ToolService) GetRawMaterials(tool *models.Tool) ([]*models.RawMaterialDTO, error) {
	var rawMaterials []*models.RawMaterial
	err := s.DB.Model(tool).Association("RawMaterials").Find(&rawMaterials)
	if err != nil {
		return nil, err
	}
	dtos := models.Map(rawMaterials, func(rm *models.RawMaterial) *models.RawMaterialDTO {
		return rm.ToDTO()
	})
	return dtos, nil
}

func (s *ToolService) GetStandardParts(tool *models.Tool) ([]*models.StandardPartDTO, error) {
	var standardParts []*models.StandardPart
	err := s.DB.Model(tool).Association("StandardParts").Find(&standardParts)
	if err != nil {
		return nil, err
	}
	dtos := models.Map(standardParts, func(sp *models.StandardPart) *models.StandardPartDTO {
		return sp.ToDTO()
	})
	return dtos, nil
}

func (s *ToolService) GetMechanicalPresses(tool *models.Tool) ([]*models.MechanicalPressDTO, error) {
	var mechanicalPresses []*models.MechanicalPress
	err := s.DB.Model(tool).Association("MechanicalPresses").Find(&mechanicalPresses)
	if err != nil {
		return nil, err
	}
	dtos := models.Map(mechanicalPresses, func(mp *models.MechanicalPress) *models.MechanicalPressDTO {
		return mp.ToDTO()
	})
	return dtos, nil
}

func (s *ToolService) GetToolRepairRecords(tool *models.Tool) ([]*models.ToolRepairRecordDTO, error) {
	var toolRepairRecords []*models.ToolRepairRecord
	err := s.DB.Model(tool).Association("ToolRepairRecords").Find(&toolRepairRecords)
	if err != nil {
		return nil, err
	}
	dtos := models.Map(toolRepairRecords, func(trr *models.ToolRepairRecord) *models.ToolRepairRecordDTO {
		return trr.ToDTO()
	})

	return dtos, nil
}
