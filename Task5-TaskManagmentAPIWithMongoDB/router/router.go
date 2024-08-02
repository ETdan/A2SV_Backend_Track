package router

import (
	"TaskManagerWithMongoDB/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/tasks", controllers.GetAllTask)
	r.GET("/tasks/:id", controllers.GetTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.POST("/tasks", controllers.PostTask)

	return r

}
