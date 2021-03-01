package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (h *userController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "user_index.html", nil)
}