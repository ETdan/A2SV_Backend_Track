package router

import (
	db "TaskManager/task-manager/DB"
	controller "TaskManager/task-manager/Delivery/contollers"
	infrastucture "TaskManager/task-manager/Infrastucture"
	repositories "TaskManager/task-manager/Repositories"
	usecases "TaskManager/task-manager/Usecases"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// var router *gin.Engine
func Setup(r *gin.Engine) *gin.Engine {
	database := db.Database{
		Url: os.Getenv("DATABASE_URL"),
	}
	if err := database.Connect(os.Getenv("DATABASE"), os.Getenv("USER_COLLECTION"), os.Getenv("TASK_COLLECTION")); err != nil {
		fmt.Println(err)
	}

	middleare := infrastucture.Middleware{}

	userRepo := repositories.NewUserRepository(*database.Database, *database.UserCollection)
	taskRepo := repositories.NewTaskRepository(database, *database.Database, *database.TaskCollection)

	userUsecase := usecases.NewUserUsecase(userRepo)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	c := controller.Controller{
		TaskUsecase: taskUsecase,
		UserUsecase: userUsecase,
	}

	r.GET("/tasks", middleare.Auth_middleware(), c.GetAllTasks)
	r.GET("/tasks/:id", middleare.Auth_middleware(), c.GetTaskByID)
	r.PUT("/tasks/:id", middleare.Auth_middleware(), c.UpdateTaskByID)
	r.DELETE("/tasks/:id", middleare.Auth_middleware(), c.DeleteTaskByID)
	r.POST("/tasks", middleare.Auth_middleware(), c.AddTask)

	r.POST("/registeradmin", middleare.Auth_middleware(), c.RegisterAdmin)
	r.POST("/register", c.RegisterUser)
	r.POST("/login", c.LoginUser)

	return r
}
