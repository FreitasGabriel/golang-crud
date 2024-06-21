package view

import (
	"github.com/FreitasGabriel/golang-crud/src/controller/model/response"
	"github.com/FreitasGabriel/golang-crud/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
