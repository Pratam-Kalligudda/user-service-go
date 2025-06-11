package helper

import (
	"github.com/Pratam-Kalligudda/user-service-go/internal/domain"
	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	secret string
}

func NewAuthHelper(secret string) Auth {
	return Auth{secret: secret}
}

func (a Auth) GenerateHashPassword(pass string) (string, error) {
	return "", nil
}

func (a Auth) ComparePassword(hashPass, pass string) error {
	return nil
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	return nil
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) error {
	return nil
}

func (a Auth) GenerateToken(id uint, role, email string) (string, error) {
	return "", nil
}

func (a Auth) VerifyToken(token string) (domain.User, error) {
	return domain.User{}, nil
}

func (a Auth) GenerateCode() (int, error) {
	return -1, nil
}
