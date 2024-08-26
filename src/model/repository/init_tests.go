package repository

import (
	"os"
	"testing"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var (
	database_name   = "user_database_test"
	collection_name = "user_collection_name"
)

func TestUserRepository_InitTests(t *testing.T) *mtest.T {

	err := os.Setenv("MONGODB_USER_DB", collection_name)
	if err != nil {
		t.FailNow()
		return nil
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mtestDb.Cleanup(func() {
		logger.Info("closing database connection")
	})

	return mtestDb
}
