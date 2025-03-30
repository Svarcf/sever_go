package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type FileService struct {
	DB *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{DB: db}
}

func (s *FileService) createFile(file *models.File) (*models.File, error) {
	err := s.DB.Create(file).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *FileService) GetFileById(id uint) (*models.File, error) {
	var file *models.File
	err := s.DB.First(&file, id).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *FileService) GetFiles() ([]*models.File, error) {
	var files []*models.File
	err := s.DB.Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}
