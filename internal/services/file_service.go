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

func (s *FileService) CreateFile(file *models.File) (*models.FileDTO, error) {
	err := s.DB.Create(file).Error
	if err != nil {
		return nil, err
	}
	return file.ToDTO(), nil
}

func (s *FileService) GetFileById(id uint) (*models.FileDTO, error) {
	var file *models.File
	err := s.DB.First(&file, id).Error
	if err != nil {
		return nil, err
	}
	return file.ToDTO(), nil
}

func (s *FileService) GetFiles() ([]*models.FileDTO, error) {
	var files []*models.File
	err := s.DB.Find(&files).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(files, func(file *models.File) *models.FileDTO {
		return file.ToDTO()
	})

	return dtos, nil
}

func (s *FileService) UpdateFile(file *models.File) (*models.FileDTO, error) {
	err := s.DB.Save(file).Error
	if err != nil {
		return nil, err
	}
	return file.ToDTO(), nil
}
