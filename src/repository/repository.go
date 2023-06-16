package repository

import "github.com/manlikeNacho/Sissors/src/models"

//go:generate mockgen -source repository.go -destination ./mocks/repository.go -package mocks Repository

type Repository interface {
	SaveUrl(u *models.Url) error
	GetUrl(s string) (string, error)
	DeleteUrl(u models.Url) error
}
