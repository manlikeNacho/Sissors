package models

import (
	"errors"
	"html"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `json:"first_name" validate:"required, min=2, max=500" bson:"first_name"`
	Last_name  string             `json:"last_name" validate:"required, min=2, max=500" bson:"last_name"`
	Email      string             `json:"email" bson:"email" bson:"email"`
	Phone      int                `json:"phone" bson:"phone"`
	Password   string             `json:"password" bson:"password"`
	User_type  string             `json:"user_type" validate:"required , eq=ADMIN|eq=USER" bson:"user_type"`
	Created_at int64              `json:"created_at" bson:"created_at"`
	Updated_at int64              `json:"updated_at" bson:"updated_at"`
}

type SignupReq struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Phone      int    `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	User_type  string `json:"user_type"`
}

func (u *SignupReq) ValidateUser() error {
	u.First_name = strings.ToLower(html.EscapeString(strings.TrimSpace(u.First_name)))
	u.Last_name = strings.ToLower(html.EscapeString(strings.TrimSpace(u.Last_name)))
	u.Email = strings.ToLower(html.EscapeString(strings.TrimSpace(u.Email)))
	u.Password = strings.ToLower(html.EscapeString(strings.TrimSpace(u.Password)))

	if u.Email == "" {
		return errors.New("email must have a value")
	}

	return nil
}
