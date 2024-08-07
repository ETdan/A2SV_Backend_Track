package data

import (
	"TaskManagerWithMongoDB/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(u models.User) (models.User, error) {
	var user models.User
	err := UsersCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: u.UserName}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// create accoutn
		user = u
		fmt.Println(u.ID, u.Password, u.Role, u.UserName)
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, errors.New("system error")
		}
		user.Password = string(password)
		_, err = UsersCollection.InsertOne(context.TODO(), user)
		if err != nil {
			return models.User{}, errors.New("system error")
		} else {
			return user, nil
		}
	} else {
		fmt.Println("username is Taken")
		return models.User{}, errors.New("username is Taken")
	}
}
func LoginUser(u models.User) (models.User, error) {
	var user models.User
	err := UsersCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: u.UserName}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, errors.New("invalid username or email")
	} else {
		return user, nil

	}
}
