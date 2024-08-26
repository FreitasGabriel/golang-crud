package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/configuration/tests/mocks"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 21)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

		returnedUserDomain, err := service.FindUserByIDService(id)

		assert.Nil(t, err)
		assert.Equal(t, returnedUserDomain.GetID(), id)
		assert.Equal(t, returnedUserDomain.GetEmail(), userDomain.GetEmail())
		assert.Equal(t, returnedUserDomain.GetPassword(), userDomain.GetPassword())
		assert.Equal(t, returnedUserDomain.GetName(), userDomain.GetName())
		assert.Equal(t, returnedUserDomain.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		returnedUserDomain, err := service.FindUserByIDService(id)

		assert.NotNil(t, err)
		assert.Nil(t, returnedUserDomain)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"
		userDomain := model.NewUserDomain(email, "test", "test", 21)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		returnedUserDomain, err := service.FindUserByEmailService(email)

		assert.Nil(t, err)
		assert.Equal(t, returnedUserDomain.GetID(), id)
		assert.Equal(t, returnedUserDomain.GetEmail(), userDomain.GetEmail())
		assert.Equal(t, returnedUserDomain.GetPassword(), userDomain.GetPassword())
		assert.Equal(t, returnedUserDomain.GetName(), userDomain.GetName())
		assert.Equal(t, returnedUserDomain.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@error.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		returnedUserDomain, err := service.FindUserByEmailService(email)

		assert.NotNil(t, err)
		assert.Nil(t, returnedUserDomain)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain(email, password, "test", 21)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		returnedUserDomain, err := service.findUserByEmailAndPasswordService(email, password)

		assert.Nil(t, err)
		assert.Equal(t, returnedUserDomain.GetID(), id)
		assert.Equal(t, returnedUserDomain.GetEmail(), userDomain.GetEmail())
		assert.Equal(t, returnedUserDomain.GetPassword(), userDomain.GetPassword())
		assert.Equal(t, returnedUserDomain.GetName(), userDomain.GetName())
		assert.Equal(t, returnedUserDomain.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@error.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
		returnedUserDomain, err := service.findUserByEmailAndPasswordService(email, password)

		assert.NotNil(t, err)
		assert.Nil(t, returnedUserDomain)
		assert.EqualValues(t, err.Message, "user not found")
	})
}
