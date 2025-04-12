package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{db: db}
}

func (s *RoleService) CreateRole(role *models.Role) (*models.RoleDTO, error) {
	err := s.db.Create(role).Error
	if err != nil {
		return nil, err
	}
	return role.ToDTO(), nil
}

func (s *RoleService) GetAllRoles() ([]*models.RoleDTO, error) {
	var roles []*models.Role
	if err := s.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	dtos := models.Map(roles, func(role *models.Role) *models.RoleDTO {
		return role.ToDTO()
	})
	return dtos, nil
}

func (s *RoleService) GetRoleByID(id uint) (*models.RoleDTO, error) {
	var role models.Role
	if err := s.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return role.ToDTO(), nil
}
