package controller

import (
	"net/http"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/validation"
	"github.com/FreitasGabriel/golang-crud/src/controller/model/request"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/FreitasGabriel/golang-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControlerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller", zap.String("journey", "loginUser"))

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err, zap.String("journey", "loginUser"))
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error("error trying to call loginUser service", err, zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("loginUser controller executed successfully",
		zap.String("userId", domain.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.Header("authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
