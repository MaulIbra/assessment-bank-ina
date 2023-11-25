package usecase

import (
	"errors"

	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"github.com/MaulIbra/assessment-bank-ina/service/repository"
)

type ITaskUsecase interface {
	CreateTask(task *models.Task) error
	GetTasks() (*[]models.Task, error)
	GetTask(id int) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	taskRepo repository.ITaskRepo
}

func NewTaskUsecase(taskRepo repository.ITaskRepo) ITaskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}

func (uu taskUsecase) CreateTask(task *models.Task) error {
	if task.Status == "" {
		task.Status = "PENDING"
	}
	return uu.taskRepo.CreateTask(task)
}

func (uu taskUsecase) GetTasks() (*[]models.Task, error) {
	return uu.taskRepo.GetTasks()
}

func (uu taskUsecase) GetTask(id int) (*models.Task, error) {
	return uu.taskRepo.GetTask(id)
}
func (uu taskUsecase) UpdateTask(task *models.Task) error {
	taskTemp, err := uu.taskRepo.GetTask(task.ID)
	if err != nil {
		return err
	}
	if taskTemp == nil {
		return errors.New("task id is not exist")
	}

	taskTemp.Title = task.Title
	taskTemp.Status = task.Status
	taskTemp.Description = task.Description
	taskTemp.UserId = task.UserId
	if task.Status == "" {
		taskTemp.Status = "PENDING"
	}

	err = uu.taskRepo.UpdateTask(taskTemp)
	if err != nil {
		return err
	}
	return nil
}

func (uu taskUsecase) DeleteTask(id int) error {
	task, err := uu.taskRepo.GetTask(id)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("task id is not exist")
	}
	return uu.taskRepo.DeleteTask(id)
}
