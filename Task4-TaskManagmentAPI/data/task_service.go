package data

import (
	"TaskManager/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Data = map[string]models.Task{
	"1": {
		ID:       "1",
		Name:     "Task 1",
		Detail:   "Detail for Task 1",
		Start:    "2024-07-31T08:00:00Z",
		Duration: "1h",
	},
	"2": {
		ID:       "2",
		Name:     "Task 2",
		Detail:   "Detail for Task 2",
		Start:    "2024-07-31T09:00:00Z",
		Duration: "2h",
	},
	"3": {
		ID:       "3",
		Name:     "Task 3",
		Detail:   "Detail for Task 3",
		Start:    "2024-07-31T10:00:00Z",
		Duration: "30m",
	},
	"4": {
		ID:       "4",
		Name:     "Task 4",
		Detail:   "Detail for Task 4",
		Start:    "2024-07-31T10:30:00Z",
		Duration: "1h 15m",
	},
}

func AddTask(task models.Task) (models.Task, error) {

	validator := validator.New()
	err := validator.Struct(task)
	if err != nil {
		return models.Task{}, errors.New("a task must have an id")
	}
	_, exists := Data[task.ID]
	if !exists {
		Data[task.ID] = task
		return task, nil
	}

	fmt.Println("no Data")
	return models.Task{}, errors.New("a task with this id exists")
}
func GetAllTask() map[string]models.Task {
	return Data
}
func GetTask(id string) (models.Task, error) {
	val, exists := Data[id]
	if !exists {
		fmt.Println("no Data")
		return models.Task{}, errors.New("no task with this ID")
	}
	return val, nil
}
func DeleteTask(id string) (models.Task, error) {
	if val, ok := Data[id]; ok {
		delete(Data, id)
		return val, nil
	}
	return models.Task{}, errors.New("NO such Task ID")
}
func UpdateTask(task models.Task, id string) (models.Task, error) {

	if val, ok := Data[id]; ok {
		if task.Detail != "" {
			val.Detail = task.Detail
		}
		if task.Duration != "" {
			val.Duration = task.Duration
		}
		if task.ID != "" {
			val.ID = task.ID
		}
		if task.Name != "" {
			val.Name = task.Name
		}
		if task.Start != "" {
			val.Start = task.Start
		}
		Data[id] = val
		return Data[id], nil
	}
	return models.Task{}, errors.New("NO Data Such ID")
}
