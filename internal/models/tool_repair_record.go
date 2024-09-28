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
	Material               string
	TimeSpent              int
	ExternalServices       string
	Note                   string
	ToolID                 uint
	Tool                   Tool
	UserID                 uint
	RepairedBy             User `gorm:"foreignKey:UserID"`
}
