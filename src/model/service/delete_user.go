package service

import (
	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("init deleteUser model", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("error trying to call deleteUser repository", err, zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("deleteUser controller executed successfully",
		zap.String("userId", id),
		zap.String("journey", "deleteUser"),
	)

	return nil
}
