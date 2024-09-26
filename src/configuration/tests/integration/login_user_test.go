package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestLoginUser(t *testing.T) {

	t.Run("login_user_with_sucess", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := "test@#123"

		_, err := Database.Collection("test_users").InsertOne(context.Background(), bson.M{"name": t.Name(), "email": email})
		if err != nil {
			t.Fatal(err)
			return
		}

		userRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      32,
		}

		bCreateUser, _ := json.Marshal(userRequest)
		stringReaderCreateUser := io.NopCloser(strings.NewReader(string(bCreateUser)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreateUser)
		UserController.CreateUser(ctxCreateUser)

		userLoginRequest := request.UserLogin{
			Email:    email,
			Password: password,
		}

		bLoginUser, _ := json.Marshal(userLoginRequest)
		stringReaderLoginUser := io.NopCloser(strings.NewReader(string(bLoginUser)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReaderLoginUser)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusOK, recorderLoginUser.Result().StatusCode)
		assert.NotEmpty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})

}
