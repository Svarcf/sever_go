package database

import (
	"fmt"
	"github.com/Svarcf/sever_go/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect to mysql database
func InitDB() (*gorm.DB, error) {
	cfg, err := config.LoadConfig()
	//take username, password, host, port, database from config file using fmt.Sprintf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
