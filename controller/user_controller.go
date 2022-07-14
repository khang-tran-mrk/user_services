package controller

import (
	"ChoTot/entity"
	"ChoTot/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ID        = "id"
	FORBIDDEN = "Forbidden"
)

type IUserController interface {
	Profile(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &userController{userService: userService}
}

func (ctrl *userController) Profile(c *gin.Context) {
	user, err := ctrl.userService.GetById(c.Param(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (ctrl *userController) Update(c *gin.Context) {
	var userInfo entity.User

	//binding
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//update
	if err := ctrl.userService.Update(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//ok
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (ctrl *userController) Delete(c *gin.Context) {
	if user := GetUserFromToken("token"); !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": FORBIDDEN})
		log.Fatalln("User is not ADMIN")
		return
	}

	if err := ctrl.userService.Del(c.Param(ID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatalln("Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func GetUserFromToken(s string) *entity.User {
	//todo things
	return &entity.User{IsAdmin: true}
}
