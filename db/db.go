package db

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const DatabaseName = "amber_db"
const MongoUrl = "mongodb://root:root@amber-db:27017/"

var Database *mongo.Database
var TerrariumCollection *mongo.Collection

func Connect(router *echo.Echo) error {
	clientOption := options.Client().ApplyURI(MongoUrl)
	client, err := mongo.Connect(clientOption)
	if err != nil {
		router.Logger.Error(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	router.Logger.Print("Connected to MongoDB")
	setupDB(client)
	return nil
}

func setupDB(client *mongo.Client) {
	Database = client.Database(DatabaseName)
	TerrariumCollection = Database.Collection("terrarium")
}
