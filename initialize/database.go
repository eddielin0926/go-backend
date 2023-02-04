package initialize

import (
	"go-backend/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Fail to connect the database")
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
