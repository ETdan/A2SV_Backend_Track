package repositories

import (
	db "TaskManager/task-manager/DB"
	domain "TaskManager/task-manager/Domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRpository struct {
	database db.Database
}

func NewTaskRepository(database db.Database) *TaskRpository {
	return &TaskRpository{
		database: database,
	}
}
func (t *TaskRpository) GetAllTasks(role string, user_id string, duration time.Duration) ([]domain.Task, error) {
	context, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	var data []domain.Task
	var task domain.Task
	if role == "admin" {
		cursor, err := t.database.GetAllTasks(context, t.database.FilterEmpty())
		if err == nil {
			for cursor.Next(context) {
				err := cursor.Decode(&task)
				if err == nil {
					data = append(data, task)
				} else {
					return []domain.Task{}, err
				}
			}
			// fmt.Println("admin //////////////////////")
			// fmt.Println(data)
			return data, nil
		}
		return nil, err
	} else {
		user_object_id, er := primitive.ObjectIDFromHex(user_id)
		if er != nil {
			return []domain.Task{}, er
		}
		cursor, err := t.database.GetAllTasks(context, t.database.FilterByCreaterID(user_object_id))
		if err == nil {
			for cursor.Next(context) {
				err := cursor.Decode(&task)
				if err == nil {
					data = append(data, task)
				} else {
					return []domain.Task{}, err
				}
			}
			return data, nil
		}
		return nil, err
	}
}
func (t *TaskRpository) GetTaskByID(role string, user_id string, task_id string, duration time.Duration) (domain.Task, error) {
	context, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	// var task domain.Task
	if role == "admin" {
		task_object_id, er := primitive.ObjectIDFromHex(task_id)
		if er != nil {
			return domain.Task{}, er
		}
		task, err := t.database.GetTaskByID(context, t.database.FilterByTaskID(task_object_id))
		if err == nil {
			return task, nil
		} else {

			return domain.Task{}, err
		}
	} else {
		task_object_id, er := primitive.ObjectIDFromHex(task_id)
		if er != nil {
			return domain.Task{}, er
		}
		user_object_id, er := primitive.ObjectIDFromHex(user_id)
		if er == nil {
			filter := primitive.D{}
			filter = append(filter, t.database.FilterByTaskID(task_object_id)...)
			filter = append(filter, t.database.FilterByCreaterID(user_object_id)...)
			task, err := t.database.GetTaskByID(context, filter)
			if err == nil {
				return task, nil
			} else {
				return domain.Task{}, err
			}
		} else {
			return domain.Task{}, er
		}
	}
}
func (t *TaskRpository) AddTask(user_id string, task domain.Task, duration time.Duration) (string, error) {
	context, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	var err error
	task.Creater_ID, err = primitive.ObjectIDFromHex(user_id)
	task.ID = primitive.NewObjectID()
	if err == nil {
		InsertedID, err := t.database.PostTask(context, task)
		if err == nil {
			return InsertedID, nil
		}
		return "", err
	}
	return "", err
}
func (t *TaskRpository) UpdateTaskByID(role string, user_id string, time_duration time.Duration, task domain.Task, id string) error {
	context, cancel := context.WithTimeout(context.Background(), time_duration)
	defer cancel()
	// task_id := task.ID.Hex()
	if role == "admin" {
		_, err := t.database.UpdateTaskByID(context, t.database.FilterEmpty(), task)
		return err
	} else {
		object_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		existing_task, e := t.database.GetTaskByID(context, t.database.FilterByTaskID(object_id))
		if e != nil {
			return e
		}

		user_object_id, err := primitive.ObjectIDFromHex(user_id)
		if err == nil {
			if existing_task.Creater_ID == user_object_id {
				_, err := t.database.UpdateTaskByID(context, t.database.FilterByTaskID(object_id), task)
				return err
			} else {
				return errors.New("task does not belong to user")
			}
		} else {
			return err
		}
	}
}
func (t *TaskRpository) DeleteTaskByID(role string, user_id string, task_id string, time_duration time.Duration) error {
	context, cancel := context.WithTimeout(context.Background(), time_duration)
	defer cancel()

	if role == "admin" {
		if task_object_id, er := primitive.ObjectIDFromHex(task_id); er == nil {
			f := t.database.FilterByTaskID(task_object_id)
			count, err := t.database.DeleteTask(context, f)
			if count == 0 {
				return errors.New("no document with this id was found and none was deleted")
			}
			return err
		} else {
			return er
		}
	} else {
		task_object_id, err := primitive.ObjectIDFromHex(task_id)
		if err != nil {
			return err
		}
		f := t.database.FilterByTaskID(task_object_id)
		existing_task, e := t.database.GetTaskByID(context, f)
		// fmt.Println(task_object_id)
		if e != nil {
			return e
		}

		user_object_id, err := primitive.ObjectIDFromHex(user_id)
		if err == nil {
			if existing_task.Creater_ID == user_object_id {
				f := t.database.FilterByTaskID(task_object_id)
				_, err := t.database.DeleteTask(context, f)
				return err
			} else {
				return errors.New("task does not belong to user")
			}
		} else {
			return err
		}
	}
}
