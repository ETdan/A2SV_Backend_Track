package usecases

import (
	domain "TaskManager/task-manager/Domain"
	"TaskManager/task-manager/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TaskUsecaseSuite struct {
	suite.Suite
	TaskRepository *mocks.TaskRepository
	TaskUsecase    *TaskUsecase
}

func (tr *TaskUsecaseSuite) SetupSuite() {
	tr.TaskRepository = new(mocks.TaskRepository)
	tr.TaskUsecase = NewTaskUsecase(tr.TaskRepository)
}

func (tr *TaskUsecaseSuite) SetupTest() {
	tr.TaskRepository = new(mocks.TaskRepository)
	tr.TaskUsecase = NewTaskUsecase(tr.TaskRepository)
}

func (tu *TaskUsecaseSuite) TestDeleteTaskByID() {
	tu.TaskRepository.On("DeleteTaskByID", "user", "1", "1", time.Second).Return(nil)
	err := tu.TaskUsecase.DeleteTaskByID("user", "1", "1")

	tu.Nil(err, "should be nil")
	tu.TaskRepository.AssertExpectations(tu.T())
}

func (tu *TaskUsecaseSuite) TestUpdateTaskByID() {
	tu.TaskRepository.On("UpdateTaskByID", "user", "1", time.Second, domain.Task{}, "1").Return(nil)
	err := tu.TaskUsecase.UpdateTaskByID("user", "1", domain.Task{}, "1")

	tu.Nil(err, "should be nil")
	tu.TaskRepository.AssertExpectations(tu.T())
}

func (tu *TaskUsecaseSuite) TestAddTask() {
	tu.TaskRepository.On("AddTask", "1", domain.Task{}, time.Second).Return("1", nil)
	inserted_id, err := tu.TaskUsecase.AddTask("user", "1", domain.Task{})

	tu.Nil(err, "should be nil")
	tu.Equal(inserted_id, "1", "should be equal")
	tu.TaskRepository.AssertExpectations(tu.T())
}

func (tu *TaskUsecaseSuite) TestGetTaskByID() {
	tu.TaskRepository.On("GetTaskByID", "user", "1", "1", time.Second).Return(domain.Task{}, nil)
	task, err := tu.TaskUsecase.GetTaskByID("user", "1", "1")

	tu.Nil(err, "should be nil")
	tu.Equal(task, domain.Task{}, "should be nil")
	tu.TaskRepository.AssertExpectations(tu.T())
}

func (tu *TaskUsecaseSuite) TestGetAllTasks() {
	tu.TaskRepository.On("GetAllTasks", "user", "1", time.Second).Return([]domain.Task{}, nil)
	task, err := tu.TaskUsecase.GetAllTasks("user", "1")

	tu.Nil(err, "should be nil")
	tu.Equal(task, []domain.Task{}, "should be nil")
	tu.TaskRepository.AssertExpectations(tu.T())
}

func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
