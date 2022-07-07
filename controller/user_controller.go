package controller

import (
	"ChoTot/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Profile(c *gin.Context)
}

type userController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &userController{userService: userService}
}

func (ctrl *userController) Profile(c *gin.Context) {
	user, err := ctrl.userService.UserProfile(1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
