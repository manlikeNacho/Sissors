package repository

import "github.com/manlikeNacho/Sissors/src/models"

type Repository interface {
	SaveUrl(u *models.Url) error
	GetUrl(s string) (string, error)
	DeleteUrl(u models.Url) error
}
