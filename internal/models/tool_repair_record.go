package models

import (
	"time"
)

type ToolRepairRecord struct {
	Id                     uint `gorm:"primarykey"`
	DateStarted            time.Time
	DateEnded              time.Time
	RepairDescription      string
	MalfunctionDescription string
	DeadlineDate           time.Time
	MaterialID             uint
	TimeSpent              int
	ExternalServices       string
	Note                   string
	ToolID                 uint
	Tool                   Tool
	UserID                 uint
	RepairedBy             User `gorm:"foreignKey:UserID"`
}

func NewToolRepairRecord(dateStarted time.Time, dateEnded time.Time, repairDescription string,
	malfunctionDescription string, deadlineDate time.Time, materialID *uint, timeSpent int, externalServices string,
	note string, toolID uint, userID uint) *ToolRepairRecord {
	return &ToolRepairRecord{
		DateStarted:            dateStarted,
		DateEnded:              dateEnded,
		RepairDescription:      repairDescription,
		MalfunctionDescription: malfunctionDescription,
		DeadlineDate:           deadlineDate,
		MaterialID:             *materialID,
		TimeSpent:              timeSpent,
		ExternalServices:       externalServices,
		Note:                   note,
		ToolID:                 toolID,
		UserID:                 userID,
	}
}
