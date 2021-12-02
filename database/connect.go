package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vasim-akram/mongo-todo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	CONN_URL := config.Config("MONGO_URL")
	DB := config.Config("DB_NAME")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONN_URL))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database")

	return client.Database(DB)
}
