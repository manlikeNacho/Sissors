package models

import "time"

type Url struct {
	ID        string
	Url       string
	ShortUrl  string
	CreatedAt time.Time
}

type UrlResponse struct {
	Url string
}
