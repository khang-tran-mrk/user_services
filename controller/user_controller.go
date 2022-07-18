package controller

import (
	"ChoTot/entity"
	product_grpc "ChoTot/grpc"
	"ChoTot/helpers"
	"ChoTot/service"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	grpcMetadata "google.golang.org/grpc/metadata"
)

const (
	ID        = "id"
	FORBIDDEN = "Forbidden"
	OFFSET    = "p"

	HEADER_KEY_AUTHORIZATION = "Authorization"
	NOT_CONTAIN_ACCESS_TOKEN = "request does not contain an access token"
)

type IUserController interface {
	Profile(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetProductsByUserId(pClient product_grpc.ProductClient) gin.HandlerFunc
	PostNewProduct(pClient product_grpc.ProductClient) gin.HandlerFunc
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

func (ctrl *userController) GetProductsByUserId(pClient product_grpc.ProductClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		//set token to context
		grpc_ctx, cancel := getContext(c)
		defer cancel()
		if grpc_ctx == nil {
			return
		}
		//get offset
		offset := c.GetInt(OFFSET)
		req := product_grpc.GetProductRequest{Offset: int32(offset)}
		res, err := pClient.GetProduct(grpc_ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": res,
		})
	}
}

func (ctrl *userController) PostNewProduct(pClient product_grpc.ProductClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		//set token to context
		grpc_ctx, cancel := getContext(c)
		defer cancel()
		if grpc_ctx == nil {
			return
		}

		//binding
		var p_input entity.Product
		if err := c.ShouldBindJSON(&p_input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		//send req to get res from Product Service
		req := helpers.P_EntityToGrpcCreateReq(&p_input)
		res, err := pClient.CreateProduct(grpc_ctx, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": res,
		})
	}
}

func getContext(c *gin.Context) (context.Context, context.CancelFunc) {
	grpc_ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	token := c.GetHeader(HEADER_KEY_AUTHORIZATION)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": NOT_CONTAIN_ACCESS_TOKEN,
		})
		return nil, cancel
	}
	grpc_ctx = grpcMetadata.AppendToOutgoingContext(grpc_ctx, "authorization", token)

	return grpc_ctx, cancel
}
