package main

import (
	"bwa-startup/auth"
	"bwa-startup/controllers"
	"bwa-startup/users"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastratup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection Database Success!")

	userRepository := users.NewRepository(db)

	userService := users.NewService(userRepository)
	authService := auth.NewService()

	userController := controllers.NewUserController(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userController.RegisterUser)
	api.POST("/sessions", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", userController.UploadAvatar)
	router.Run()

	
}