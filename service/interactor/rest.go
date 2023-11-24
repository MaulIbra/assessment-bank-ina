package interactor

import (
	"net/http"
	"strconv"

	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"github.com/MaulIbra/assessment-bank-ina/service/repository"
	"github.com/MaulIbra/assessment-bank-ina/service/usecase"
	"github.com/gin-gonic/gin"
)

type Interactor struct {
	UserUsecase usecase.IUserUsecase
	AuthRepo    repository.IAuthRepo
}

func Auth(repository repository.IAuthRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "you are unauthorized to make this request.",
			})
			c.Abort()
		} else {
			err := repository.ClaimToken(token)
			if err != nil {
				c.AbortWithStatusJSON(401, gin.H{
					"message": "you are unauthorized to make this request.",
				})
			}
			c.Next()
		}
	}
}

func (i *Interactor) Routes(router *gin.RouterGroup) {
	router.POST("/user", i.CreateUser)
	router.GET("/user", Auth(i.AuthRepo), i.ReadUsers)
	router.GET("/user/:id", Auth(i.AuthRepo), i.ReadUser)
	router.PUT("/user/:id", Auth(i.AuthRepo), i.UpdateUser)
	router.DELETE("/user/:id", Auth(i.AuthRepo), i.DeleteUser)
}

func (i *Interactor) CreateUser(context *gin.Context) {
	var wrapper models.Wrapper
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}
	err = i.UserUsecase.CreateUser(&user)
	if err != nil {
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

	err = context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
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
