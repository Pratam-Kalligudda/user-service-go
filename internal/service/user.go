package service

import (
	"github.com/Pratam-Kalligudda/user-service-go/internal/dto"
	"github.com/Pratam-Kalligudda/user-service-go/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s UserService) Login(u dto.LoginDTO) (string, error) {
	return "", nil
}

func (s UserService) Register(u dto.SignupDTO) (string, error) {
	return "", nil
}

func (s UserService) UpdateUser(u any) error {
	return nil
}

func (s UserService) Refresh(u any) error {
	return nil
}
