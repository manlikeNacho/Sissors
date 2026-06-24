package sliceRepo

import (
	"context"
	"errors"
	"time"

	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/utils"
	"github.com/manlikeNacho/Sissors/src/utils/rerrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersDBRepository interface {
	GetUserByID(userId string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CheckUserExistsByEmail(email string) bool
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	SaveUser(user *models.SignupReq) (*models.User, error)
	InitailizeUserDb(db *mongo.Client, dbName, usersCollection string)
}

type userRepository struct {
	client          *mongo.Client
	dbName          string
	usersCollection string
}

var UserRepo UsersDBRepository = &userRepository{}

func (u *userRepository) InitailizeUserDb(db *mongo.Client, dbName, usersCollection string) {
	u.client = db
	u.dbName = dbName
	u.usersCollection = usersCollection
}

func (u *userRepository) collection() *mongo.Collection {
	return u.client.Database(u.dbName).Collection(u.usersCollection)
}

func (u *userRepository) SaveUser(user *models.SignupReq) (*models.User, error) {
	//Hash password
	password, err := utils.HashPasswords(user)
	if err != nil {
		return nil, err
	}
	//map user
	newUser := &models.User{
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Password:   password,
		Phone:      user.Phone,
		User_type:  user.User_type,
		Created_at: time.Now().Unix(),
		Updated_at: time.Now().Unix(),
	}
	//save to db
	_, err = u.collection().InsertOne(context.Background(), newUser)
	//return user or error
	return newUser, nil
}

func (u *userRepository) GetUserByID(userId string) (*models.User, error) {
	user := &models.User{}
	query := bson.M{
		"id": userId,
	}
	if err := u.collection().FindOne(context.Background(), query).Decode(user); err != nil {
		return nil, rerrors.Format(rerrors.NotFoundErr, errors.New("invalid user ID"))
	}

	return user, nil
}

func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := bson.M{
		"email": email,
	}
	if err := u.collection().FindOne(context.Background(), query).Decode(user); err != nil {
		return nil, rerrors.Format(rerrors.NotFoundErr, errors.New("invalid user ID"))
	}

	return user, nil
}

func (u *userRepository) CheckUserExistsByEmail(email string) bool {
	query := bson.M{
		"email": email,
	}

	collection := u.collection()
	count, err := collection.CountDocuments(context.Background(), query)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func (u *userRepository) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var users []models.User
	query := bson.M{}
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	collection := u.collection()
	cursor, err := collection.Find(context.Background(), query, opts)
	if err != nil {
		return nil, rerrors.Format(rerrors.InternalErr, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, rerrors.Format(rerrors.InternalErr, err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, rerrors.Format(rerrors.InternalErr, err)
	}
	return users, nil
}

func (u *userRepository) UpdateUser(user *models.User) error {
	filter := bson.M{"id": user.ID}

	updateQuery := bson.M{
		"$set": bson.M{
			"phone":      user.Phone,
			"first_name": user.First_name,
			"last_name":  user.Last_name,
		},
	}

	collection := u.collection()
	result, err := collection.UpdateOne(context.Background(), filter, updateQuery)
	if err != nil {
		return rerrors.Format(rerrors.InternalErr, err)
	}

	if result.MatchedCount < 1 {
		return rerrors.Format(rerrors.InternalErr, err)
	}

	return nil
}
