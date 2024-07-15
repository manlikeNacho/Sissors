package sliceRepo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/repository"
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
	// client *mongo.Client
}

var _ repository.Repository = &Db{}

func New() *Db {
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	username := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// colName := os.Getenv("COL_NAME")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://"+username+":"+password+"@cluster0.bbfxqpc.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection success")

	// db := client.Database(dbName).Collection(colName)
	return &Db{
		// Db:     db,
		// client: client,
		Db: client,
	}
}

func (d Db) Close() error {
	return d.Db.Disconnect(context.Background())
}

func (d Db) SaveUrl(u *models.Url) error {
	// if _, err := d.Db.InsertOne(context.Background(), u); err != nil {
	// 	return err
	// }
	return nil
}

func (d Db) GetUrl(s string) (string, error) {
	// filter := bson.D{{Key: "ShortUrl", Value: s}}
	var result models.Url

	// err := d.Db.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		return "", err
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }

	// fmt.Println("Found document:", result)
	return result.Url, nil
}

func (d Db) DeleteUrl(u models.Url) error {
	//TODO implement me
	panic("implement me")
}

// func (d Db) SaveUser(u models.User) error {
// 	if _, err := d.Db.InsertOne(ctx, u); err != nil {
// 		return err
// 	}

// 	return nil
// }
