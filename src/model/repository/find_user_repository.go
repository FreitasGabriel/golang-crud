package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/FreitasGabriel/golang-crud/src/model/repository/entity"
	"github.com/FreitasGabriel/golang-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init findUserByEmail repository", zap.String("journey", "findUserByEmail"))
	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errMessage := fmt.Sprintf(
				"User not found with this email: %s", email,
			)
			logger.Error(errMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errMessage)
		}

		errMessage := "Error trying to find user by email"
		logger.Error(errMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewNotFoundError(errMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init findUserByID repository", zap.String("journey", "findUserByID"))
	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errMessage := fmt.Sprintf(
				"User not found with this id: %s", id,
			)
			logger.Error(errMessage, err, zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errMessage)
		}

		errMessage := "Error trying to find user by id"
		logger.Error(errMessage, err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewNotFoundError(errMessage)
	}

	logger.Info(
		"FindUserByid repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"))
	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errMessage := "User not found with this email and password"

			logger.Error(errMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
			return nil, rest_err.NewForbiddenError(errMessage)
		}

		errMessage := "Error trying to find user by email and password"
		logger.Error(errMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errMessage)
	}

	logger.Info(
		"FindUserByEmailAndPassword repository executed successfully",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}
