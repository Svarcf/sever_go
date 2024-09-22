package models

import (
	"gorm.io/gorm"
	"time"
)

type ToolRepairRecord struct {
	gorm.Model
	DateStarted            time.Time
	DateEnded              time.Time
	RepairDescription      string
	MalfunctionDescription string
	DeadlineDate           time.Time
	Material               string
	TimeSpent              uint
	ExternalServices       string
	Note                   string
	ToolID                 uint
	Tool                   Tool
	UserID                 uint
	RepairedBy             User `gorm:"foreignKey:UserID"`
}
