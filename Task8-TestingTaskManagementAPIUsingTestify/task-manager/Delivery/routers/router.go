package router

import (
	db "TaskManager/task-manager/DB"
	controller "TaskManager/task-manager/Delivery/contollers"
	Infrastructure "TaskManager/task-manager/Infrastructure"
	repositories "TaskManager/task-manager/Repositories"
	usecases "TaskManager/task-manager/Usecases"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// var router *gin.Engine
func Setup(r *gin.Engine) *gin.Engine {
	database := db.Database{Url: os.Getenv("DATABASE_URL")}
	fmt.Println(os.Getenv("DATABASE_URL"), os.Getenv("USER_COLLECTION"), os.Getenv("TASK_COLLECTION"))

	if err := database.Connect(os.Getenv("DATABASE"), os.Getenv("USER_COLLECTION"), os.Getenv("TASK_COLLECTION")); err != nil {
		log.Fatal("could not connect with DB", err)
	}

	service := Infrastructure.NewService()
	middleare := Infrastructure.NewMiddleware(service)
	password := Infrastructure.NewPassword()

	userRepo := repositories.NewUserRepository(*database.Database, *database.UserCollection, *password)
	taskRepo := repositories.NewTaskRepository(database, *database.Database, *database.TaskCollection)

	userUsecase := usecases.NewUserUsecase(userRepo, *password)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	c := controller.NewController(userUsecase, taskUsecase)

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
