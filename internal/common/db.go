package common

import (
	"fmt"
	"github.com/Svarcf/sever_go/internal/config"
	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB { return DB }

func InitDB() *gorm.DB {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if cfg.DB_INIT {
		db.AutoMigrate(&models.ToolType{}, &models.Tool{}, &models.Role{}, &models.User{},
			&models.File{}, &models.StandardPart{}, &models.MechanicalPress{},
			&models.ToolRepairRecord{}, &models.RawMaterial{})
	}

	DB = db
	return db
}
