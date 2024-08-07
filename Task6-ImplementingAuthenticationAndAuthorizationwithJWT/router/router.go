package router

import (
	"TaskManagerWithMongoDB/controllers"
	"TaskManagerWithMongoDB/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/tasks", middleware.Middleware(), controllers.GetAllTask)
	r.GET("/tasks/:id", middleware.Middleware(), controllers.GetTask)
	r.PUT("/tasks/:id", middleware.Middleware(), controllers.UpdateTask)
	r.DELETE("/tasks/:id", middleware.Middleware(), controllers.DeleteTask)
	r.POST("/tasks", middleware.Middleware(), controllers.PostTask)

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	return r

}
