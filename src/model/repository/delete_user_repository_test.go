package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	mtestDb := TestUserRepository_InitTests(t)

	mtestDb.Run("when_sending_a_valid_userId_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)

		err := repo.DeleteUser("test")

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_Database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)

		err := repo.DeleteUser("tes")

		assert.NotNil(t, err)
	})
}