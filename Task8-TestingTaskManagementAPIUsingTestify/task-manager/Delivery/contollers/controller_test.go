package controller

import (
	domain "TaskManager/task-manager/Domain"
	"TaskManager/task-manager/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ControllerSuite struct {
	suite.Suite
	TaskUsecase *mocks.TaskUsecase
	UserUsecase *mocks.UserUsecase
	controller  *Controller
	router      *gin.Engine
}

func (cs *ControllerSuite) SetupSuite() {
	cs.TaskUsecase = new(mocks.TaskUsecase)
	cs.UserUsecase = new(mocks.UserUsecase)
	cs.controller = NewController(cs.UserUsecase, cs.TaskUsecase)
	// &Controller{
	// 	UserUsecase: cs.UserUsecase,
	// 	TaskUsecase: cs.TaskUsecase,
	// }
	cs.router = gin.Default()
	cs.router.POST("/register", cs.controller.RegisterUser)
	cs.router.POST("/login", cs.controller.LoginUser)
	// cs.router.GET("/tasks", cs.controller.GetAllTasks)
}

func (cs *ControllerSuite) SetupTest() {
	cs.TaskUsecase = new(mocks.TaskUsecase)
	cs.UserUsecase = new(mocks.UserUsecase)
	cs.controller = &Controller{
		UserUsecase: cs.UserUsecase,
		TaskUsecase: cs.TaskUsecase,
	}
	cs.router = gin.Default()
	cs.router.POST("/register", cs.controller.RegisterUser)
	cs.router.POST("/login", cs.controller.LoginUser)
}

// Register User
func (cs *ControllerSuite) TestRegisterUser() {
	user := domain.User{
		UserName: "x",
		Password: "x",
	}

	cs.UserUsecase.On("RegisterUser", user).Return(user, nil)

	jsonPayload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	cs.router.ServeHTTP(w, req)

	cs.Equal(http.StatusAccepted, w.Code, "expected status code 202")

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)

	cs.NoError(err, "should not be error")

	cs.UserUsecase.AssertExpectations(cs.T())
}

func (cs *ControllerSuite) TestRegisterUser_InvalidJSON() {
	// Create an invalid request with incorrect JSON payload
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Act: Serve the HTTP request
	cs.router.ServeHTTP(w, req)

	// Assert
	cs.Equal(http.StatusInternalServerError, w.Code, "expected status code 500")

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	cs.NoError(err, "should not be error")
	// cs.Contains(response["message"], "error decoding JSON", "response message should contain JSON error")

	// Verify mocks
	cs.UserUsecase.AssertExpectations(cs.T())
}

func (cs *ControllerSuite) TestRegisterUser_Error() {
	user := domain.User{
		UserName: "x",
		Password: "x",
	}

	// Mock behavior of RegisterUser to return an error
	cs.UserUsecase.On("RegisterUser", user).Return(domain.User{}, errors.New("some error"))

	// Create the request with JSON payload
	jsonPayload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Act: Serve the HTTP request
	cs.router.ServeHTTP(w, req)

	// Assert
	// cs.Equal(http.StatusBadRequest, w.Code, "expected status code 400")

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	cs.NoError(err, "should not be error")
	cs.Contains(response["message"], "some error", "response message should contain the error message")

	// Verify mock expectations
	cs.UserUsecase.AssertExpectations(cs.T())
}

func (cs *ControllerSuite) TestLoginUser_Success() {
	user := domain.User{
		UserName: "testuser",
		Password: "password123",
	}

	token := "mockToken"

	// Mock the LoginUser method to return a token without error
	cs.UserUsecase.On("LoginUser", user).Return(token, nil)

	// Create the request with JSON payload
	jsonPayload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Act: Serve the HTTP request
	cs.router.ServeHTTP(w, req)
	fmt.Println("//////////////////////////////////////////////")
	fmt.Println("Response Body:", w.Body.String())
	// Assert
	cs.Equal(http.StatusAccepted, w.Code, "expected status code 202")

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	cs.NoError(err, "should not be error")
	cs.Equal(token, response["message"], "response message should match the token")

	// Verify mock expectations
	cs.UserUsecase.AssertExpectations(cs.T())
}

// Register Admin
func (cs *ControllerSuite) TestRegisterAdmin() {
	user := domain.User{
		UserName: "x",
		Password: "x",
	}

	cs.UserUsecase.On("RegisterUser", user).Return(user, nil)

	jsonPayload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	cs.router.ServeHTTP(w, req)

	cs.Equal(http.StatusAccepted, w.Code, "expected status code 202")

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)

	cs.NoError(err, "should not be error")

	cs.UserUsecase.AssertExpectations(cs.T())
}

// Get All Tasks
func (cs *ControllerSuite) TestGetAllTasks_positive() {

	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "UserName",
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	tasks := []domain.Task{
		{ID: primitive.ObjectID{}, Name: "Test Task", Creater_ID: primitive.ObjectID{}},
		{ID: primitive.ObjectID{}, Name: "x", Creater_ID: primitive.ObjectID{}},
	}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("GetAllTasks", "user", "1").Return(tasks, nil)

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks", middleware, cs.controller.GetAllTasks)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusOK, w.Code)
	// Assert the response body contains the expected tasks
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	// fmt.Println("///////////////////////////////")
	// fmt.Println(response["tasks:"])
	// fmt.Println("///////////////////////////////")

	tasksResponse, ok := response["tasks:"].([]interface{})
	if !ok {
		cs.T().Fatal("Expected 'tasks' in response")
	}

	// Convert tasks to a more workable format for comparison
	var tasksList []domain.Task
	for _, task := range tasksResponse {
		taskMap, ok := task.(map[string]interface{})
		if !ok {
			cs.T().Fatal("Invalid task format in response")
		}
		// Create a domain.Task from the map for comparison
		id, _ := primitive.ObjectIDFromHex(taskMap["ID"].(string))
		name, _ := taskMap["Name"].(string)
		creator_id, _ := primitive.ObjectIDFromHex(taskMap["Creater_ID"].(string))

		t := domain.Task{
			ID:         id,
			Name:       name,
			Creater_ID: creator_id,
		}
		tasksList = append(tasksList, t)
	}
	cs.Equal(tasks, tasksList)
	cs.TaskUsecase.AssertExpectations(cs.T())
}

func (cs *ControllerSuite) TestGetAllTasks_noTasks() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "UserName",
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// Mock the TaskUsecase method to return an empty slice
	cs.TaskUsecase.On("GetAllTasks", "user", "1").Return([]domain.Task{}, nil)

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks", middleware, cs.controller.GetAllTasks)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusOK, w.Code)
	// Ensure the response body contains the "no tasks" message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Equal("this user got no tasks", response["tasks:"].(string))
}

func (cs *ControllerSuite) TestGetAllTasks_taskRetrievalFailure() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "UserName",
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// Mock the TaskUsecase method to return an error
	cs.TaskUsecase.On("GetAllTasks", "user", "1").Return([]domain.Task{}, errors.New("database error"))

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks", middleware, cs.controller.GetAllTasks)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message:"], "database error")
}

func (cs *ControllerSuite) TestGetAllTasks_invalidClaims() {
	// Setup claims of incorrect type
	invalidClaims := struct{}{}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("GetAllTasks", "user", "1").Return([]domain.Task{}, nil)

	// Create middleware that sets invalid claims
	middleware := func(c *gin.Context) {
		c.Set("claims", invalidClaims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks", middleware, cs.controller.GetAllTasks)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "could not converts claims")
}

func (cs *ControllerSuite) TestGetAllTasks_missingClaims() {
	// Mock the TaskUsecase method
	cs.TaskUsecase.On("GetAllTasks", "user", "1").Return([]domain.Task{}, nil)

	// Create middleware that does not set claims
	middleware := func(c *gin.Context) {
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks", middleware, cs.controller.GetAllTasks)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "could not find token claims")
}

// Get Tasks By ID

func (cs *ControllerSuite) TestGetTaskByID_success() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "UserName",
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	task := domain.Task{
		ID:         primitive.NewObjectID(),
		Name:       "Test Task",
		Creater_ID: primitive.NewObjectID(),
	}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("GetTaskByID", "user", "1", "taskID").Return(task, nil)

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks/:id", middleware, cs.controller.GetTaskByID)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/taskID", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusOK, w.Code)
	// Ensure the response body contains the expected task
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}

	taskResponse, ok := response["tasks:"].(map[string]interface{})
	if !ok {
		cs.T().Fatal("Expected 'tasks' in response")
	}

	// Create a domain.Task from the map for comparison
	id, _ := primitive.ObjectIDFromHex(taskResponse["ID"].(string))
	name, _ := taskResponse["Name"].(string)
	creatorID, _ := primitive.ObjectIDFromHex(taskResponse["Creater_ID"].(string))

	t := domain.Task{
		ID:         id,
		Name:       name,
		Creater_ID: creatorID,
	}

	cs.Equal(task, t)
	cs.TaskUsecase.AssertExpectations(cs.T())
}

func (cs *ControllerSuite) TestGetTaskByID_missingClaims() {
	// Mock the TaskUsecase method
	cs.TaskUsecase.On("GetTaskByID", "user", "1", "taskID").Return(domain.Task{}, nil)

	// Create middleware that does not set claims
	middleware := func(c *gin.Context) {
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks/:id", middleware, cs.controller.GetTaskByID)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/taskID", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "could not find token claims")
}

func (cs *ControllerSuite) TestGetTaskByID_invalidClaims() {
	// Setup invalid claims
	invalidClaims := struct{}{}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("GetTaskByID", "user", "1", "taskID").Return(domain.Task{}, nil)

	// Create middleware that sets invalid claims
	middleware := func(c *gin.Context) {
		c.Set("claims", invalidClaims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks/:id", middleware, cs.controller.GetTaskByID)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/taskID", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "could not converts claims")
}

func (cs *ControllerSuite) TestGetTaskByID_taskRetrievalFailure() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "UserName",
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// Mock the TaskUsecase method to return an error
	cs.TaskUsecase.On("GetTaskByID", "user", "1", "taskID").Return(domain.Task{}, errors.New("database error"))

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.GET("/tasks/:id", middleware, cs.controller.GetTaskByID)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/taskID", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message:"], "database error")
}

// Add Task
func (cs *ControllerSuite) TestAddTask_success() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "UserName",
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	task := domain.Task{
		Name: "Test Task",
	}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("AddTask", "user", "1", task).Return(task.ID.Hex(), nil)

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.POST("/tasks", middleware, cs.controller.AddTask)

	// Create and send request
	taskJSON, _ := json.Marshal(task)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusAccepted, w.Code)
	// Ensure the response body contains the expected task ID
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}

	insertedID, ok := response["Task Insered at ID"].(string)
	if !ok {
		cs.T().Fatal("Expected 'Task Insered at ID' in response")
	}

	cs.Equal(task.ID.Hex(), insertedID)
	cs.TaskUsecase.AssertExpectations(cs.T())
}

func (cs *ControllerSuite) TestAddTask_bindingFailure() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "UserName",
		"role":     "user",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.POST("/tasks", middleware, cs.controller.AddTask)

	// Create and send request with invalid JSON body
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader("invalid json"))
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "invalid character")
}

func (cs *ControllerSuite) TestAddTask_invalidClaims() {
	// Setup invalid claims
	invalidClaims := struct{}{}

	// Create a sample task
	task := domain.Task{Name: "New Task", Creater_ID: primitive.NewObjectID()}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("AddTask", "user", "1", task).Return(primitive.NewObjectID(), nil)

	// Create middleware that sets invalid claims
	middleware := func(c *gin.Context) {
		c.Set("claims", invalidClaims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.POST("/tasks", middleware, cs.controller.AddTask)

	// Create and send request
	body, _ := json.Marshal(task)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "could not converts claims")
}

func (cs *ControllerSuite) TestAddTask_missingClaims() {
	// Create a sample task
	task := domain.Task{Name: "New Task", Creater_ID: primitive.NewObjectID()}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("AddTask", "user", "1", task).Return(primitive.NewObjectID(), nil)

	// Create middleware that does not set claims
	middleware := func(c *gin.Context) {
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.POST("/tasks", middleware, cs.controller.AddTask)

	// Create and send request
	body, _ := json.Marshal(task)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "could not find token claims")
}

// update Task By ID

// TestUpdateTaskBy ID_invalidJSON
func (cs *ControllerSuite) TestUpdateTaskByID_invalidClaims() {
	// Setup invalid claims
	invalidClaims := struct{}{}

	// Create a sample task
	task := domain.Task{Name: "New Task", Creater_ID: primitive.NewObjectID()}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("UpdateTaskByID", "user", "1", task, "taskID").Return(nil)

	// Create middleware that sets invalid claims
	middleware := func(c *gin.Context) {
		c.Set("claims", invalidClaims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.PUT("/tasks/:id", middleware, cs.controller.UpdateTaskByID)

	// Create and send request
	body, _ := json.Marshal(task)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/taskID", bytes.NewReader(body))
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusInternalServerError, w.Code)
	// Ensure the response body contains the error message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}
	cs.Contains(response["message"], "could not converts claims")
}

// TestUpdateTaskBy ID_success
func (cs *ControllerSuite) TestUpdateTaskByID_success() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id": "1",
		"role":    "user",
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}

	task := domain.Task{
		Name:       "Updated Task",
		Creater_ID: primitive.NewObjectID(),
	}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("UpdateTaskByID", "user", "1", task, "taskID").Return(nil)

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.PUT("/tasks/:id", middleware, cs.controller.UpdateTaskByID)

	// Create and send request
	taskJSON, _ := json.Marshal(task)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/taskID", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusAccepted, w.Code)
	// Ensure the response body contains the success message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}

	message, ok := response["Message"].(string)
	if !ok || message != "successfully updated" {
		cs.T().Fatal("Expected 'Message' with value 'successfully updated' in response")
	}

	cs.TaskUsecase.AssertExpectations(cs.T())
}

// Delete Task By ID
// TestDeleteTaskBy ID_success
func (cs *ControllerSuite) TestDeleteTaskByID_success() {
	// Setup JWT claims
	claims := jwt.MapClaims{
		"user_id": "1",
		"role":    "user",
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}

	// Mock the TaskUsecase method
	cs.TaskUsecase.On("DeleteTaskByID", "user", "1", "taskID").Return(nil)

	// Create middleware that sets claims
	middleware := func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}

	// Setup router with middleware and controller
	cs.router.DELETE("/tasks/:id", middleware, cs.controller.DeleteTaskByID)

	// Create and send request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/taskID", nil)
	cs.router.ServeHTTP(w, req)

	// Assertions
	cs.Equal(http.StatusAccepted, w.Code)
	// Ensure the response body contains the success message
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		cs.T().Fatal(err)
	}

	message, ok := response["Message"].(string)
	if !ok || message != "successfully Deleted" {
		cs.T().Fatal("Expected 'Message' with value 'successfully Deleted' in response")
	}

	cs.TaskUsecase.AssertExpectations(cs.T())
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}
