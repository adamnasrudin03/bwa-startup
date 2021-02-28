package controllers

import (
	"bwa-startup/auth"
	"bwa-startup/helpers"
	"bwa-startup/users"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userController struct {
	userService users.Service
	authService auth.Service
}

func NewUserController (userService users.Service, authService auth.Service) *userController{
	return &userController{userService, authService}
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

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		response := helpers.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(user, token)

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

	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helpers.APIResponse("Loginfailed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(loggedInUser, token)

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

	//get context isinya user yang telah di isi dari authMiddleware
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID
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


func (h *userController) FetchUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(users.User)

	formatter := users.FormatUser(currentUser, "")

	response := helpers.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
