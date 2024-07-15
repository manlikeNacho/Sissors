package models

import "github.com/manlikeNacho/Sissors/src/utils/rerrors"

type AccessToken struct {
	AccessToken string `json:"access_token" bson:"access_token"`
}

type RefreshToken struct {
	UserID        string `json:"-" bson:"user_id"`
	TokenID       string `json:"-" bson:"token_id"`
	RefreshString string `json:"refresh_token" bson:"refresh_string"`
}

// TokenPair is used to return pairs of id and refresh tokens
type TokenPair struct {
	RefreshToken
	AccessToken
}

type TokensReq struct {
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

func (t *TokensReq) Validate() error {
	if t.RefreshToken == "" {
		return rerrors.Format(rerrors.UnproccessibleEntityErr, nil)
	}

	return nil
}
