package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserName string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}
type Task struct {
	ID         primitive.ObjectID `bson:"_id"`
	Creater_ID primitive.ObjectID `bson:"creater_id"`
	Name       string             `bson:"name"`
	Detail     string             `bson:"detail"`
	Start      string             `bson:"start"`
	Duration   string             `bson:"duration"`
}

type UserUsecase interface {
	RegisterUser(User) (User, error)
	LoginUser(User) (string, error)
	RegisterAdmin(User) (User, error)
}

type UserRepository interface {
	CreateAdmin(user User) (User, error)
	CreateUser(user User) (User, error)
	GetUserByID(id primitive.ObjectID) (User, error)
	GetUserByUserName(username string) (User, error)
}

type TaskUsecase interface {
	DeleteTaskByID(role string, user_id string, task_id string) error
	UpdateTaskByID(role string, user_id string, task Task, id string) error
	AddTask(role string, user_id string, task Task) (string, error)
	GetTaskByID(role string, user_id string, task_id string) (Task, error)
	GetAllTasks(role string, user_id string) ([]Task, error)
}

type TaskRepository interface {
	DeleteTaskByID(role string, user_id string, task_id string, time_duration time.Duration) error
	UpdateTaskByID(role string, user_id string, time_duration time.Duration, task Task, id string) error
	AddTask(user_id string, task Task, duration time.Duration) (string, error)
	GetTaskByID(role string, user_id string, task_id string, duration time.Duration) (Task, error)
	GetAllTasks(role string, user_id string, duration time.Duration) ([]Task, error)
}
type PasswordInterface interface {
	ValidatePasswordHash(existingUserPassword string, claimPassword string) error
	GeneratePasswordHash(existingUserPassword string, claimPassword string) ([]byte, error)
}
