package main

import (
	"log"

	router "TaskManager/task-manager/Delivery/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	r = router.Setup(r)
	r.Run()
}
