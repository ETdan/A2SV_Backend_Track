package usecases

import (
	domain "TaskManager/task-manager/Domain"
	"fmt"
	"time"
)

type TaskUsecase struct {
	TaskRepository domain.TaskRepository
}

func NewTaskUsecase(repository domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		TaskRepository: repository,
	}
}

func (t *TaskUsecase) GetAllTasks(role string, user_id string) ([]domain.Task, error) {
	data, err := t.TaskRepository.GetAllTasks(role, user_id, time.Duration(time.Second))
	return data, err
}
func (t *TaskUsecase) GetTaskByID(role string, user_id string, task_id string) (domain.Task, error) {
	data, err := t.TaskRepository.GetTaskByID(role, user_id, task_id, time.Duration(time.Second))
	return data, err
}
func (t *TaskUsecase) AddTask(role string, user_id string, task domain.Task) (string, error) {
	fmt.Println(user_id)
	insertedid, err := t.TaskRepository.AddTask(user_id, task, time.Second)
	if err == nil {
		return insertedid, nil
	} else {
		return "", err
	}
}
func (t *TaskUsecase) UpdateTaskByID(role string, user_id string, task domain.Task, id string) error {
	result := t.TaskRepository.UpdateTaskByID(role, user_id, time.Second, task, id)
	return result
}
func (t *TaskUsecase) DeleteTaskByID(role string, user_id string, task_id string) error {
	result := t.TaskRepository.DeleteTaskByID(role, user_id, task_id, time.Second)
	return result
}
