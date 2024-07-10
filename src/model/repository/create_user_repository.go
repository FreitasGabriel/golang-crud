package repository

import (
	"context"
	"os"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/FreitasGabriel/golang-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init createUser repository", zap.String("journey", "createUser"))
	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("error trying to create user", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("CreateUser repository executed successfully",
		zap.String("userID", value.ID.Hex()),
		zap.String("journey", "createUser"))

	return converter.ConvertEntityToDomain(*value), nil
}
