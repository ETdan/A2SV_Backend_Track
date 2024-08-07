package main

import (
	"TaskManagerWithMongoDB/data"
	"TaskManagerWithMongoDB/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	data.StartMongoDB()
	r := gin.Default()
	r = router.Router(r)
	r.Run()

}
