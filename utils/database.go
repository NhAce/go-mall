package utils

import (
	"fmt"
	"go-mall/config"
	"go-mall/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	dbConfig := config.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.Charset,
		dbConfig.ParseTime,
		dbConfig.Loc,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	// 自动迁移表结构
	err = AutoMigrate()
	if err != nil {
		return err
	}

	log.Println("数据库连接成功")
	return nil
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	return DB.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Product{},
		&model.Cart{},
		&model.Address{},
		&model.Order{},
		&model.OrderItem{},
	)
} 