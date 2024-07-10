package controller

import (
	"net/http"
	"net/mail"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControlerInterface) FindUserById(c *gin.Context) {

	logger.Info("Init findUserByID controller", zap.String("journey", "findUserByID"))

	userId, _ := c.Params.Get("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadResquestError(
			"UserID is not a valid id",
		)
		logger.Error("Error to trying validate userId", err, zap.String("journey", "findUserByID"))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail services", err, zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("jounrey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControlerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	userEmail, _ := c.Params.Get("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadResquestError(
			"UserEmail is not a valid email",
		)
		logger.Error("Error to trying to validate userEmail", err, zap.String("journey", "FindUserByEmail"))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail services", err, zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("jounrey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
