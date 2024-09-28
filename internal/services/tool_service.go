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

func (s *ToolService) GetToolById(id uint) (*models.Tool, error) {
	var tool *models.Tool
	err := s.DB.First(&tool, id).Error
	if err != nil {
		return nil, err
	}
	return tool, nil
}

func (s *ToolService) GetTools() ([]*models.Tool, error) {
	var tools []*models.Tool
	err := s.DB.Find(&tools).Error
	if err != nil {
		return nil, err
	}
	return tools, nil
}

func (s *ToolService) GetToolByCode(code string) (*models.Tool, error) {
	var tool *models.Tool
	err := s.DB.Where("code = ?", code).First(&tool).Error
	if err != nil {
		return nil, err
	}
	return tool, nil
}

func (s *ToolService) GetRawMaterials(tool *models.Tool) ([]*models.RawMaterial, error) {
	var rawMaterials []*models.RawMaterial
	err := s.DB.Model(tool).Association("RawMaterials").Find(&rawMaterials)
	if err != nil {
		return nil, err
	}
	return rawMaterials, nil
}

func (s *ToolService) GetStandardParts(tool *models.Tool) ([]*models.StandardPart, error) {
	var standardParts []*models.StandardPart
	err := s.DB.Model(tool).Association("StandardParts").Find(&standardParts)
	if err != nil {
		return nil, err
	}
	return standardParts, nil
}

func (s *ToolService) GetMechanicalPresses(tool *models.Tool) ([]*models.MechanicalPress, error) {
	var mechanicalPresses []*models.MechanicalPress
	err := s.DB.Model(tool).Association("MechanicalPresses").Find(&mechanicalPresses)
	if err != nil {
		return nil, err
	}
	return mechanicalPresses, nil
}

func (s *ToolService) GetToolRepairRecords(tool *models.Tool) ([]*models.ToolRepairRecord, error) {
	var toolRepairRecords []*models.ToolRepairRecord
	err := s.DB.Model(tool).Association("ToolRepairRecords").Find(&toolRepairRecords)
	if err != nil {
		return nil, err
	}
	return toolRepairRecords, nil
}
