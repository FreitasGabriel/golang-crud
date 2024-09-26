package tests

import (
	"context"
	"encoding/json"
	"io"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateUser(t *testing.T) {
	t.Run("user_updated_with_success", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		_, err := Database.Collection("test_users").InsertOne(context.Background(), bson.M{"_id": id, "name": "OLA NAME", "age": 10, "email": "teste@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		userRequest := request.UserUpdateRequest{
			Name: "NEW NAME",
			Age:  30,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
		UserController.UpdateUser(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)

		filter := bson.D{{Key: "_id", Value: id}}
		userEntity := entity.UserEntity{}

		_ = Database.Collection("test_users").FindOne(context.Background(), filter).Decode(&userEntity)

		assert.EqualValues(t, userRequest.Name, userEntity.Name)
		assert.EqualValues(t, userRequest.Age, userEntity.Age)

	})

}
