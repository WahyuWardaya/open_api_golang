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

		
	router.Run(":3000")
}
