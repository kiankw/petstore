package dal

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatalln("MySQL not found")
	}
	DB = db
}

func GetDB(ctx context.Context) *gorm.DB {
	return DB
}
