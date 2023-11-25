package repository

import (
	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"gorm.io/gorm"
)

type ITaskRepo interface {
	CreateTask(Task *models.Task) error
	GetTasks() (*[]models.Task, error)
	GetTask(id int) (*models.Task, error)
	UpdateTask(Task *models.Task) error
	DeleteTask(id int) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) ITaskRepo {
	return &taskRepo{
		db: db,
	}
}

func (ur taskRepo) CreateTask(task *models.Task) error {
	return ur.db.Model(models.Task{}).Create(&task).Error
}

func (ur taskRepo) GetTasks() (*[]models.Task, error) {
	tasks := make([]models.Task, 0)
	result := ur.db.Model(models.Task{}).Find(&tasks)
	return &tasks, result.Error
}

func (ur taskRepo) GetTask(id int) (*models.Task, error) {
	task := models.Task{}
	result := ur.db.Model(models.Task{}).Where("id = ?", id).Find(&task)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &task, result.Error
}

func (ur taskRepo) UpdateTask(task *models.Task) error {
	return ur.db.Model(models.Task{}).Where("id = ?", task.ID).Updates(&task).Error
}

func (ur taskRepo) DeleteTask(id int) error {
	return ur.db.Model(models.Task{}).Delete(models.Task{}, id).Error
}
