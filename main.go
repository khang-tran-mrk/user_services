package main

import (
	"ChoTot/config"
	"ChoTot/controller"
	"ChoTot/middleware"
	"ChoTot/repository"
	"ChoTot/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDatabase()

	userRepo repository.IUserRepository = repository.NewUserRepository(db)

	userService service.IUserService = service.NewUserService(userRepo)

	userController controller.IUserController = controller.NewUserController(userService)
)

func main() {
	defer config.CloseDatabase(db)
	r := gin.Default()

	userRoutes := r.Group("/user").Use(middleware.Auth())
	{
		userRoutes.GET("/ping", middleware.Pong)
		userRoutes.GET("/profile/:id", userController.Profile)
	}

	r.Run(":8080")
}
