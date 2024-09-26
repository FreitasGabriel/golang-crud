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
	"github.com/FreitasGabriel/golang-crud/src/model/repository/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateUser(t *testing.T) {
	t.Run("user_already_registered_with_this_email", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		email := fmt.Sprintf("%d@test.com", rand.Int())

		_, err := Database.Collection("test_users").InsertOne(context.Background(), bson.M{"name": t.Name(), "email": email})
		if err != nil {
			t.Fatal(err)
			return
		}

		userRequest := request.UserRequest{
			Email:    email,
			Password: "test@#123",
			Name:     "Test User",
			Age:      32,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("user_is_not_registered_in_database", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		email := fmt.Sprintf("%d@test.com", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: "test@#123",
			Name:     "Test User",
			Age:      32,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		filter := bson.D{{Key: "email", Value: email}}
		userEntity := entity.UserEntity{}

		_ = Database.Collection("test_users").FindOne(context.Background(), filter).Decode(&userEntity)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, userEntity.Email, userRequest.Email)
		assert.EqualValues(t, userEntity.Age, userRequest.Age)
		assert.EqualValues(t, userEntity.Name, userRequest.Name)

	})

}
