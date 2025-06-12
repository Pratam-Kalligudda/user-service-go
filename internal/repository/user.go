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
	tx := repo.db.Create(&u)
	return u, tx.Error
}

func (repo userRepository) UpdateUser(u domain.User) (domain.User, error) {
	tx := repo.db.Model(&u).Updates(u)
	return u, tx.Error
}

func (repo userRepository) FindUserByEmail(email string) (user domain.User, err error) {
	err = repo.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (repo userRepository) FindUserById(id uint) (user domain.User, err error) {
	err = repo.db.First(&user, "id = ?", id).Error
	return user, err
}
