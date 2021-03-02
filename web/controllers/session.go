package controllers

import (
	"bwa-startup/users"
	"net/http"

	"github.com/gin-gonic/gin"
)


type sessionController struct {
	userService users.Service
}
func NewSessionController (userService users.Service) *sessionController{
	return &sessionController{userService}
}

func (h *sessionController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}