package service

import (
	"os"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/configuration/tests/mock"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_LoginUserServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_calling_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 21)
		userDomain.SetID(id)

		//Isso é necessário pois fazendo um <userDomainMock := userDomain> ainda continuaremos tomando erro
		//por causa de deep copy. Onde qualquer alteração em userDomainMock também alterará em userDomain
		userDomainMock := model.NewUserDomain(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetName(),
			userDomain.GetAge(),
		)

		userDomainMock.EncryptPassword()

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(), userDomainMock.GetPassword()).Return(
			nil, rest_err.NewInternalServerError("error trying to find user by email and password"))

		user, token, err := service.LoginUserService(userDomain)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to find user by email and password")
	})

	t.Run("when_calling_create_token_returns_error", func(t *testing.T) {

		userDomainMock := mock.NewMockUserDomainInterface(ctrl)

		userDomainMock.EXPECT().GetEmail().Return("teste@test.com")
		userDomainMock.EXPECT().GetPassword().Return("teste")
		userDomainMock.EXPECT().EncryptPassword().Return()

		userDomainMock.EXPECT().GenerateToken().Return("", rest_err.NewInternalServerError("error trying to create token"))

		repository.EXPECT().FindUserByEmailAndPassword(
			"teste@test.com", "teste").Return(
			userDomainMock, nil)

		user, token, err := service.LoginUserService(userDomainMock)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create token")

	})

	t.Run("when_user_and_password_is_valid_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		os.Setenv("JWT_SECRET_KEY", "test")
		defer os.Clearenv()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 21)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(), gomock.Any()).Return(
			userDomain, nil)

		user, _, err := service.LoginUserService(userDomain)

		assert.Nil(t, err)
		assert.Equal(t, user.GetID(), id)
		assert.Equal(t, user.GetEmail(), userDomain.GetEmail())
		assert.Equal(t, user.GetPassword(), userDomain.GetPassword())
		assert.Equal(t, user.GetName(), userDomain.GetName())
		assert.Equal(t, user.GetAge(), userDomain.GetAge())
	})
}
