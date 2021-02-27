package controllers

import (
	"bwa-startup/helpers"
	"bwa-startup/users"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userController struct {
	userService users.Service
}

func NewUserController (userService users.Service) *userController{
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context){
	//Tangkap input dari user
	//Map input user ke struct RegisterUserInput
	//Struct diatas di pasrsing sebagai parameter service

	var input users.RegisterUserInput
	err := c.BindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helpers.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(user, "JWT token")

	response := helpers.APIResponse("Account has ben registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userController) Login(c *gin.Context) {
	//step login
	//user memasukan input (email & password)
	//input ditangkap handler/controller
	//mapping dati input user ke input struct
	//input struct passing ke service 
	//di service mencari dgn bantuan repositori user dengan email x
	//mencocokan password
	var input users.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := users.FormatUser(loggedInUser, "JWT token")

	response := helpers.APIResponse("Login Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userController)  CheckEmailAvailability(c *gin.Context){
	//Input email dari user
	//input email di mapping ke struct input
	//struct input di passing ke service
	//service akan manggil repository - email masih tersedia atau tidak
	//repositori - db
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Email Checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helpers.APIResponse("Email Checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helpers.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userController)  UploadAvatar(c *gin.Context){
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//userID harusnya dapat dari JWT
	userID := 1
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helpers.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
