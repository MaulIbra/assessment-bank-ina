package interactor

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (i *Interactor) CreateTask(context *gin.Context) {
	var wrapper models.Wrapper
	var task models.Task
	if err := context.ShouldBindJSON(&task); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = models.ErrorMsg{fe.Field(), models.GetErrorMsg(fe)}
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	user := context.Request.Header.Get("user_id")
	if user == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "token is not have specific id",
		})
		return
	}

	userId, err := strconv.Atoi(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	task.UserId = userId
	err = i.TaskUsecase.CreateTask(&task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}
	wrapper.Data = task
	context.JSON(http.StatusCreated, wrapper)
}

func (i *Interactor) ReadTasks(context *gin.Context) {
	var wrapper models.Wrapper
	tasks, err := i.TaskUsecase.GetTasks()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}
	wrapper.Data = tasks
	context.JSON(http.StatusOK, wrapper)
}

func (i *Interactor) ReadTask(context *gin.Context) {
	var wrapper models.Wrapper
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "id can't be empty",
		})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	task, err := i.TaskUsecase.GetTask(taskId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}
	if task == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"messages": "no data",
		})
		return
	}

	wrapper.Data = task
	context.JSON(http.StatusOK, wrapper)
}

func (i *Interactor) UpdateTask(context *gin.Context) {
	var wrapper models.Wrapper
	var task models.Task
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "id can't be empty",
		})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	user := context.Request.Header.Get("user_id")
	if user == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "token is not have specific id",
		})
		return
	}

	userId, err := strconv.Atoi(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}
	if err := context.ShouldBindJSON(&task); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = models.ErrorMsg{fe.Field(), models.GetErrorMsg(fe)}
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	task.UserId = userId
	task.ID = taskId

	err = i.TaskUsecase.UpdateTask(&task)
	if err != nil {
		if err.Error() == "task id is not exist" {
			context.JSON(http.StatusBadRequest, gin.H{
				"messages": err.Error(),
			})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}
	wrapper.Data = task
	context.JSON(http.StatusOK, wrapper)
}

func (i *Interactor) DeleteTask(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "id can't be empty",
		})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	err = i.TaskUsecase.DeleteTask(taskId)
	if err != nil {
		if err.Error() == "task id is not exist" {
			context.JSON(http.StatusBadRequest, gin.H{
				"messages": err.Error(),
			})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"messages": "Task sucessfully deleted",
	})
}
