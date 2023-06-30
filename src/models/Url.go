package models

type Url struct {
	ID       string `json:"ID" bson:"ID"`
	Url      string `json:"Url" bson:"Url"`
	ShortUrl string `json:"ShortUrl" bson:"ShortUrl"`
}

type UrlResponse struct {
	Url string
}
