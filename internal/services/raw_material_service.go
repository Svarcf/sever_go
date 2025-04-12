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

func (s *RawMaterialService) CreateRawMaterial(rawMaterial *models.RawMaterial) (*models.RawMaterialDTO, error) {
	err := s.DB.Create(rawMaterial).Error
	if err != nil {
		return nil, err
	}
	return rawMaterial.ToDTO(), nil
}

func (s *RawMaterialService) GetRawMaterialById(id uint) (*models.RawMaterialDTO, error) {
	var rawMaterial *models.RawMaterial
	err := s.DB.First(&rawMaterial, id).Error
	if err != nil {
		return nil, err
	}
	return rawMaterial.ToDTO(), nil
}

func (s *RawMaterialService) GetRawMaterials() ([]*models.RawMaterialDTO, error) {
	var rawMaterials []*models.RawMaterial
	err := s.DB.Find(&rawMaterials).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(rawMaterials, func(rm *models.RawMaterial) *models.RawMaterialDTO {
		return rm.ToDTO()
	})
	return dtos, nil
}
