package sliceRepo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx, cancel   = context.WithTimeout(context.TODO(), 10*time.Second)
	CacheDuration = 6 * time.Hour
)

type Db struct {
	Db *mongo.Client
}

func New() *Db {
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	username := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://"+username+":"+password+"@cluster0.bbfxqpc.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection success")
	return &Db{
		Db: client,
	}
}

func (d Db) Close() error {
	return d.Db.Disconnect(context.Background())
}
