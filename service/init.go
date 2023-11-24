package service

import (
	"os"
	"strconv"

	"github.com/MaulIbra/assessment-bank-ina/config"
	"github.com/MaulIbra/assessment-bank-ina/service/interactor"
	"github.com/MaulIbra/assessment-bank-ina/service/repository"
	"github.com/MaulIbra/assessment-bank-ina/service/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func Init() *gin.Engine {

	router := gin.Default()

	db := config.DbConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}.MysqlConnection()

	log.SetReportCaller(true)

	mySigningKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	expiredToken, _ := strconv.Atoi(os.Getenv("EXPIRED_TIME_TOKEN"))

	passSecret := os.Getenv("PASS_SECRET")

	v1 := router.Group("/api")

	iAuthRepo := repository.NewAuthRepository(token, mySigningKey, expiredToken)

	iUserRepo := repository.NewUserRepo(db)
	iUserUsecase := usecase.NewUserUsecase(iUserRepo, passSecret)

	interactor := interactor.Interactor{
		UserUsecase: iUserUsecase,
		AuthRepo:    iAuthRepo,
	}
	interactor.Routes(v1)

	return router
}
