package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type ToolTypeService struct {
	DB *gorm.DB
}

func NewToolTypeService(db *gorm.DB) *ToolTypeService {
	return &ToolTypeService{DB: db}
}

func (s *ToolTypeService) CreateToolType(toolType *models.ToolType) (*models.ToolTypeDTO, error) {
	err := s.DB.Create(toolType).Error
	if err != nil {
		return nil, err
	}
	return toolType.ToDTO(), nil
}

func (s *ToolTypeService) UpdateToolType(toolType *models.ToolType) (*models.ToolTypeDTO, error) {
	err := s.DB.Save(toolType).Error
	if err != nil {
		return nil, err
	}
	return toolType.ToDTO(), nil
}

func (s *ToolTypeService) GetToolTypeByCode(code string) (*models.ToolTypeDTO, error) {
	var toolType *models.ToolType
	err := s.DB.Where("code = ?", code).First(&toolType).Error
	if err != nil {
		return nil, err
	}
	return toolType.ToDTO(), nil
}

func (s *ToolTypeService) GetToolTypeById(id uint) (*models.ToolTypeDTO, error) {
	var toolType *models.ToolType
	err := s.DB.First(&toolType, id).Error
	if err != nil {
		return nil, err
	}
	return toolType.ToDTO(), nil
}

func (s *ToolTypeService) GetToolTypes() ([]*models.ToolTypeDTO, error) {
	var toolTypes []*models.ToolType
	err := s.DB.Find(&toolTypes).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(toolTypes, func(tt *models.ToolType) *models.ToolTypeDTO {
		return tt.ToDTO()
	})
	return dtos, nil
}

func (s *ToolTypeService) GetTools(toolType *models.ToolType) ([]*models.ToolDTO, error) {
	var tools []*models.Tool
	err := s.DB.Model(toolType).Association("Tools").Find(&tools)
	if err != nil {
		return nil, err
	}
	dtos := models.Map(tools, func(t *models.Tool) *models.ToolDTO {
		return t.ToDTO()
	})
	return dtos, nil
}
