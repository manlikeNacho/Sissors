package utils

import (
	"github.com/manlikeNacho/Sissors/src/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswords(u *models.User) (string, error) {
	// Hash passwords using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(passcode, hashPasscode string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasscode), []byte(passcode))
	if err != nil {
		return err
	}

	return nil
}
