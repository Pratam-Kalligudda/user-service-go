package repository

import (
	"github.com/Pratam-Kalligudda/user-service-go/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	UpdateUser(u domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo userRepository) CreateUser(u domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (repo userRepository) UpdateUser(u domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (repo userRepository) FindUserByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (repo userRepository) FindUserById(id uint) (domain.User, error) {
	return domain.User{}, nil
}
