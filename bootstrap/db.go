package bootstrap

import mongo "github.com/go-bongo/bongo"

var Connection *mongo.Connection

func DatabaseConnection() (*mongo.Connection, error){

	config := &mongo.Config{
		ConnectionString: "localhost",
		Database:         "tracker",
	}

	connection, err := mongo.Connect(config)
	Connection = connection

	return connection, err
}

