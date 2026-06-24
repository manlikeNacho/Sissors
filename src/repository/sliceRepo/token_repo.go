package sliceRepo

import (
	"context"

	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/utils/rerrors"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository interface {
	SaveRefreshToken(userId, tokenId, refreshString string) error
	InitializeTokenDB(db *mongo.Client, dbName, token_collection string)
	DeleteUserRefreshToken(userid string) error
}

type tokenDBRepository struct {
	client          *mongo.Client
	dbName          string
	tokenCollection string
}

// tokenDBRepository implements TokenRepository interface
var TokenRepo TokenRepository = &tokenDBRepository{}

func (d *tokenDBRepository) InitializeTokenDB(db *mongo.Client, dbName, token_collection string) {
	d.client = db
	d.dbName = dbName
	d.tokenCollection = token_collection
}

// initTokenCollection setup token collection in db.
func (d *tokenDBRepository) initTokenCollection() *mongo.Collection {
	return d.client.Database(d.dbName).Collection(d.tokenCollection)
}

// SaveRefreshToken save new refresh token in collection
func (d *tokenDBRepository) SaveRefreshToken(userId, tokenId, refreshString string) error {
	refresh_token := &models.RefreshToken{
		UserID:        userId,
		TokenID:       tokenId,
		RefreshString: refreshString,
	}

	_, err := d.initTokenCollection().InsertOne(context.Background(), refresh_token)
	if err != nil {
		return rerrors.Format(rerrors.InternalErr, err)
	}
	return nil
}

// DeleteUserRefreshToken delete refresh tokens
func (d *tokenDBRepository) DeleteUserRefreshToken(userid string) error {
	_, err := d.initTokenCollection().DeleteOne(context.Background(), userid)
	if err != nil {
		return rerrors.Format(rerrors.InternalErr, err)
	}

	return nil
}
