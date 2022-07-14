package controller

import (
	"ChoTot/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICateController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
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
