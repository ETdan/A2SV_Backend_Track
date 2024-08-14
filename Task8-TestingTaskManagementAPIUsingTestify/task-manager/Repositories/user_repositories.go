package repositories

import (
	domain "TaskManager/task-manager/Domain"
	Infrastructure "TaskManager/task-manager/Infrastructure"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db              mongo.Database
	collection      mongo.Collection
	PasswordService Infrastructure.PasswordService
}

func NewUserRepository(db mongo.Database, collection mongo.Collection, password Infrastructure.PasswordService) *UserRepository {
	return &UserRepository{
		db:              db,
		collection:      collection,
		PasswordService: password,
	}
}

func (r *UserRepository) GetUserByUserName(username string) (domain.User, error) {
	if username == "" {
		return domain.User{}, errors.New("username should not be empty")
	}
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
	var emptyID primitive.ObjectID
	if id == emptyID {
		return domain.User{}, errors.New("id should not be empty")
	}
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
	if user.Password == "" || user.UserName == "" {
		return domain.User{}, errors.New("username and password should not be empty")
	}
	user.ID = primitive.NewObjectID()
	user.Role = "user"
	password, err := r.PasswordService.GeneratePasswordHash(user.Password)
	if err != nil {
		return domain.User{}, errors.New("could not generate bycrpt from password")
	}
	user.Password = string(password)

	// check existance of user with the same username
	_, e := r.GetUserByUserName(user.UserName)
	if e == nil {
		return domain.User{}, errors.New("username exist")
	}
	_, err = r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}
func (r *UserRepository) CreateAdmin(user domain.User) (domain.User, error) {
	if user.Password == "" || user.UserName == "" {
		return domain.User{}, errors.New("username and password should not be empty")
	}
	user.ID = primitive.NewObjectID()
	user.Role = "admin"
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, errors.New("could not generate bycrpt from password")
	}
	user.Password = string(password)

	// check existance of user with the same username
	_, e := r.GetUserByUserName(user.UserName)
	if e == nil {
		return domain.User{}, errors.New("username exist")
	}

	_, err = r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return domain.User{}, err
	} else {
		return user, nil
	}
}
