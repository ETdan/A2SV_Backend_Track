package data

import (
	"TaskManagerWithMongoDB/models"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DB *mongo.Database
var Collection *mongo.Collection

func StartMongoDB() {
	// Set client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://ETdan:<password>@cluster0.f79ysrp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	DB = client.Database("taskManager")
	Collection = DB.Collection("tasks")
	// .Collection("tasks")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	}
	// disconnect with the db
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	} else {
	// 		fmt.Println("Connection to MongoDB closed.")
	// 	}
	// }()
}

var Data = map[string]models.Task{}

func AddTask(task models.Task) (*mongo.InsertOneResult, error) {
	// fmt.Println(client.Database("taskManager").Collection("tasks").CountDocuments(context.TODO(), bson.D{{}}))

	result, err := Collection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
func GetAllTask() ([]models.Task, error) {
	findOptions := options.Find()

	var data []models.Task
	var task models.Task

	cursor, err := Collection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err == nil {
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&task)
			if err == nil {
				data = append(data, task)
			} else {
				return []models.Task{}, err
			}
		}
		return data, nil
	}
	return []models.Task{}, err
}
func GetTask(id string) (models.Task, error) {
	var task models.Task
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	err = Collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: ID}}).Decode(&task)

	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
func DeleteTask(id string) (models.Task, error) {
	var task models.Task
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	// fmt.Println(ID)
	err = Collection.FindOneAndDelete(context.TODO(), bson.D{{Key: "_id", Value: ID}}).Decode(&task)

	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
func UpdateTask(task models.Task, id string) (models.Task, error) {
	// var updatedTask models.Task
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	fmt.Println(ID)
	updateFields := bson.M{}
	if task.Name != "" {
		updateFields["name"] = task.Name
	}
	if task.Detail != "" {
		updateFields["detail"] = task.Detail
	}
	if task.Start != "" {
		updateFields["start"] = task.Start
	}
	if task.Duration != "" {
		updateFields["duration"] = task.Duration
	}
	// updateFields["_id"] = ID

	// Check if there are fields to update
	if len(updateFields) == 0 {
		return models.Task{}, err
	}

	// Define the filter and update
	filter := bson.M{"_id": ID}
	update := bson.M{"$set": updateFields}

	// Perform the update operation
	result, err := Collection.UpdateOne(context.TODO(), filter, update)
	fmt.Println(result)
	if err != nil {
		return models.Task{}, err
	}
	// Send the response
	if result.MatchedCount == 0 {
		return models.Task{}, errors.New("Document not found")
	} else {
		return task, nil
	}
}
