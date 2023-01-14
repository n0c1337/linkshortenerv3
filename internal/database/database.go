package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (db *gorm.DB, err error) {
	dsn := "host=localhost user=root password=root dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return
}
