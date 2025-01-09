package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Connect(host, port string) error {
	_ = host
	_ = port

	logs := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Writer คือ os.Stdout
		logger.Config{
			SlowThreshold:             time.Second, // เกณฑ์การเตือน query ที่ช้า
			LogLevel:                  logger.Info, // ระดับ log (Info, Warn, Error)
			IgnoreRecordNotFoundError: true,        // ละเว้น error record not found
			Colorful:                  true,        // แสดง log มีสี
		},
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"agent",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logs,
	})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&Account{},
		&Role{},
		&Permission{},
		&RolePermission{},
		&Layer{},
	)

	if err != nil {
		panic("Failed to migrate database schema")
	}

	DB = db
	fmt.Println("Database connected!")
	return nil
}
