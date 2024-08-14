package db

import (
	domain "TaskManager/task-manager/Domain"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Url            string
	Database       *mongo.Database
	UserCollection *mongo.Collection
	TaskCollection *mongo.Collection
	Client         *mongo.Client
}

func NewDatabase() {}

type DatabaseInterface interface {
	Connect(database_name string, user_collectionName string, task_collectionName string) error
	GetAllTasks(filter interface{}) (*mongo.Cursor, error)
	GetTaskByID(filter interface{}) *mongo.SingleResult
	UpdateTaskByID(context context.Context, filter interface{}, task domain.Task) (*mongo.UpdateResult, error)
	PostTask(ctx context.Context, task domain.Task) (string, error)
	DeleteTask(ctx context.Context, filter interface{}) (int64, error)

	FilterByUserID(user_id string) primitive.D
	FilterByCreaterID(user_id string) primitive.D
	FilterEmpty() primitive.D
}

func (d *Database) Connect(database_name string, user_collectionName string, task_collectionName string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(d.Url).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	var err error
	d.Client, err = mongo.Connect(context.TODO(), clientOptions)
	d.Database = d.Client.Database(database_name)
	d.TaskCollection = d.Database.Collection(task_collectionName)
	d.UserCollection = d.Database.Collection(user_collectionName)
	// .Collection("tasks")

	if err != nil {
		return err
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	// Check the connection
	if err = d.Client.Ping(context.TODO(), nil); err != nil {
		return err
	} else {
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	}
	return nil
}

func (d *Database) GetAllTasks(context context.Context, filter interface{}) (*mongo.Cursor, error) {
	findOptions := options.Find()
	cursor, err := d.TaskCollection.Find(context, filter.(primitive.D), findOptions)

	if err == nil {
		return cursor, nil
	} else {
		return nil, err
	}
}
func (d *Database) GetTaskByID(ctx context.Context, filter interface{}) (domain.Task, error) {
	findOptions := options.FindOne()
	var task domain.Task
	if err := d.TaskCollection.FindOne(ctx, filter, findOptions).Decode(&task); err == nil {
		return task, nil
	} else {
		return domain.Task{}, err
	}
}
func (d *Database) UpdateTaskByID(context context.Context, filter interface{}, task domain.Task) (*mongo.UpdateResult, error) {
	var update primitive.D
	if task.Name != "" {
		update = append(update, primitive.E{Key: "name", Value: task.Name})
	}
	if task.Detail != "" {
		update = append(update, primitive.E{Key: "detail", Value: task.Detail})
	}
	if task.Start != "" {
		update = append(update, primitive.E{Key: "start", Value: task.Start})
	}
	if task.Duration != "" {
		update = append(update, primitive.E{Key: "duration", Value: task.Duration})
	}
	updateDocument := bson.M{"$set": update}
	result, err := d.TaskCollection.UpdateOne(context, filter, updateDocument)
	return result, err
}
func (d *Database) PostTask(ctx context.Context, task domain.Task) (string, error) {
	result, err := d.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return "", err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		insertedID := id.Hex()
		return insertedID, nil
	} else {
		return "", errors.New("failed to convert InsertedID to primitive.ObjectID")
	}
}
func (d *Database) DeleteTask(ctx context.Context, filter interface{}) (int64, error) {
	deleteResult, err := d.TaskCollection.DeleteOne(ctx, filter)
	if err == nil {
		return deleteResult.DeletedCount, nil
	} else {
		return 0, err
	}
}

func (d *Database) FilterByCreaterID(user_id primitive.ObjectID) bson.D {

	filter := bson.D{{Key: "creater_id", Value: user_id}}
	return filter
}
func (d *Database) FilterByUserID(user_id string) primitive.D {
	filter := bson.D{{Key: "_id", Value: user_id}}
	return filter
}
func (d *Database) FilterByTaskID(task_id primitive.ObjectID) bson.D {
	filter := bson.D{{Key: "_id", Value: task_id}}
	return filter
}
func (d *Database) FilterEmpty() primitive.D {
	filter := bson.D{{}}
	return filter
}

// func (d *Database) Find(filter interface{}) (*mongo.Cursor, error) {
// 	findOptions := options.Find()
// 	cursor, err := d.TaskCollection.Find(context.TODO(), filter, findOptions)

// 	if err == nil {
// 		return cursor, nil
// 	}
// 	return nil, err
// }
