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
	databaseName = config.Env.DATABASE_NAME

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(config.Env.MONGO_URI))
	if err != nil {
		return err
	}

	if err = mongoClient.Ping(context.Background(), nil); err != nil {
		return err
	}

	return nil
}

func CloseConnection() {
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database(databaseName).Collection(name)
}
