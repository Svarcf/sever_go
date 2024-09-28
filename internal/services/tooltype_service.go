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

func (s *ToolTypeService) GetToolTypeByCode(code string) (*models.ToolType, error) {
	var toolType *models.ToolType
	err := s.DB.Where("code = ?", code).First(&toolType).Error
	if err != nil {
		return nil, err
	}
	return toolType, nil
}

func (s *ToolTypeService) GetToolTypeById(id uint) (*models.ToolType, error) {
	var toolType *models.ToolType
	err := s.DB.First(&toolType, id).Error
	if err != nil {
		return nil, err
	}
	return toolType, nil
}

func (s *ToolTypeService) GetToolTypes() ([]*models.ToolType, error) {
	var toolTypes []*models.ToolType
	err := s.DB.Find(&toolTypes).Error
	if err != nil {
		return nil, err
	}
	return toolTypes, nil
}
