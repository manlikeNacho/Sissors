package models

import "time"

type Url struct {
	ID        string
	Url       string
	CreatedAt time.Time
}

type UrlResponse struct {
	Url string
}
