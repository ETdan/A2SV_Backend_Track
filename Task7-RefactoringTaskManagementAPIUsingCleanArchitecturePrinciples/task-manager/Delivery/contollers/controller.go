package controller

import (
	domain "TaskManager/task-manager/Domain"
	usecases "TaskManager/task-manager/Usecases"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserUsecase usecases.UserUsecase
	TaskUsecase usecases.TaskUsecase
}
type ControllerInterface interface {
	RegisterUser(*gin.Context)
	RegisterAdmin(c *gin.Context)
	LoginUser(c *gin.Context)

	GetAllTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
	AddTask(c *gin.Context)
	UpdateTaskByID(c *gin.Context)
	DeleteTaskByID(c *gin.Context)
}

func (controller *Controller) RegisterUser(c *gin.Context) {

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		user, err = controller.UserUsecase.RegisterUser(user)
		if err == nil {
			c.JSON(http.StatusAccepted, gin.H{"message": user})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		}
	}
}
func (controller *Controller) RegisterAdmin(c *gin.Context) {

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		user, err = controller.UserUsecase.RegisterAdmin(user)
		if err == nil {
			c.JSON(http.StatusAccepted, gin.H{"message": user})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
	}
}
func (controller *Controller) LoginUser(c *gin.Context) {

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		fmt.Println(user.UserName, user.Password)
		token, err := controller.UserUsecase.LoginUser(user)
		if err == nil {
			c.JSON(http.StatusAccepted, gin.H{"message": token})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
	}
}

func (controller *Controller) GetAllTasks(c *gin.Context) {
	// var
	if claims, ok := c.Get("claims"); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find token claims"})
		c.Abort()
		return
	} else {
		if jwtClaims, ok := claims.(jwt.MapClaims); !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not converts claims"})
			c.Abort()
			return

		} else {
			tasks, err := controller.TaskUsecase.GetAllTasks(jwtClaims["role"].(string), jwtClaims["user_id"].(string))
			// fmt.Println(jwtClaims)
			fmt.Println("///////////////////////")
			fmt.Println(jwtClaims["user_id"])
			fmt.Println(jwtClaims["role"])
			fmt.Println(jwtClaims["user_id"])
			// fmt.Println(tasks, err)
			fmt.Println("///////////////////////")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
				c.Abort()
			} else {
				c.JSON(http.StatusOK, gin.H{"tasks:": tasks})
			}
		}
	}

}
func (controller *Controller) GetTaskByID(c *gin.Context) {
	if claims, ok := c.Get("claims"); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find token claims"})
		c.Abort()
		return
	} else {
		if jwtClaims, ok := claims.(jwt.MapClaims); !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not converts claims"})
			c.Abort()
			return

		} else {
			tasks, err := controller.TaskUsecase.GetTaskByID(jwtClaims["role"].(string), jwtClaims["user_id"].(string), jwtClaims["task_id"].(string))
			// // fmt.Println(jwtClaims)
			// fmt.Println("///////////////////////")
			// fmt.Println(jwtClaims["user_id"])
			// fmt.Println(jwtClaims["role"])
			// fmt.Println(jwtClaims["user_id"])
			// // fmt.Println(tasks, err)
			// fmt.Println("///////////////////////")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
				c.Abort()
			} else {
				c.JSON(http.StatusOK, gin.H{"tasks:": tasks})
			}
		}
	}
}
func (controller *Controller) AddTask(c *gin.Context) {}
func (controller *Controller) UpdateTaskByID(c *gin.Context) {

}
func (controller *Controller) DeleteTaskByID(c *gin.Context) {}
