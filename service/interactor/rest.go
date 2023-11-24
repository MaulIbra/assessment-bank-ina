package interactor

import (
	"net/http"

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
			claim, err := repository.ClaimToken(token)
			if err != nil {
				c.AbortWithStatusJSON(401, gin.H{
					"message": "you are unauthorized to make this request.",
				})
			}
			userId, ok := claim["user_id"]
			if ok {
				c.Request.Header.Add("user_id", userId.(string))
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
