package sliceRepo

import (
	"context"
	"fmt"
	"log"

	"github.com/manlikeNacho/Sissors/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlRepository interface {
	SaveUrl(u *models.Url) error
	GetUrl(urlString string) (string, error)
	DeleteUrl(urlID string) error
	InitializeUrlDB(db *mongo.Client, dbName, urlCollection string)
}

type urlDBRepository struct {
	client        *mongo.Client
	dbName        string
	urlCollection string
}

// urlDBRepository implements UrlRepository interface
var UrlRepo UrlRepository = &urlDBRepository{}

func (d *urlDBRepository) InitializeUrlDB(db *mongo.Client, dbName, urlCollection string) {
	d.client = db
	d.dbName = dbName
	d.urlCollection = urlCollection
}

func (d *urlDBRepository) initUrlCollection() *mongo.Collection {
	return d.client.Database(d.dbName).Collection(d.urlCollection)
}

func (d *urlDBRepository) SaveUrl(u *models.Url) error {
	if _, err := d.initUrlCollection().InsertOne(context.Background(), u); err != nil {
		return err
	}
	return nil
}

func (d *urlDBRepository) GetUrl(urlString string) (string, error) {
	filter := bson.D{{Key: "ShortUrl", Value: urlString}}
	var result models.Url

	err := d.initUrlCollection().FindOne(context.TODO(), filter).Decode(&result)
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

func (d *urlDBRepository) DeleteUrl(urlID string) error {
	filter := bson.D{{Key: "Url", Value: urlID}}
	_, err := d.initUrlCollection().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
