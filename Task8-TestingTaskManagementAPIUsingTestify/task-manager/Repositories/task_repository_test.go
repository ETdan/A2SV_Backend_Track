package repositories

import (
	db "TaskManager/task-manager/DB"
	domain "TaskManager/task-manager/Domain"
	"TaskManager/task-manager/mocks"
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRpositorySuite struct {
	suite.Suite
	database      db.Database
	collection    *mocks.DatabaseInterface
	taskRpository TaskRpository
}

func (tr *TaskRpositorySuite) SetupSuite() {
	tr.database = db.Database{Url: os.Getenv("DATABASE_URL")}
	tr.collection = new(mocks.DatabaseInterface)
	tr.taskRpository = *NewTaskRepository(tr.database)
}
func (tr *TaskRpositorySuite) SetupSTest() {
	// tr.database = db.Database{Url: os.Getenv("DATABASE_URL")}
	// tr.db = mongo.Database{}
	// tr.c = mongo.Collection{}
	tr.collection = new(mocks.DatabaseInterface)
	tr.taskRpository = *NewTaskRepository(tr.database)
}
func (tr *TaskRpositorySuite) TestDeleteTaskByID() {
	// Create context with timeout
	cxt, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Set up ObjectID
	object_id, err := primitive.ObjectIDFromHex("507f191e810c19729de860ea")
	if err != nil {
		tr.T().Fatalf("Failed to convert hex to ObjectID: %v", err)
	}

	// Set up mock expectations
	tr.collection.On("FilterByTaskID", object_id).Return(primitive.D{{Key: "_id", Value: object_id}})
	tr.collection.On("DeleteTask", cxt, primitive.D{{Key: "_id", Value: object_id}}).Return(int64(1), nil)
	tr.collection.On("GetTaskByID", cxt, primitive.D{{Key: "_id", Value: object_id}}).Return(domain.Task{Creater_ID: object_id}, nil)

	// Call the method under test
	err = tr.taskRpository.DeleteTaskByID("user", "507f191e810c19729de860ea", "507f191e810c19729de860ea", time.Second)

	// Assert no error
	tr.Nil(err, "Should be nil")

	// Verify mock expectations
	tr.collection.AssertExpectations(tr.T())
}

func (tr *TaskRpositorySuite) TestUpdateTaskByID() {
	// role string, user_id string, time_duration time.Duration, task Task, id string
	// error
}
func (tr *TaskRpositorySuite) TestAddTask() {
	// user_id string, task Task, duration time.Duration
	// (string, error)
}
func (tr *TaskRpositorySuite) TestGetTaskByID() {
	// role string, user_id string, task_id string, duration time.Duration
	//  (Task, error)3
}
func (tr *TaskRpositorySuite) TestGetAllTasks() {
	// role string, user_id string, duration time.Duration
	// ([]Task, error)
}

func TestTaskRpositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRpositorySuite))
}
