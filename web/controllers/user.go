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
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

func (h *userController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "user_new.html", nil)
}


func (h *userController) Create(c *gin.Context) {
	var input users.FormCreateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		//skip
	}

	registerInput := users.RegisterUserInput{}
	registerInput.Name = input.Name
	registerInput.Email = input.Email
	registerInput.Occupation = input.Occupation
	registerInput.Password = input.Password

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		//skip
	}

	c.Redirect(http.StatusFound, "/users")
}
