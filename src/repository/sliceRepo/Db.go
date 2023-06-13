package sliceRepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/repository"
	"github.com/redis/go-redis/v9"
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

func (d Db) SaveUrl(u *models.Url) error {
	if err := d.Db.Set(ctx, u.ShortUrl, u.Url, CacheDuration).Err(); err != nil {
		return err
	}
	return nil
}

func (d Db) GetUrl(s string) (string, error) {
	val, err := d.Db.Get(ctx, s).Result()
	if err != nil {
		return "", errors.New("url not found")
	}
	return val, nil
}

func (d Db) DeleteUrl(u models.Url) error {
	//TODO implement me
	panic("implement me")
}
