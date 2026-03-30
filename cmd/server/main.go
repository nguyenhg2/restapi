package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}

var DB *gorm.DB

func connectDB() {
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

	DB.AutoMigrate(&Todo{})
}

func main() {
	connectDB()

	r := gin.Default()

	r.GET("/api/v1/todos", func(c *gin.Context) {
		var todos []Todo
		DB.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	r.POST("/api/v1/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.Create(&todo)
		c.JSON(http.StatusCreated, todo)
	})

	r.DELETE("/api/v1/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo
		if err := DB.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo khong ton tai"})
			return
		}
		DB.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"message": "Xoa thanh cong"})
	})

	r.Run(":8080")
}
