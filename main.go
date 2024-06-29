package main

import (
	"context"
	"log"

	"github.com/FreitasGabriel/golang-crud/src/configuration/database/mongodb"
	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/controller"
	"github.com/FreitasGabriel/golang-crud/src/controller/routes"
	"github.com/FreitasGabriel/golang-crud/src/model/repository"
	"github.com/FreitasGabriel/golang-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error lading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("error trying to connect to dabatase, error=%s \n", err.Error())
		return
	}

	//init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
