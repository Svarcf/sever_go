package services

import (
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type ToolRepairRecordService struct {
	DB *gorm.DB
}

func NewToolRepairRecordService(db *gorm.DB) *ToolRepairRecordService {
	return &ToolRepairRecordService{DB: db}
}

func (s *ToolRepairRecordService) GetToolRepairRecordById(id uint) (*models.ToolRepairRecord, error) {
	var toolRepairRecord *models.ToolRepairRecord
	err := s.DB.First(&toolRepairRecord, id).Error
	if err != nil {
		return nil, err
	}
	return toolRepairRecord, nil
}

func (s *ToolRepairRecordService) GetToolRepairRecords() ([]*models.ToolRepairRecord, error) {
	var toolRepairRecords []*models.ToolRepairRecord
	err := s.DB.Find(&toolRepairRecords).Error
	if err != nil {
		return nil, err
	}
	return toolRepairRecords, nil
}
