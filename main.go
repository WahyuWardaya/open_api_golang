package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct{
	ID 			uint		`gorm:"column:id;primary_key"`
	Name 		string		`gorm:"column:name"`
	Email 		string		`gorm:"column:email"`
	Age 		int			`gorm:"column:age"`
	CreatedAt 	time.Time	`gorm:"column:created_at"`
	UpdatedAt 	time.Time	`gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}


func main() {
	dsn := "root:@tcp(localhost:3306)/openapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("Tidak berhasil konek ke database")
	}
	router := gin.Default()
	
	router.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(http.StatusOK, gin.H{"data":users})
	})

	router.POST("/users", func(c *gin.Context){
		var newUsers User

		if err := c.ShouldBindJSON(&newUsers); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newUsers)
		c.JSON(http.StatusCreated, gin.H{"data": newUsers})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
	
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

		
	router.Run(":3000")
}
