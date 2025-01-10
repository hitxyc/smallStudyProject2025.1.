package mapper

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"your_project/config"
)

var _db *gorm.DB

func InitDataBase(config *config.DatabaseConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	_db = db
}

func CloseDataBase() {
	sqlDB, err := _db.DB()
	if err != nil {
		log.Println("Error getting DB:", err)
		return
	}

	// 关闭数据库连接
	err = sqlDB.Close()
	if err != nil {
		log.Println("Error closing DB connection:", err)
	}
}

func GetDB() *gorm.DB {
	return _db
}
