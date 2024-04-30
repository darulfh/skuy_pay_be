package repository

import (
	"BE-Golang/model"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterRepository(user model.User) (*model.User, error)
	LoginRepository(user *model.User) (*model.User, error)
	GetUserByEmailRepository(email string) (*model.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) RegisterRepository(user model.User) (*model.User, error) {
	var count int64
	r.db.Model(&model.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		return nil, errors.New("email already exists")
	}
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) LoginRepository(user *model.User) (*model.User, error) {
	result := r.db.Where("email = ? AND password = ?", user.Email, user.Password).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("incorrect email or password")
		}
		return nil, errors.New("failed to get user")
	}

	return user, nil
}

func (r *authRepository) GetUserByEmailRepository(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
