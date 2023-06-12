package sliceRepo

import (
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/repository"
	"github.com/redis/go-redis/v9"
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

	return &Db{
		Db: rdb,
	}
}

func (d Db) SaveUrl(u models.Url) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d Db) GetUrl(u models.Url) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d Db) DeleteUrl(u models.Url) error {
	//TODO implement me
	panic("implement me")
}
