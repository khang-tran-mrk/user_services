package main

import (
	"ChoTot/config"
	"ChoTot/controller"
	product_grpc "ChoTot/grpc"
	"ChoTot/middleware"
	"ChoTot/repository"
	"ChoTot/service"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDatabase()

	userRepo       repository.IUserRepository = repository.NewUserRepository(db)
	userService    service.IUserService       = service.NewUserService(userRepo)
	userController controller.IUserController = controller.NewUserController(userService)

	cateRepo       repository.ICateRepository = repository.NewCateRepository(db)
	cateService    service.ICateService       = service.NewCateService(cateRepo)
	cateController controller.ICateController = controller.NewCateController(cateService)
)

const (
	ADDR                 = "localhost"
	GRPC_PORT            = ":5000"
	API_USERSERVICE_PORT = ":8080"
)

func main() {
	//get grpc connection from product service
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(ADDR+GRPC_PORT, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	productClient := product_grpc.NewProductClient(conn)

	defer config.CloseDatabase(db)
	r := gin.Default()

	userRoutes := r.Group("/user").Use(middleware.Auth())
	{
		userRoutes.GET("/ping", middleware.Pong)
		userRoutes.GET("/profile/:id", userController.Profile)
		userRoutes.POST("/update/:id", userController.Update)

		userRoutes.GET("/products", userController.GetProductsByUserId(productClient))
		userRoutes.POST("/products", userController.PostNewProduct(productClient))
	}

	cateRoutes := r.Group("/cate")
	{
		cateRoutes.GET("/", cateController.GetAll)
		cateRoutes.GET("/:id", cateController.GetById)
	}

	r.Run(API_USERSERVICE_PORT)
}
