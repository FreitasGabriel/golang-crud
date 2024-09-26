package service

import (
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/configuration/tests/mock"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_UpdateUserServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 21)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateUser(id, userDomain)

		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_user_and_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 21)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(
			rest_err.NewInternalServerError("error trying to update user"),
		)

		err := service.UpdateUser(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}