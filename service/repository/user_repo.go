package repository

import (
	"github.com/MaulIbra/assessment-bank-ina/service/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	CreateUser(user *models.User) error
	GetUsers() (*[]models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{
		db: db,
	}
}

func (ur userRepo) CreateUser(user *models.User) error {
	return ur.db.Model(models.User{}).Create(&user).Error
}

func (ur userRepo) GetUsers() (*[]models.User, error) {
	users := make([]models.User, 0)
	result := ur.db.Model(models.User{}).Find(&users)
	return &users, result.Error
}

func (ur userRepo) GetUser(id int) (*models.User, error) {
	user := models.User{}
	result := ur.db.Model(models.User{}).Where("id = ?", id).Find(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, result.Error
}

func (ur userRepo) UpdateUser(user *models.User) error {
	return ur.db.Model(models.User{}).Where("id = ?", user.ID).Updates(&user).Error
}

func (ur userRepo) DeleteUser(id int) error {
	return ur.db.Model(models.User{}).Delete(models.User{}, id).Error
}

func (ur userRepo) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	result := ur.db.Model(models.User{}).Where("email = ?", email).Find(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, result.Error
}
