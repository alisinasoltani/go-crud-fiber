package database

import (
	"log"

	models "github.com/alisinasoltani/goFiber/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting to the Database\n", err.Error())
	}
	log.Println("Database Connected Successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations...")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.User{})
	Database = DbInstance{Db: db}
}
