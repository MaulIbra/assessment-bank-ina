package usecase

import (
	"errors"

	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"github.com/MaulIbra/assessment-bank-ina/service/repository"
	"github.com/MaulIbra/assessment-bank-ina/service/utils"
)

type IUserUsecase interface {
	CreateUser(user *models.User) error
	GetUsers() (*[]models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userUsecase struct {
	userRepo   repository.IUserRepo
	passSecret string
}

func NewUserUsecase(userRepo repository.IUserRepo, passSecret string) IUserUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		passSecret: passSecret,
	}
}

func (uu userUsecase) CreateUser(user *models.User) error {

	userTemp, err := uu.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if userTemp != nil {
		return errors.New("user email already exist")
	}
	encryptedPass, err := utils.Encrypt(user.Password, uu.passSecret)
	if err != nil {
		return err
	}
	user.Password = encryptedPass
	return uu.userRepo.CreateUser(user)
}

func (uu userUsecase) GetUsers() (*[]models.User, error) {
	return uu.userRepo.GetUsers()
}

func (uu userUsecase) GetUser(id int) (*models.User, error) {
	return uu.userRepo.GetUser(id)
}
func (uu userUsecase) UpdateUser(user *models.User) error {
	userTemp, err := uu.userRepo.GetUser(user.ID)
	if err != nil {
		return err
	}
	if userTemp == nil {
		return errors.New("user id is not exist")
	}

	pass, err := utils.Decrypt(userTemp.Password, uu.passSecret)
	if err != nil {
		return err
	}

	if user.Password != "" && pass != user.Password {
		encryptPass, err := utils.Encrypt(user.Password, uu.passSecret)
		if err != nil {
			return err
		}
		userTemp.Password = encryptPass
	}

	userTemp.Name = user.Name
	userTemp.Email = user.Email

	err = uu.userRepo.UpdateUser(userTemp)
	if err != nil {
		return err
	}
	return nil
}

func (uu userUsecase) DeleteUser(id int) error {
	user, err := uu.userRepo.GetUser(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user id is not exist")
	}
	return uu.userRepo.DeleteUser(id)
}
