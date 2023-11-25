package interactor

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (i *Interactor) CreateUser(context *gin.Context) {
	var wrapper models.Wrapper
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
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
	err := i.UserUsecase.CreateUser(&user)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == "user email already exist" {
			context.JSON(http.StatusBadRequest, gin.H{
				"messages": err.Error(),
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"messages": "something wrong in the server",
			})
		}
		return
	}

	token, err := i.AuthRepo.GenerateJWT(user.ID, user.Email, user.Password)
	if err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}

	wrapper.Data = user
	wrapper.MetaData = &models.MetaData{
		Token: token,
	}
	context.JSON(http.StatusCreated, wrapper)
}

func (i *Interactor) ReadUsers(context *gin.Context) {
	var wrapper models.Wrapper
	users, err := i.UserUsecase.GetUsers()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}
	wrapper.Data = users
	context.JSON(http.StatusOK, wrapper)
}

func (i *Interactor) ReadUser(context *gin.Context) {
	var wrapper models.Wrapper
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "id can't be empty",
		})
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	user, err := i.UserUsecase.GetUser(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "something wrong in the server",
		})
		return
	}
	if user == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"messages": "no data",
		})
		return
	}

	wrapper.Data = user

	context.JSON(http.StatusOK, wrapper)
}

func (i *Interactor) UpdateUser(context *gin.Context) {
	var wrapper models.Wrapper
	var user models.User
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "id can't be empty",
		})
		return
	}

	UserId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	if err := context.ShouldBindJSON(&user); err != nil {
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

	user.ID = UserId

	err = i.UserUsecase.UpdateUser(&user)
	if err != nil {
		if err.Error() == "user id is not exist" {
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
	wrapper.Data = user
	context.JSON(http.StatusOK, wrapper)
}

func (i *Interactor) DeleteUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": "id can't be empty",
		})
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	err = i.UserUsecase.DeleteUser(userId)
	if err != nil {
		if err.Error() == "user id is not exist" {
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
		"messages": "User sucessfully deleted",
	})
}
