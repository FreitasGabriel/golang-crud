package repository

import (
	"os"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var (
	database_name   = "user_database_test"
	collection_name = "user_collection_name"
)

func MockedDatabase(t *testing.T) *mtest.T {
	return mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
}

func TestUserRepository_CreateUser(t *testing.T) {

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := MockedDatabase(t)
	mtestDb.Cleanup(func() {
		logger.Info("closing database connection")
	})

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)

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
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.CreateUser(model.NewUserDomain(
			"test@test.com", "test", "test", 90,
		))

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
