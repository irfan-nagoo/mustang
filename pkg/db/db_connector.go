package db

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetDatabaseClient() *mongo.Client {

	uri := os.Getenv("DATABASE_URL")
	
	loggerOptions := options.
			Logger().
			SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug) 
	
	clientOpts := options.Client().
								ApplyURI(uri).
								SetLoggerOptions(loggerOptions)
	
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		panic(err)
	}
	return client
}