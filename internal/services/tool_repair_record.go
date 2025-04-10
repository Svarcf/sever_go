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

func (s *ToolRepairRecordService) CreateToolRepairRecord(toolRepairRecord *models.ToolRepairRecord) (*models.ToolRepairRecordDTO, error) {
	err := s.DB.Create(toolRepairRecord).Error
	if err != nil {
		return nil, err
	}
	return toolRepairRecord.ToDto(), nil
}

func (s *ToolRepairRecordService) GetToolRepairRecordById(id uint) (*models.ToolRepairRecordDTO, error) {
	var toolRepairRecord *models.ToolRepairRecord
	err := s.DB.First(&toolRepairRecord, id).Error
	if err != nil {
		return nil, err
	}
	return toolRepairRecord.ToDto(), nil
}

func (s *ToolRepairRecordService) GetToolRepairRecords() ([]*models.ToolRepairRecordDTO, error) {
	var toolRepairRecords []*models.ToolRepairRecord
	err := s.DB.Find(&toolRepairRecords).Error
	if err != nil {
		return nil, err
	}
	dtos := models.Map(toolRepairRecords, func(trr *models.ToolRepairRecord) *models.ToolRepairRecordDTO {
		return trr.ToDto()
	})

	return dtos, nil
}
