package service

import (
	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init createUser model", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadResquestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("error trying to call CreateUser repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
	)

	return userDomainRepository, nil
}
