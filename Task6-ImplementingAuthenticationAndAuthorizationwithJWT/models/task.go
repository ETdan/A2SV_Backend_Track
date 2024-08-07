package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID         primitive.ObjectID `bson:"_id"`
	Creater_ID primitive.ObjectID `bson:"creater_id"`
	Name       string             `bson:"name"`
	Detail     string             `bson:"detail"`
	Start      string             `bson:"start"`
	Duration   string             `bson:"duration"`
}
