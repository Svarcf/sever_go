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

func (us *UserService) GetUser(id uint) (*models.User, error) {
	var user *models.User
	err := us.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := us.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetUserRole(obj *models.User) (*models.Role, error) {
	var role *models.Role
	err := us.DB.First(&role, obj.RoleID).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}
