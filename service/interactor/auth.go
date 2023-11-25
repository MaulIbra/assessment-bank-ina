package interactor

import (
	"errors"
	"net/http"

	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (i *Interactor) Login(context *gin.Context) {
	var wrapper models.Wrapper
	var auth models.Auth
	if err := context.ShouldBindJSON(&auth); err != nil {
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

	user, token, err := i.AuthUsecase.Login(auth)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}

	wrapper.Data = user
	wrapper.MetaData = &models.MetaData{
		Token: token,
	}
	context.JSON(http.StatusCreated, wrapper)
}
