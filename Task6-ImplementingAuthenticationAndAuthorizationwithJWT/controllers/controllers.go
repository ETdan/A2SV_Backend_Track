package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"TaskManagerWithMongoDB/data"
	"TaskManagerWithMongoDB/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var secrate_key = []byte(os.Getenv("SECRATE_KEY"))

// Donee
func GetAllTask(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get claims"})
		c.Abort()
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not process claims"})
		c.Abort()
		return
	}

	user_id, ok := jwtClaims["user_id"].(string)
	var user_object_id primitive.ObjectID
	Data := []models.Task{}
	var err error
	if jwtClaims["role"] == "admin" {
		Data, err = data.GetAllTask()
	} else {
		user_object_id, err = primitive.ObjectIDFromHex(user_id)
		if err == nil {
			if ok {
				Data, err = data.GetAllUserTask(user_object_id)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error could not get primitive of user id"})
			c.Abort()
			return
		}
	}
	if err == nil {
		if len(Data) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "There is no Task created by user"})
		} else {
			c.JSON(http.StatusOK, Data)
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
}

// Donee
func GetTask(c *gin.Context) {
	id := c.Param("id")
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get claims"})
		c.Abort()
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not process claims"})
		c.Abort()
		return
	}

	user_id, ok := jwtClaims["user_id"].(string)
	var user_object_id primitive.ObjectID
	var err error
	if ok {
		user_object_id, err = primitive.ObjectIDFromHex(user_id)
		if err != nil {
			fmt.Println("Error converting string to ObjectID:", err)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": jwtClaims["user_id"]})
		c.Abort()
		return
	}

	task, err := data.GetTask(id, user_object_id, jwtClaims["role"].(string))
	if err == nil {
		c.IndentedJSON(http.StatusFound, gin.H{"message": "successfully Found", "data": task})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

// Done
func UpdateTask(c *gin.Context) {
	var task models.Task
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get claims"})
		c.Abort()
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not process claims"})
		c.Abort()
		return
	}

	user_id, ok := jwtClaims["user_id"].(string)
	var user_object_id primitive.ObjectID
	var err error
	if ok {
		user_object_id, err = primitive.ObjectIDFromHex(user_id)
		if err != nil {
			fmt.Println("Error converting string to ObjectID:", err)
			return
		}
	}
	id := c.Param("id")
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val, err := data.UpdateTask(task, id, user_object_id, jwtClaims["role"].(string))
	if err == nil {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "successfully ADD", "data": val})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

// Donee
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get claims"})
		c.Abort()
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not process claims"})
		c.Abort()
		return
	}

	user_id, ok := jwtClaims["user_id"].(string)
	var user_object_id primitive.ObjectID
	var err error
	if ok {
		user_object_id, err = primitive.ObjectIDFromHex(user_id)
		if err == nil {
			task, err := data.DeleteTask(id, user_object_id, jwtClaims["role"].(string))
			if err == nil {
				c.IndentedJSON(http.StatusAccepted, gin.H{"message": "successfully Deleted", "data": task})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			}
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while getting user id form claim"})
		}

	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while converting user id ro objectId"})
	}
}

// Donee
func PostTask(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get claims"})
		c.Abort()
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not process claims"})
		c.Abort()
		return
	}

	user_id, ok := jwtClaims["user_id"].(string)
	var user_object_id primitive.ObjectID
	var err error
	if ok {
		user_object_id, err = primitive.ObjectIDFromHex(user_id)
		if err != nil {
			fmt.Println("Error converting string to ObjectID:", err)
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": jwtClaims["user_id"]})
		c.Abort()
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = primitive.NewObjectID()
	task.Creater_ID = user_object_id
	val, err := data.AddTask(task)
	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully ADD", "data": val})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

}

// Donee
func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username and password can not be empty"})
	}
	user.ID = primitive.NewObjectID()
	user.Role = "user"

	if user, err := data.RegisterUser(user); err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Account successfully Created", "data": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	}
}

// DOnee
func LoginUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username and password can not be empty"})
	}

	ExistingUser, err := data.LoginUser(user)
	fmt.Println("ExistingUser: ", ExistingUser)
	if err == nil {
		if bcrypt.CompareHashAndPassword([]byte(ExistingUser.Password), []byte(user.Password)) != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			c.Abort()
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":  ExistingUser.ID,
			"username": ExistingUser.UserName,
			"role":     ExistingUser.Role,
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Add(time.Hour * 2).Unix(),
		})

		jwtToken, err := token.SignedString(secrate_key)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully logged in", "jwt-Token": jwtToken})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}
