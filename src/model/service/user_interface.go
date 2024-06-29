package service

import (
	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/FreitasGabriel/golang-crud/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
