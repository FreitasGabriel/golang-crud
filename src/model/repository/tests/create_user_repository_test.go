package repository

import (
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/FreitasGabriel/golang-crud/src/model/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {

	mtestDb := TestUserRepository_InitTests(t)

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.CreateUser(model.NewUserDomain(
			"test@test.com", "test", "test", 90,
		))

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
	})

	mtestDb.Run("return_error_from_Database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.CreateUser(model.NewUserDomain(
			"test@test.com", "test", "test", 90,
		))

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
