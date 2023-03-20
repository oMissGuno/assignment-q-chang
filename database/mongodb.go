package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenDB(ctx context.Context, uri string, dbstr string) (*mongo.Database, error) {
	fmt.Println("db uri:", uri)
	fmt.Println("db dbstr:", dbstr)

	credential := options.Credential{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	opt := options.Client().ApplyURI(uri).SetAuth(credential)

	//connect mongodb
	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbstr)

	return db, nil
}
