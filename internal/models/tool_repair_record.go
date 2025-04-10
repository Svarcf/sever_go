package models

import (
	"database/sql"
	"time"
)

type ToolRepairRecord struct {
	Id                     uint `gorm:"primarykey"`
	DateStarted            time.Time
	DateEnded              sql.NullTime
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

type ToolRepairRecordDTO struct {
	Id                     uint       `json:"id"`
	DateStarted            time.Time  `json:"date_started"`
	DateEnded              *time.Time `json:"date_ended"`
	RepairDescription      string     `json:"repair_description"`
	MalfunctionDescription string     `json:"malfunction_description"`
	DeadlineDate           time.Time  `json:"deadline_date"`
	MaterialID             uint       `json:"material_id"`
	TimeSpent              int        `json:"time_spent"`
	ExternalServices       string     `json:"external_services"`
	Note                   string     `json:"note"`
	ToolID                 uint       `json:"tool_id"`
	Tool                   Tool       `json:"tool"`
	UserID                 uint       `json:"user_id"`
	RepairedBy             User       `json:"repaired_by"`
}

func (trrDto ToolRepairRecordDTO) ToModel() *ToolRepairRecord {
	trr := &ToolRepairRecord{
		Id:                     trrDto.Id,
		DateStarted:            trrDto.DateStarted,
		RepairDescription:      trrDto.RepairDescription,
		MalfunctionDescription: trrDto.MalfunctionDescription,
		DeadlineDate:           trrDto.DeadlineDate,
		MaterialID:             trrDto.MaterialID,
		TimeSpent:              trrDto.TimeSpent,
		ExternalServices:       trrDto.ExternalServices,
		Note:                   trrDto.Note,
		ToolID:                 trrDto.ToolID,
		UserID:                 trrDto.UserID,
	}

	if trrDto.DateEnded == nil {
		trr.DateEnded = sql.NullTime{Valid: false}
	} else {
		trr.DateEnded = sql.NullTime{Valid: true, Time: *trrDto.DateEnded}

	}
	return trr
}

func (trr ToolRepairRecord) ToDto() *ToolRepairRecordDTO {
	return &ToolRepairRecordDTO{
		Id:                     trr.Id,
		DateStarted:            trr.DateStarted,
		DateEnded:              &trr.DateEnded.Time,
		RepairDescription:      trr.RepairDescription,
		MalfunctionDescription: trr.MalfunctionDescription,
		DeadlineDate:           trr.DeadlineDate,
		MaterialID:             trr.MaterialID,
		TimeSpent:              trr.TimeSpent,
		ExternalServices:       trr.ExternalServices,
		Note:                   trr.Note,
		ToolID:                 trr.ToolID,
		UserID:                 trr.UserID,
	}
}

func NewToolRepairRecord(dateStarted time.Time, dateEnded sql.NullTime, repairDescription string,
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
