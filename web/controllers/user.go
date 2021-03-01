package controllers

import (
	"bwa-startup/users"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userController struct {
	userService users.Service
}

func NewUserController(userService users.Service) *userController {
	return &userController{userService}
}

func (h *userController) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		//later
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}