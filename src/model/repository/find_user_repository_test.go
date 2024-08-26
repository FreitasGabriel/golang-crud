package repository

import (
	"fmt"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"github.com/FreitasGabriel/golang-crud/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {

	mtestDb := TestUserRepository_InitTests(t)
	mtestDb.Cleanup(func() {
		logger.Info("closing database connection")
	})

	mtestDb.Run("when_sending_a_valid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@tes.com",
			Password: "test",
			Name:     "Teste",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBSON(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
	})

	mtestDb.Run("return_error_when_mongodb_returns_error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("tt@tt.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("return_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("teste")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {

	mtestDb := TestUserRepository_InitTests(t)
	mtestDb.Cleanup(func() {
		logger.Info("closing database connection")
	})

	mtestDb.Run("when_sending_a_valid_id_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@tes.com",
			Password: "test",
			Name:     "Teste",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBSON(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
	})

	mtestDb.Run("return_error_when_mongodb_returns_error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("tt@tt.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("return_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("teste")

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, fmt.Sprintf("User not found with this id: %s", "teste"))
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {

	mtestDb := TestUserRepository_InitTests(t)
	mtestDb.Cleanup(func() {
		logger.Info("closing database connection")
	})

	mtestDb.Run("when_sending_a_valid_email_and_password_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@tes.com",
			Password: "test",
			Name:     "Teste",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBSON(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
	})

	mtestDb.Run("return_error_when_mongodb_returns_error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("tt@tt.com", "123123123")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("return_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("teste@asdasd.com", "123")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBSON(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
