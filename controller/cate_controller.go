package controller

import (
	"ChoTot/entity"
	"ChoTot/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICateController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type cateController struct {
	cateSvc service.ICateService
}

func NewCateController(cateSvc service.ICateService) ICateController {
	return &cateController{cateSvc: cateSvc}
}

func (ctrl *cateController) GetAll(c *gin.Context) {

	cates, err := ctrl.cateSvc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cates})

}

func (ctrl *cateController) GetById(c *gin.Context) {

	cate, err := ctrl.cateSvc.GetById(c.Param(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cate})

}

func (ctrl *cateController) Create(c *gin.Context) {
	var cateInfo entity.SubCategories

	//binding
	if err := c.ShouldBindJSON(&cateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//update
	if err := ctrl.cateSvc.Create(&cateInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//ok
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (ctrl *cateController) Update(c *gin.Context) {
	var cateInfo entity.SubCategories

	//binding
	if err := c.ShouldBindJSON(&cateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//update
	if err := ctrl.cateSvc.Update(&cateInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//ok
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (ctrl *cateController) Delete(c *gin.Context) {

	//func from user_controller
	if user := GetUserFromToken("token"); !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": FORBIDDEN})
		log.Fatalln("User is not ADMIN")
		return
	}

	if err := ctrl.cateSvc.Del(c.Param(ID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatalln("Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
