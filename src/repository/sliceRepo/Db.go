package sliceRepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx, cancel   = context.WithTimeout(context.TODO(), 10*time.Second)
	CacheDuration = 6 * time.Hour
	dbName        = "snipbit"
	colName       = "url-shortener"
	password      = "GHPDB2SlSSf99ElA"
	username      = "snipbit"
)

type Db struct {
	Db     *mongo.Collection
	client *mongo.Client
}

var _ repository.Repository = &Db{}

func New() *Db {
	defer cancel()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://"+username+":"+password+"@cluster0.bbfxqpc.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo connection success")

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName).Collection(colName)
	return &Db{
		Db:     db,
		client: client,
	}
}

func (d Db) Close() error {
	return d.client.Disconnect(ctx)
}

func (d Db) SaveUrl(u *models.Url) error {
	if _, err := d.Db.InsertOne(context.Background(), u); err != nil {
		return err
	}
	return nil
}

func (d Db) GetUrl(s string) (string, error) {
	filter := bson.D{{Key: "ShortUrl", Value: s}}
	var result models.Url

	err := d.Db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("Found document:", result)
	return result.Url, err
}

func (d Db) DeleteUrl(u models.Url) error {
	//TODO implement me
	panic("implement me")
}
