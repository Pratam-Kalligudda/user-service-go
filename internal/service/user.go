package service

import (
	"errors"
	"time"

	"github.com/Pratam-Kalligudda/user-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/user-service-go/internal/dto"
	"github.com/Pratam-Kalligudda/user-service-go/internal/helper"
	"github.com/Pratam-Kalligudda/user-service-go/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) Login(u dto.LoginDTO) (string, string, error) {
	user, err := s.Repo.FindUserByEmail(u.Email)
	if err != nil {
		return "", "", err
	}

	if err := s.Auth.ComparePassword(user.Password, u.Password); err != nil {
		return "", "", errors.New("incorrect username or password")
	}

	refreshToken, err := s.Auth.GenerateToken(user.ID, user.UserType, user.Email, time.Hour*1)
	if err != nil {
		return "", "", err
	}

	jwtToken, err := s.Auth.GenerateToken(user.ID, user.UserType, user.Email, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	return jwtToken, refreshToken, nil
}

func (s UserService) Register(u dto.SignupDTO) (string, string, error) {

	hashPass, err := s.Auth.GenerateHashPassword(u.Password)
	if err != nil {
		return "", "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    u.Email,
		Password: hashPass,
		Phone:    u.Phone,
	})
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.Auth.GenerateToken(user.ID, user.UserType, user.Email, time.Hour*1)
	if err != nil {
		return "", "", err
	}

	jwtToken, err := s.Auth.GenerateToken(user.ID, user.UserType, user.Email, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	return jwtToken, refreshToken, nil
}

func (s UserService) UpdateUser(u any) error {
	return nil
}

func (s UserService) Refresh(user domain.User) (string, error) {
	token, err := s.Auth.GenerateToken(user.ID, user.Email, user.UserType, time.Minute*15)
	return token, err
}
func (s UserService) GetVerificationCode() error {
	return nil
}
func (s UserService) BecomeSeller(email string) error {
	return nil
}
func (s UserService) VerifyUser() error {
	return nil
}
