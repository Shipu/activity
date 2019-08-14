package bootstrap

import (
	mongo "github.com/go-bongo/bongo"
	"os"
)

var Connection *mongo.Connection

func DatabaseConnection() (*mongo.Connection, error) {

	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")

	config := &mongo.Config{
		ConnectionString: host,
		Database:         database,
	}

	connection, err := mongo.Connect(config)
	Connection = connection

	return connection, err
}
