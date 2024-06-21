package controller

import (
	"github.com/FreitasGabriel/golang-crud/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControlerInterface{
		service: serviceInterface,
	}
}

type userControlerInterface struct {
	service service.UserDomainService
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
}
