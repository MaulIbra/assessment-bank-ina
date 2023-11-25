package usecase

import (
	"errors"

	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"github.com/MaulIbra/assessment-bank-ina/service/repository"
	"github.com/MaulIbra/assessment-bank-ina/service/utils"
)

type IAuthUsecase interface {
	Login(auth models.Auth) (*models.User, string, error)
}

type authUsecase struct {
	authRepo   repository.IAuthRepo
	userRepo   repository.IUserRepo
	passSecret string
}

func NewAuthUsecase(authRepo repository.IAuthRepo, userRepo repository.IUserRepo, passSecret string) IAuthUsecase {
	return &authUsecase{
		authRepo:   authRepo,
		userRepo:   userRepo,
		passSecret: passSecret,
	}
}

func (au authUsecase) Login(auth models.Auth) (*models.User, string, error) {
	user, err := au.userRepo.GetUserByEmail(auth.Email)
	if err != nil {
		return nil, "", errors.New("You have entered an invalid email or password")
	}

	if user == nil {
		return nil, "", errors.New("You have entered an invalid email or password")
	}

	pass, err := utils.Decrypt(user.Password, au.passSecret)
	if err != nil {
		return nil, "", err
	}

	if pass != auth.Password {
		return nil, "", errors.New("You have entered an invalid email or password")
	}

	token, err := au.authRepo.GenerateJWT(user.ID, user.Email, user.Password)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
