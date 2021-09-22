package driver

import (
	"context"
	"fmt"
	"log"
	"go_mongo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Db *mongo.Database
}

var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {
	// user, password string, host string,
	/*
	   Connect to my cluster
	*/

	connStr := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Global.DB_USERNAME, config.Global.DB_PASSWORD, config.Global.DB_HOST, config.Global.DB_PORT)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection Ok")
	Mongo.Db = client.Database(config.Global.DB_DATABASE)
	return Mongo
}
