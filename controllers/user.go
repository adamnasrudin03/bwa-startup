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
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	formatter := users.FormatUser(user, "JWT token")

	response := helpers.APIResponse("Account has ben registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}