package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserName string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}
