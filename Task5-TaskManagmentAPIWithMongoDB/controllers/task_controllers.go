package controllers

import (
	"fmt"
	"net/http"

	"TaskManagerWithMongoDB/data"
	"TaskManagerWithMongoDB/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// data.Data = {
// 	"1": models.Task{ /* Initialize with values */ },
// 	"2": models.Task{ /* Initialize with values */ },
// 	"3": models.Task{ /* Initialize with values */ },
// 	"4": models.Task{ /* Initialize with values */ },
// }

// Done
func GetAllTask(c *gin.Context) {
	fmt.Println(data.Data)
	Data, err := data.GetAllTask()
	if err == nil {
		c.JSON(http.StatusOK, Data)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
}

func GetTask(c *gin.Context) {

	id := c.Param("id")
	task, err := data.GetTask(id)
	if err == nil {
		// fmt.Println("data found")
		c.IndentedJSON(http.StatusFound, gin.H{"message": "successfully Found", "data": task})
	} else {
		// fmt.Println("/////////////////error///////////")
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

// Done
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val, err := data.UpdateTask(task, id)
	if err == nil {
		fmt.Println("data found")
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "successfully ADD", "data": val})
	} else {
		fmt.Println("/////////////////error///////////")
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

}

// Done
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := data.DeleteTask(id)
	if err == nil {
		fmt.Println("data found")
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "successfully Deleted", "data": task})
	} else {
		fmt.Println("/////////////////error///////////")
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func PostTask(c *gin.Context) {

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = primitive.NewObjectID()

	val, err := data.AddTask(task)
	if err == nil {
		fmt.Println("data found")
		c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully ADD", "data": val})
	} else {
		fmt.Println("/////////////////error///////////")
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Task received", "data": task})
}
