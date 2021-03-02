package main

import (
	"bwa-startup/auth"
	"bwa-startup/campaign"
	"bwa-startup/controllers"
	"bwa-startup/helpers"
	"bwa-startup/payment"
	"bwa-startup/transaction"
	"bwa-startup/users"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	webController "bwa-startup/web/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	campaignRepository := campaign.NewRepository(db)
	transactinRepository := transaction.NewRepository(db)

	userService := users.NewService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactinRepository, campaignRepository, paymentService)

	userController := controllers.NewUserController(userService, authService)
	campaignController := controllers.NewCampaignController(campaignService)
	transactionController := controllers.NewTransactionController(transactionService)

	userWebController := webController.NewUserController(userService)
	campaignWebController := webController.NewcampaignController(campaignService, userService)
	transactionWebController := webController.NewTransactionController(transactionService)

	router := gin.Default()
	router.Use(cors.Default())

	store := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("bwastartup", store))

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	api := router.Group("/api/v1")

	api.POST("/users", userController.RegisterUser)
	api.POST("/sessions", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userController.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userController.FetchUser)

	api.GET("/campaigns", campaignController.GetCampaigns)
	api.GET("/campaigns/:id", campaignController.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignController.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignController.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignController.UploadCampaignImage)


	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionController.GetCampaignTransaction)
	api.GET("/transactions", authMiddleware(authService, userService), transactionController.GetUserTransaction)
	api.POST("/transactions", authMiddleware(authService, userService), transactionController.CreateTransaction)
	api.POST("/transactions/notification", transactionController.GetNotification)


	router.GET ("/users", authAdminMiddleware(), userWebController.Index)
	router.GET ("/users/new", authAdminMiddleware(), userWebController.New)
	router.POST ("/users", authAdminMiddleware(), userWebController.Create)
	router.GET ("/users/edit/:id", authAdminMiddleware(), userWebController.Edit)
	router.POST ("/users/update/:id", authAdminMiddleware(), userWebController.Update)
	router.GET ("/users/avatar/:id", authAdminMiddleware(), userWebController.NewAvatar)
	router.POST ("/users/avatar/:id", authAdminMiddleware(), userWebController.CreateAvatar)

	router.GET ("/campaigns", authAdminMiddleware(), campaignWebController.Index)
	router.GET ("/campaigns/new", authAdminMiddleware(), campaignWebController.New)
	router.POST ("/campaigns", authAdminMiddleware(), campaignWebController.Create)
	router.GET ("/campaigns/image/:id", authAdminMiddleware(), campaignWebController.NewImage)
	router.POST ("/campaigns/image/:id", authAdminMiddleware(), campaignWebController.CreateImage)
	router.GET ("/campaigns/edit/:id", authAdminMiddleware(), campaignWebController.Edit)
	router.POST ("/campaigns/update/:id", authAdminMiddleware(), campaignWebController.Update)
	router.GET ("/campaigns/show/:id", authAdminMiddleware(), campaignWebController.Show)

	router.GET ("/transactions", authAdminMiddleware(), transactionWebController.Index)

	router.Run()
}



func authMiddleware(authService auth.Service, userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		playload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(playload["user_id"].(float64))

		user, err := userService.GetUserById(userID)
		if err != nil {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//set context isinya user
		c.Set("currentUser", user)
	}
}

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions := sessions.Default(c)

		userIDSessions := sessions.Get("userID")

		if userIDSessions == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}

	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
