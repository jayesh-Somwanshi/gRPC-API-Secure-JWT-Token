package config

import (
	"SecureAPIWithgrpc/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func InitDB() {
// 	dsn := "root:root@tcp(127.0.0.1:3306)/grpcDatabase?charset=utf8mb4&parseTime=True&loc=Local"

// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to database")
// 	}
// 	DB.AutoMigrate(&model.Employee{})
// }

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/grpcDatabase?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Assign to global `DB`
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate models
	if err := DB.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
