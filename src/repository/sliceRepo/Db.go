package sliceRepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/repository"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ctx           = context.Background()
	CacheDuration = 6 * time.Hour
)

type Db struct {
	Db *redis.Client
}

var _ repository.Repository = &Db{}

func New() *Db {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Initialized Redis database")

	return &Db{
		Db: rdb,
	}
}

func (d Db) SaveUrl(u models.Url) error {
	if err := d.Db.Set(ctx, u.Url, u.ShortUrl, CacheDuration).Err(); err != nil {
		return err
	}
	return nil
}

func (d Db) GetUrl(u models.Url) (string, error) {
	val, err := d.Db.Get(ctx, u.Url).Result()
	if err == redis.Nil {
		return "", errors.New("url not found")
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (d Db) DeleteUrl(u models.Url) error {
	//TODO implement me
	panic("implement me")
}
