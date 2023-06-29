package sliceRepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/url"
	"time"

	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/repository"
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

	// Parse the Redis URL
	parsedURL, err := url.Parse("rediss://red-cie9ngunqql22egdj1ng:FQbBlx08hh4v42PJlrPnlVvu4txzmiXE@frankfurt-redis.render.com:6379")
	if err != nil {
		log.Fatal("Failed to parse Redis URL:", err)
	}
	pwd, _ := parsedURL.User.Password()
	log.Println(parsedURL.Host)
	log.Println(parsedURL.User.Username())
	log.Println(pwd)

	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     parsedURL.Host,
		Username: parsedURL.User.Username(),
		Password: pwd,
	})

	_ = redis.NewClient(&redis.Options{
		//Addr:     "redisDb:6379",
		Addr:     "rediss://red-cie9ngunqql22egdj1ng:FQbBlx08hh4v42PJlrPnlVvu4txzmiXE@frankfurt-redis.render.com:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Initialized Redis database")

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to Redis:", pong)

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
