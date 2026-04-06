package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"restapi/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=pass dbname=tododb port=5432 sslmode=disable"
	}

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Khong the ket noi database: ", err)
	}

	DB.AutoMigrate(&models.Todo{})
}
