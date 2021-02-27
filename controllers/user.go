package controllers

import (
	"bwa-startup/helpers"
	"bwa-startup/users"
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

func (h *userController) login(c *gin.Context) {
	//step login
	//user memasukan input (email & password)
	//input ditangkap handler/controller
	//mapping dati input user ke input struct
	//input struct passing ke service 
	//di service mencari dgn bantuan repositori user dengan email x
	//mencocokan password
}