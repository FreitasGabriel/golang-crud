package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/tests/connection"
	"github.com/FreitasGabriel/golang-crud/src/controller"
	"github.com/FreitasGabriel/golang-crud/src/model/repository"
	"github.com/FreitasGabriel/golang-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

func TestMain(m *testing.M) {

	os.Setenv("MONGODB_USER_DB", "users")
	os.Setenv("MONGODB_COLLECTION", "test_users")
	closeConnection := func() {}

	Database, closeConnection = connection.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service.NewUserDomainService(repo)
	UserController = controller.NewUserControllerInterface(userService)

	defer func() {
		os.Clearenv()
		closeConnection()
	}()

	os.Exit(m.Run())
}

func TestFindUserbyEmail(t *testing.T) {

	t.Run("user_not_found_with_this_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "teste@test.com",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_with_this_email_with_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		_, err := Database.Collection("test_users").InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "teste@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "teste@test.com",
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestFindUserbyId(t *testing.T) {

	t.Run("user_not_found_with_this_id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		UserController.FindUserById(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_with_this_id_with_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		_, err := Database.Collection("test_users").InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "teste@test.com"})
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

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserById(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.URL.RawQuery = u.Encode()
	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
