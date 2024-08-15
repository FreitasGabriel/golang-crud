package service

import (
	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init updateUser model", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		logger.Error("error trying to call updateUser repository", err, zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("updateUser controller executed successfully",
		zap.String("userId", id),
		zap.String("journey", "updateUser"),
	)

	return nil
}
