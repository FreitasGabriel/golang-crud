package main

import (
	"context"
	"log"

	"github.com/FreitasGabriel/golang-crud/src/configuration/database/mongodb"
	configuration "github.com/FreitasGabriel/golang-crud/src/configuration/dependencies"
	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Meu primeiro CRUD em Go | Gabriel Freitas
// @version 1.0.0
// @description API for crud operations on users
// @host localhost:8080
// @Basepath /
// @schemes http
// @license MIT
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

	userController := configuration.InitDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
