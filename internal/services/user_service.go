package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetUserById(id uint) (*models.UserDTO, error) {
	var user *models.User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user.ToDTO(), nil
}

func (s *UserService) GetUsers() ([]*models.UserDTO, error) {
	var users []*models.User
	err := s.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(users, func(user *models.User) *models.UserDTO {
		return user.ToDTO()
	})
	return dtos, nil
}
