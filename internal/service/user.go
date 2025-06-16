package service

import (
	"errors"
	"math"
	"math/rand/v2"
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

	if err := s.Auth.Validate(u); err != nil {
		return "", "", err
	}

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
func (s UserService) GetVerificationCode(userId uint) (int, error) {
	randInt := s.randNumber(6)
	exp := time.Now().Add(time.Minute * 5)

	_, err := s.Repo.UpdateUser(domain.User{
		Code:   randInt,
		Expiry: exp,
	})
	if err != nil {
		return -1, nil
	}

	return randInt, nil
}
func (s UserService) BecomeSeller(email string) error {
	return nil
}
func (s UserService) VerifyUser(verification dto.VerificationCodeDTO, id uint) error {
	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return err
	}
	if user.Verified {
		return errors.New("user already verified")
	}
	if user.Code != verification.Code {
		return errors.New("wrong code")
	}

	if time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	user.Verified = true

	if _, err = s.Repo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
func (s UserService) randNumber(n int) int {
	minLimit := int(math.Pow10(n))
	maxLimit := int(math.Pow10(n - 1))
	randInt := int(rand.Float64() * float64(minLimit))
	if randInt < maxLimit {
		randInt += maxLimit
	}
	return randInt
}
