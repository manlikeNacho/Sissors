package models

import (
	"errors"
	"html"
	"strings"
)

type User struct {
	ID       string
	FullName string `json:"full_name" bson:"full_name"`
	Email    string `json:"email" bson:"email"`
	Phone    int    `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
}

func (u *User) ValidateUser() error {
	u.FullName = strings.ToLower(html.EscapeString(strings.TrimSpace(u.FullName)))
	u.Email = strings.ToLower(html.EscapeString(strings.TrimSpace(u.Email)))
	u.Password = strings.ToLower(html.EscapeString(strings.TrimSpace(u.Password)))

	if u.FullName == "" {
		return errors.New("fullname must have a value")
	}
	if u.Email == "" {
		return errors.New("email must have a value")
	}

	return nil
}
