package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// FindUserByID retrieves user ionformation based on the provided user ID
// @Summary Find user by id
// @Description Retrieve user details based on the user ID provided as a parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int tru "ID of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: invalid user ID"
// @Failure 404 {object} rest_err.RestErr "Error: user not found"
// @Router /user/id/{id} [get]
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
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("jounrey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

// FindUserByEmail retrieves user ionformation based on the provided user Email
// @Summary Find user by Email
// @Description Retrieve user details based on the user Email provided as a parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int tru "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: invalid user Email"
// @Failure 404 {object} rest_err.RestErr "Error: user not found"
// @Router /user/email/{id} [get]
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
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("jounrey", "FindUserByEmail"))
	fmt.Println(userDomain)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
