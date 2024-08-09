package repositories

import (
	domain "TaskManager/task-manager/Domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db         mongo.Database
	collection mongo.Collection
}

func NewUserRepository(db mongo.Database, collection mongo.Collection) *UserRepository {
	return &UserRepository{
		db:         db,
		collection: collection,
	}
}

func (r *UserRepository) GetUserByUserName(username string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"username": username}
	result := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if result == nil {
		return user, nil
	} else {
		return domain.User{}, errors.New("could not find user")
	}
}
func (r *UserRepository) GetUserByID(id primitive.ObjectID) (domain.User, error) {
	var user domain.User
	filter := bson.M{"_id": id}
	result := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if result == nil {
		return user, nil
	} else {
		return domain.User{}, nil
	}
}
func (r *UserRepository) CreateUser(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	user.Role = "user"
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, errors.New("Could not generate bycrpt from password")
	}
	user.Password = string(password)
	_, err = r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}
func (r *UserRepository) CreateAdmin(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	user.Role = "admin"
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, errors.New("Could not generate bycrpt from password")
	}
	user.Password = string(password)
	_, err = r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return domain.User{}, err
	} else {
		return user, nil
	}
}
