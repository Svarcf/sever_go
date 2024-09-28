package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type StandardPartService struct {
	DB *gorm.DB
}

func NewStandardPartService(db *gorm.DB) *StandardPartService {
	return &StandardPartService{DB: db}
}

func (s *StandardPartService) GetStandardPartById(id uint) (*models.StandardPart, error) {
	var standardPart *models.StandardPart
	err := s.DB.First(&standardPart, id).Error
	if err != nil {
		return nil, err
	}
	return standardPart, nil
}

func (s *StandardPartService) GetStandardParts() ([]*models.StandardPart, error) {
	var standardParts []*models.StandardPart
	err := s.DB.Find(&standardParts).Error
	if err != nil {
		return nil, err
	}
	return standardParts, nil
}
