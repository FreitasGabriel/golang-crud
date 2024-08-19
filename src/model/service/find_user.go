package service

import (
	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init findUserByEmail service", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init findUserByID service", zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) findUserByEmailAndPasswordService(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init findUserByID service", zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
