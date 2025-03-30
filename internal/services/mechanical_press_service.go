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

func (s *MechanicalPressService) CreateMechanicalPress(mechanicalPress *models.MechanicalPress) (*models.MechanicalPress, error) {
	err := s.DB.Create(mechanicalPress).Error
	if err != nil {
		return nil, err
	}
	return mechanicalPress, nil
}

func (s *MechanicalPressService) GetMechanicalPressById(id uint) (*models.MechanicalPress, error) {
	var mechanicalPress *models.MechanicalPress
	err := s.DB.First(&mechanicalPress, id).Error
	if err != nil {
		return nil, err
	}
	return mechanicalPress, nil
}

func (s *MechanicalPressService) GetMechanicalPresses() ([]*models.MechanicalPress, error) {
	var mechanicalPresses []*models.MechanicalPress
	err := s.DB.Find(&mechanicalPresses).Error
	if err != nil {
		return nil, err
	}
	return mechanicalPresses, nil
}
