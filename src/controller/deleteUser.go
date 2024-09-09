package controller

import (
	"net/http"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControlerInterface) DeleteUser(c *gin.Context) {

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadResquestError("invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
		return
	}
	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("error trying to call deleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	c.Status(http.StatusOK)

}
