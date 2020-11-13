package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	var pass string
	var uname string
	var uri string

	pass = os.Getenv("DATABASE_PASSWORD")
	uname = os.Getenv("DATABASE_UNAME")
	//fmt.Println(uname)
	//fmt.Println(pass)

	uri = "mongodb+srv://" + uname + ":" + pass + "@cluster0.rztnm.mongodb.net/test"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
