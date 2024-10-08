package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/rest_err"
	"github.com/FreitasGabriel/golang-crud/src/configuration/tests/mock"
	"github.com/FreitasGabriel/golang-crud/src/controller/model/request"
	"github.com/FreitasGabriel/golang-crud/src/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_LoginUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLogin{
			Email:    "ERROR@_EMAIL",
			Password: "error",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		user := request.UserLogin{
			Email:    "gabriielfs96@gmail.com",
			Password: "GFSp1q2w#E$R",
		}

		domain := model.NewUserLoginDomain(
			user.Email,
			user.Password,
		)

		b, _ := json.Marshal(user)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginUserService(domain).Return(nil, "", rest_err.NewInternalServerError("error test"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_and_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		token := primitive.NewObjectID().Hex()

		userRequest := request.UserLogin{
			Email:    "test@test.com",
			Password: "teste@#123",
		}

		userDomain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginUserService(userDomain).Return(userDomain, token, nil)

		fmt.Println(stringReader)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("Authorization")[0], token)
	})
}
