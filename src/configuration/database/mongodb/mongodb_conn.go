package mongodb

import (
	"context"
	"os"

	"github.com/FreitasGabriel/golang-crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL = "MONGODB_URL"
	DB_NAME     = "MONGODB_USER_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {

	mongodb_uri := os.Getenv(MONGODB_URL)
	db_name := os.Getenv(DB_NAME)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("connection successful with database")

	return client.Database(db_name), nil
}
