package controller

import (
	domain "TaskManager/task-manager/Domain"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserUsecase domain.UserUsecase
	TaskUsecase domain.TaskUsecase
}

func NewController(uu domain.UserUsecase, tu domain.TaskUsecase) *Controller {
	return &Controller{
		UserUsecase: uu,
		TaskUsecase: tu,
	}
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
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
				c.Abort()
			} else {
				if len(tasks) == 0 {
					c.JSON(http.StatusOK, gin.H{"tasks:": "this user got no tasks"})
				} else {
					c.JSON(http.StatusOK, gin.H{"tasks:": tasks})
				}
			}
		}
	}
}
func (controller *Controller) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
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
			// fmt.Println(jwtClaims["role"].(string), jwtClaims["user_id"].(string), id)
			tasks, err := controller.TaskUsecase.GetTaskByID(jwtClaims["role"].(string), jwtClaims["user_id"].(string), id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
				c.Abort()
			} else {
				c.JSON(http.StatusOK, gin.H{"tasks:": tasks})
			}
		}
	}
}
func (controller *Controller) AddTask(c *gin.Context) {
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
			var task domain.Task
			if err := c.ShouldBindJSON(&task); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			} else {
				insertedId, err := controller.TaskUsecase.AddTask(jwtClaims["role"].(string), jwtClaims["user_id"].(string), task)
				if err == nil {
					c.JSON(http.StatusAccepted, gin.H{"Task Insered at ID": insertedId})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				}
			}
		}
	}
}
func (controller *Controller) UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")
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
			var task domain.Task
			if err := c.ShouldBindJSON(&task); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			} else {
				err := controller.TaskUsecase.UpdateTaskByID(jwtClaims["role"].(string), jwtClaims["user_id"].(string), task, id)
				if err == nil {
					c.JSON(http.StatusAccepted, gin.H{"Message": "successfully updated"})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				}
			}
		}
	}

}
func (controller *Controller) DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")
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
			err := controller.TaskUsecase.DeleteTaskByID(jwtClaims["role"].(string), jwtClaims["user_id"].(string), id)
			if err == nil {
				c.JSON(http.StatusAccepted, gin.H{"Message": "successfully Deleted"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			}
		}
	}
}
