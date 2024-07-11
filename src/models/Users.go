package models

import (
	"errors"
	"html"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `json:"first_name" validate:"required, min=2, max=500"`
	Last_name  string             `json:"last_name" validate:"required, min=2, max=500"`
	Email      string             `json:"email" bson:"email"`
	Phone      int                `json:"phone" bson:"phone"`
	Password   string             `json:"password" bson:"password"`
	User_type  string             `json:"user_type" validate:"required , eq=ADMIN|eq=USER"`
}

func (u *User) ValidateUser() error {
	// u.FullName = strings.ToLower(html.EscapeString(strings.TrimSpace(u.FullName)))
	u.Email = strings.ToLower(html.EscapeString(strings.TrimSpace(u.Email)))
	u.Password = strings.ToLower(html.EscapeString(strings.TrimSpace(u.Password)))

	// if u.FullName == "" {
	// 	return errors.New("fullname must have a value")
	// }
	if u.Email == "" {
		return errors.New("email must have a value")
	}

	return nil
}
