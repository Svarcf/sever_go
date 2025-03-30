package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type RawMaterialService struct {
	DB *gorm.DB
}

func NewRawMaterialService(db *gorm.DB) *RawMaterialService {
	return &RawMaterialService{DB: db}
}

func (s *RawMaterialService) CreateRawMaterial(rawMaterial *models.RawMaterial) (*models.RawMaterial, error) {
	err := s.DB.Create(rawMaterial).Error
	if err != nil {
		return nil, err
	}
	return rawMaterial, nil
}

func (s *RawMaterialService) GetRawMaterialById(id uint) (*models.RawMaterial, error) {
	var rawMaterial *models.RawMaterial
	err := s.DB.First(&rawMaterial, id).Error
	if err != nil {
		return nil, err
	}
	return rawMaterial, nil
}

func (s *RawMaterialService) GetRawMaterials() ([]*models.RawMaterial, error) {
	var rawMaterials []*models.RawMaterial
	err := s.DB.Find(&rawMaterials).Error
	if err != nil {
		return nil, err
	}
	return rawMaterials, nil
}
