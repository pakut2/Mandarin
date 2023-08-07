package database

import (
	"context"

	"github.com/pakut2/mandarin/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var databaseName string

func InitConnection() error {
	mongoUri := config.Env.MONGO_URI
	databaseName = config.Env.DATABASE_NAME

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		return err
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	return nil
}

func CloseConnection() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database(databaseName).Collection(name)
}
