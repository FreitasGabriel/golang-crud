package routes

import (
	"github.com/FreitasGabriel/golang-crud/src/controller"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, controller controller.UserControllerInterface) {

	r.GET("/user/id/:userId", model.VerifyTokenMiddleware, controller.FindUserById)
	r.GET("/user/email/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
	r.POST("/user/login", controller.LoginUser)
}
