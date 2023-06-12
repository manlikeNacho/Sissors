package repository

import "github.com/manlikeNacho/Sissors/src/models"

type Repository interface {
	SaveUrl(u models.Url) (string, error)
	DeleteUrl(u models.Url) error
}
