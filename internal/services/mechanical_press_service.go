package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type MechanicalPressService struct {
	DB *gorm.DB
}

func NewMechanicalPressService(db *gorm.DB) *MechanicalPressService {
	return &MechanicalPressService{DB: db}
}

func (s *MechanicalPressService) CreateMechanicalPress(mechanicalPress *models.MechanicalPress) (*models.MechanicalPressDTO, error) {
	err := s.DB.Create(mechanicalPress).Error
	if err != nil {
		return nil, err
	}
	return mechanicalPress.ToDTO(), nil
}

func (s *MechanicalPressService) UpdateMechanicalPress(mechanicalPress *models.MechanicalPress) (*models.MechanicalPressDTO, error) {
	err := s.DB.Save(mechanicalPress).Error
	if err != nil {
		return nil, err
	}
	return mechanicalPress.ToDTO(), nil
}

func (s *MechanicalPressService) GetMechanicalPressById(id uint) (*models.MechanicalPressDTO, error) {
	var mechanicalPress *models.MechanicalPress
	err := s.DB.First(&mechanicalPress, id).Error
	if err != nil {
		return nil, err
	}
	return mechanicalPress.ToDTO(), nil
}

func (s *MechanicalPressService) GetMechanicalPresses() ([]*models.MechanicalPressDTO, error) {
	var mechanicalPresses []*models.MechanicalPress
	err := s.DB.Find(&mechanicalPresses).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(mechanicalPresses, func(mp *models.MechanicalPress) *models.MechanicalPressDTO {
		return mp.ToDTO()
	})

	return dtos, nil
}
