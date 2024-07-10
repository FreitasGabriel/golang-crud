package configuration

import (
	"github.com/FreitasGabriel/golang-crud/src/controller"
	"github.com/FreitasGabriel/golang-crud/src/model/repository"
	"github.com/FreitasGabriel/golang-crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
