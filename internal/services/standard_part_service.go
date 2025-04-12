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

func (s *StandardPartService) CreateStandardPart(standardPart *models.StandardPart) (*models.StandardPartDTO, error) {
	err := s.DB.Create(standardPart).Error
	if err != nil {
		return nil, err
	}
	return standardPart.ToDTO(), nil
}

func (s *StandardPartService) GetStandardPartById(id uint) (*models.StandardPartDTO, error) {
	var standardPart *models.StandardPart
	err := s.DB.First(&standardPart, id).Error
	if err != nil {
		return nil, err
	}
	return standardPart.ToDTO(), nil
}

func (s *StandardPartService) GetStandardParts() ([]*models.StandardPartDTO, error) {
	var standardParts []*models.StandardPart
	err := s.DB.Find(&standardParts).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(standardParts, func(sp *models.StandardPart) *models.StandardPartDTO {
		return sp.ToDTO()
	})
	return dtos, nil
}
