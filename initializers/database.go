package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() {
	dsn := "go_user:go_pass@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
